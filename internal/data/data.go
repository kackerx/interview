package data

import (
	"context"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/kackerx/interview/internal/conf"
)

const ctxTxKey = "TxKey"

type Data struct {
	master *gorm.DB
	slave  *gorm.DB
}

// 用wire依赖注入参数不能同类型, 但是再搞结构体有点麻烦, 简单点吧
// type MasterDB struct {
// 	master *gorm.DB
// }
//
// type SlaveDB struct {
// 	slave *gorm.DB
// }

func (r *Data) DB(ctx context.Context) *gorm.DB {
	v := ctx.Value(ctxTxKey)
	if v != nil {
		if tx, ok := v.(*gorm.DB); ok {
			return tx
		}
	}
	return r.master.WithContext(ctx)
}

type Transaction interface {
	Transaction(ctx context.Context, fn func(ctx context.Context) error) error
}

func NewTransaction(r *Data) Transaction {
	return r
}

func (r *Data) Transaction(ctx context.Context, fn func(ctx context.Context) error) error {
	return r.master.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		ctx = context.WithValue(ctx, ctxTxKey, tx)
		return fn(ctx)
	})
}

func NewData(masterDB *gorm.DB) *Data {
	return &Data{master: masterDB}
}

func NewDb(cfg *conf.Conf) *gorm.DB {
	master := cfg.Data.Master
	db, err := gorm.Open(mysql.Open(master.Dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	sqlDb, _ := db.DB()
	sqlDb.SetMaxOpenConns(master.MaxOpen)
	sqlDb.SetMaxIdleConns(master.MaxIdle)
	sqlDb.SetConnMaxLifetime(time.Minute * time.Duration(master.MaxLifeTime))

	return db
}

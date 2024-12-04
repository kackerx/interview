package data

import (
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/kackerx/interview/internal/conf"
)

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

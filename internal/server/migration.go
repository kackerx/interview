package server

import (
	"gorm.io/gorm"

	"github.com/kackerx/interview/internal/data/model"
)

type Migrate struct {
	db *gorm.DB
}

func NewMigrate(db *gorm.DB) *Migrate {
	return &Migrate{db: db}
}

func (m *Migrate) Start() {
	if err := m.db.Migrator().AutoMigrate(&model.User{}, &model.Task{}, &model.Document{}); err != nil {
		panic(err)
	}
}

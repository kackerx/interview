package model

import (
	"time"

	"gorm.io/gorm"
)

type Document struct {
	gorm.Model
	// Status      string `gorm:"column:status;type:varchar(32);not null;default:''"`
	Content     string `gorm:"content;type:text;not null"`
	I18nContent string `gorm:"i18n_content;type:text;not null"`
	CreatedBy   string `gorm:"column:created_by;type:varchar(30);not null;default:''"`
	CompleteAt  time.Time
}

func (t *Document) TableName() string {
	return "t_document"
}

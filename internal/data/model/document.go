package model

import (
	"gorm.io/gorm"
)

type Document struct {
	gorm.Model
	// Status      string `gorm:"column:status;type:varchar(32);not null;default:''"`
	TaskID      uint   `gorm:"task_id;not null;default:0"`
	Content     string `gorm:"content;type:text;not null"`
	I18nContent string `gorm:"i18n_content;type:text;not null"`
	CreatedBy   string `gorm:"column:created_by;type:varchar(30);not null;default:''"`
}

func (t *Document) TableName() string {
	return "t_document"
}

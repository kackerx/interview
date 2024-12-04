package model

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Status string `gorm:"column:status;type:varchar(32);not null;default:''"`
	// Content   string `gorm:"content;type:text';not null"`
	CreatedBy string `gorm:"column:created_by;type:varchar(30);not null;default:''"`
}

func (t *Task) TableName() string {
	return "t_task"
}

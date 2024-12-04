package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName string `gorm:"column:user_name;type:varchar(30);not null;default:'';unique"`
	Password string `gorm:"column:password;type:varchar(128);not null;default:''"`
	NickName string `gorm:"column:nick_name;type:varchar(30);not null;default:''"`
	Avatar   string `gorm:"column:avatar;type:varchar(256);not null;default:''"`
	Email    string `gorm:"column:email;type:varchar(30);not null;default:''"`
	Status   int    `gorm:"column:status;type:tinyint;not null;default:1"`
	Gender   int    `gorm:"column:gender;type:tinyint;not null;default:1"`
}

func (u *User) TableName() string {
	return "t_user"
}

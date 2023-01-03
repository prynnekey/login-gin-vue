package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null;unique"`
	Password string `gorm:"type:varchar(20);not null"`
	Tel      string `gorm:"type:varchar(20);not null;unique"`
}

// 映射表的名称为user
func (*User) TableName() string {
	return "user"
}

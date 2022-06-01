package domain

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Nama     string `gorm:"type:varchar(50)"`
	Email    string `gorm:"unique"`
	Password string
	Token    string
}

package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDB() *gorm.DB {
	db, _ := gorm.Open(mysql.Open("root@tcp(localhost:3306)/go_api?parseTime=true"), &gorm.Config{})
	return db
}

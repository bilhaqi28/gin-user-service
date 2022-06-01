package migration

import (
	"gin-user-service/model/domain"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&domain.User{})
}

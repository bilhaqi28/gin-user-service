package controller

import (
	"gin-user-service/repository"
	"gin-user-service/service"

	"gorm.io/gorm"
)

type MainControllerImpl struct {
	db *gorm.DB
}

// UserController implements MainController
func (main *MainControllerImpl) UserController() ControllerUser {
	repository := repository.NewRepositoryUser()
	service := service.NewServiceUser(repository, main.db)
	controller := NewControllerUser(service)
	return controller
}

func NewMainController(db *gorm.DB) MainController {
	return &MainControllerImpl{
		db: db,
	}
}

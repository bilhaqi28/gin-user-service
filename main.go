package main

import (
	"gin-user-service/config"
	"gin-user-service/controller"
	"gin-user-service/migration"
	"gin-user-service/route"
)

func main() {
	db := config.NewDB()
	// proses migration
	migration.Migrate(db)
	// product
	controller := controller.NewMainController(db)
	router := route.NewRouter(controller)
	router.Run(":8080")
}

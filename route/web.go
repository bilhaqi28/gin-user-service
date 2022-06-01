package route

import (
	"gin-user-service/controller"
	"gin-user-service/exception"

	"github.com/gin-gonic/gin"
)

func NewRouter(controller controller.MainController) *gin.Engine {
	router := gin.Default()
	router.Use(gin.CustomRecovery(exception.InternalServerError))
	service := router.Group("/service")
	{
		service.POST("/register", func(c *gin.Context) {
			controller.UserController().Register(c)
		})
		service.POST("/login", func(c *gin.Context) {
			controller.UserController().Login(c)
		})
	}

	return router
}

package controller

import (
	"net/http"

	"gin-user-service/model/web/request"
	"gin-user-service/model/web/response"
	"gin-user-service/service"

	"github.com/gin-gonic/gin"
)

type ControllerUserImpl struct {
	serviceUser service.ServiceUser
}

// Login implements ControllerUser
func (controller *ControllerUserImpl) Login(c *gin.Context) {
	var request request.Login
	err := c.ShouldBindJSON(&request)
	if err != nil {
		errorResponse := response.ErrorResponse{
			Code:   http.StatusBadRequest,
			Status: false,
			Error:  err.Error(),
		}
		c.JSON(http.StatusBadRequest, errorResponse)
		return
	}
	userLogin, err := controller.serviceUser.Login(c.Request.Context(), request)
	if err != nil {
		errorResponse := response.ErrorResponse{
			Code:   http.StatusBadRequest,
			Status: false,
			Error:  err.Error(),
		}
		c.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	apiResponse := response.ApiResponse{
		Code:   http.StatusOK,
		Status: true,
		Data:   userLogin,
	}
	c.JSON(http.StatusOK, apiResponse)
}

// Destroy implements ControllerProduct

func (controller *ControllerUserImpl) Register(c *gin.Context) {
	var request request.CreateUser
	err := c.ShouldBindJSON(&request)
	if err != nil {
		errorResponse := response.ErrorResponse{
			Code:   http.StatusBadRequest,
			Status: false,
			Error:  err.Error(),
		}
		c.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	userRegister := controller.serviceUser.Register(c.Request.Context(), request)
	apiResponse := response.ApiResponse{
		Code:   http.StatusOK,
		Status: true,
		Data:   userRegister,
	}

	c.JSON(http.StatusOK, apiResponse)

}

func NewControllerUser(serviceUser service.ServiceUser) ControllerUser {
	return &ControllerUserImpl{
		serviceUser: serviceUser,
	}
}

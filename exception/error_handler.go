package exception

import (
	"net/http"

	"gin-user-service/model/web/response"

	"github.com/gin-gonic/gin"
)

func InternalServerError(c *gin.Context, err interface{}) {
	// exception, _ := err.(NotFoundError)
	errorResponse := response.ErrorResponse{
		Code:   http.StatusInternalServerError,
		Status: false,
		Error:  "Something went wrong",
	}
	c.JSON(http.StatusInternalServerError, errorResponse)
}

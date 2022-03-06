package utils

import (
	"jira-backend/constants"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetSuccessResponse(msg string, val interface{}) gin.H {
	return gin.H{
		constants.Status:   true,
		constants.Message:  msg,
		constants.Response: val,
	}
}

func GetErrorResponse(error_msg string) gin.H {
	return gin.H{
		constants.Status:  false,
		constants.Message: error_msg,
	}
}

func ThrowBadRequest(c *gin.Context) {
	c.JSON(http.StatusBadRequest, GetErrorResponse(constants.BAD_REQUEST))
	c.AbortWithStatus(http.StatusBadRequest)
}

func GetResponse(status bool, msg string, val interface{}) gin.H {
	return gin.H{
		constants.Status:   status,
		constants.Message:  msg,
		constants.Response: val,
	}
}

package utils

import (
	"jira-backend/constants"

	"github.com/gin-gonic/gin"
)

func GetResponse(status bool, msg string, val interface{}) gin.H {
	return gin.H{
		constants.Status:   status,
		constants.Message:  msg,
		constants.Response: val,
	}
}

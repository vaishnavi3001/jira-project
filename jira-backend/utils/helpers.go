package utils

import (
	"fmt"
	ct "jira-backend/constants"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func GetSuccessResponse(msg string, val interface{}) gin.H {
	return gin.H{
		ct.Status:   true,
		ct.Message:  msg,
		ct.Response: val,
	}
}

func GetErrorResponse(error_msg string) gin.H {
	return gin.H{
		ct.Status:  false,
		ct.Message: error_msg,
	}
}

func ThrowBadRequest(c *gin.Context) {
	c.JSON(http.StatusBadRequest, GetErrorResponse(ct.BAD_REQUEST))
	c.AbortWithStatus(http.StatusBadRequest)
}

func GetResponse(status bool, msg string, val interface{}) gin.H {
	return gin.H{
		ct.Status:   status,
		ct.Message:  msg,
		ct.Response: val,
	}
}

func DateFormat(dateStr string) (time.Time, error) {
	t, err := time.Parse(ct.SPRINT_DATE_FORMAT, dateStr)
	if err != nil {
		fmt.Println(err.Error())
	}
	return t, err
}

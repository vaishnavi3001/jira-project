package handlers

import "github.com/gin-gonic/gin"

func Userlogin(c *gin.Context) {
	c.JSON(200, gin.H{"message": "mandar"})
}

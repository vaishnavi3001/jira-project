package dbutils

import (
	"fmt"
	ct "jira-backend/constants"
	md "jira-backend/models"
	sk "jira-backend/skeletons"
	ut "jira-backend/utils"

	"github.com/gin-gonic/gin"
)

func CheckLoginCreds(username string, password string) int64 {
	var count int64
	var user md.User

	DB.Where("username = ? AND password = ?", username, password).Find(&user).Count(&count)
	if count == 1 {
		return int64(user.UserId)
	}
	return -1
}

func RegisterUser(data sk.UserRegister) gin.H {
	username := data.Username
	email := data.Email

	var count int64

	DB.Where("username = ? OR email_id = ?", username, email).Find(&md.User{}).Count(&count)
	fmt.Println(count)
	if count != 0 {
		return ut.GetErrorResponse(ct.USER_ALREADY_EXISTS)
	}

	DB.Create(&md.User{Username: username, EmailId: email, Firstname: data.FirstName, Lastname: data.LastName, Password: data.Password})

	return ut.GetSuccessResponse(ct.REGISTERATION_SUCCESSFUL, "")
}

func GetProfileUser(user_id uint) gin.H {

	var count int64
	var user md.User

	DB.Where("user_id = ?", user_id).Find(&user)
	fmt.Println(count)
	if count == 0 {
		return ut.GetErrorResponse(ct.USER_DOESNT_EXIST)
	}
	
	resp := sk.UserProfile{Username: user.Username, Email: user.EmailId, FirstName: user.Firstname, LastName: user.Lastname}
	return ut.GetSuccessResponse(ct.USER_FOUND, resp)
}
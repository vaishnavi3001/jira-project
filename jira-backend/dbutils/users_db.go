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

	DB.Where("user_id = ?", user_id).Find(&user).Count(&count)
	fmt.Println(count)
	if count == 0 {
		return ut.GetErrorResponse(ct.USER_DOESNT_EXIST)
	}
	
	resp := sk.UserProfile{Username: user.Username, Email: user.EmailId, FirstName: user.Firstname, LastName: user.Lastname}
	return ut.GetSuccessResponse(ct.USER_FOUND, resp)
}

func UpdateProfileUser(user_id uint, data sk.UserProfile) gin.H {
	username := data.Username
	email := data.Email
	var count int64
	var user md.User
	
	DB.Where("username = ? OR email_id = ?", username, email).Find(&user).Count(&count)
	
	if count != 0  && user.UserId != user_id{
		return ut.GetErrorResponse(ct.USER_ALREADY_EXISTS)
	}

	DB.Model(&md.User{}).Where("user_id = ?", user_id).Updates(md.User{Username: data.Username, EmailId: data.Email, Firstname: data.FirstName, Lastname: data.LastName})
	resp := sk.UserProfile{Username: data.Username, Email: data.Email, FirstName: data.FirstName, LastName: data.LastName}
	return ut.GetSuccessResponse(ct.USER_PROFILE_UPDATE_SUCCESS, resp)
}

func ChangePassword(user_id uint, data sk.ChangePasswordReq) gin.H {
	oldPassword := data.OldPassword
	
	newPassword := data.NewPassword
	fmt.Println(oldPassword)
	fmt.Println("###")
	fmt.Println(newPassword)
	fmt.Println("@@@")
	var count int64
	var user md.User
	
	DB.Where("user_id = ? AND password = ?", user_id, oldPassword).Find(&user).Count(&count)
	
	if count == 0{
		return ut.GetErrorResponse(ct.INVALID_CREDENTIALS)
	}

	DB.Model(&md.User{}).Where("user_id = ?", user_id).Updates(md.User{Password: newPassword})
	resp := ""
	return ut.GetSuccessResponse(ct.PASSWORD_CHANGE_SUCCESSFUL, resp)
}
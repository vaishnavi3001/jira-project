package dbutils

import (
	md "jira-backend/models"
)

func CheckTokenInDb(tokenStr string, userId uint) bool {
	var count int64
	DB.Where("token = ? AND user_id = ?", tokenStr, userId).Find(&md.UserAuth{}).Count(&count)

	return count == 1
}

func DeleteToken(tokenStr string) {
	DB.Where("token = ?", tokenStr).Delete(&md.UserAuth{})
}

func AddUserToken(tokenStr string, userId uint) {
	DB.Create(&md.UserAuth{UserId: userId, Token: tokenStr})
}

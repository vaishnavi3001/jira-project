package dbutils

import md "jira-backend/models"

func CheckTokenInDb(tokenStr string, userId uint) bool {
	var count int64
	db.Where("token = ? AND user_id = ?", tokenStr, userId).Find(&md.UserAuth{}).Count(&count)

	return count == 1
}

func DeleteToken(tokenStr string) {
	db.Where(" token = ?").Delete(&md.UserAuth{})
}

func AddUserToken(tokenStr string, userId uint) {
	db.Create(&md.UserAuth{UserId: userId, Token: tokenStr})
}

package dbutils

import md "jira-backend/models"

func CheckLoginCreds(username string, password string) int64 {
	var count int64
	var user md.User

	db.Where("username = ? AND password = ?", username, password).Find(&user).Count(&count)
	if count == 1 {
		return int64(user.UserId)
	}
	return -1
}

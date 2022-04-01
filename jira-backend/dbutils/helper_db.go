package dbutils

import md "jira-backend/models"

func IsUserPartOfTheProject(userId uint, projectId uint) (bool, md.UserRole) {
	var count int64
	var userRole md.UserRole
	DB.Preload("User").Where("user_id = ? AND project_id = ?", userId, projectId).Find(&userRole).Count(&count)

	if count == 0 {
		return false, md.UserRole{}
	}

	return true, userRole
}

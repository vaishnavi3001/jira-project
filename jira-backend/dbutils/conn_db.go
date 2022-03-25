package dbutils

import (
	"jira-backend/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func InitializeConn() error {
	db, err = gorm.Open(sqlite.Open("jira.db"), &gorm.Config{})
	if err != nil {
		return err
	}

	db.AutoMigrate(&models.User{}, &models.Project{}, &models.Sprint{}, &models.Issue{}, &models.Permission{}, &models.Issue{}, &models.UserRole{}, &models.UserAuth{})
	return nil
}

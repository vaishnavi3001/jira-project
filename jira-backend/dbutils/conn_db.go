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

func InitializeConn() {
	db, err = gorm.Open(sqlite.Open("jira.db"), &gorm.Config{})
	if err != nil {
		panic("Could Not Initialize Connection To The Database")

	}

	db.AutoMigrate(&models.User{}, &models.Project{}, &models.Sprint{}, &models.Issue{}, &models.Permission{}, &models.Issue{}, &models.UserRole{})
}

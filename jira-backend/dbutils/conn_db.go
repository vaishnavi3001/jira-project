package dbutils

import (
	"fmt"
	md "jira-backend/models"
	ut "jira-backend/utils"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func InitializeConn(isTest bool) error {
	db_name := ut.Vconfig.GetString("db.db_name")
	if isTest {
		db_name = ut.Vconfig.GetString("db.test_db_name")
	}

	DB, err = gorm.Open(sqlite.Open(db_name), &gorm.Config{})
	if err != nil {
		return err
	}
	fmt.Println(db_name)
	DB.AutoMigrate(&md.User{}, &md.Project{}, &md.Sprint{}, &md.Issue{}, &md.Permission{}, &md.Issue{}, &md.UserRole{}, &md.UserAuth{}, &md.Comment{}, &md.Invite{})
	return nil
}

func GetDBInstance(isTest bool) *gorm.DB {
	if err := InitializeConn(isTest); err != nil {
		return nil
	}
	return DB
}

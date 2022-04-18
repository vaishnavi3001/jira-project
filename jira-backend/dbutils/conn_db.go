package dbutils

import (
	ct "jira-backend/constants"
	md "jira-backend/models"
	ut "jira-backend/utils"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func InitializeConn(env string) error {
	db_name := ut.Vconfig.GetString("db.db_name")
	if env == ct.TEST {
		db_name = ut.Vconfig.GetString("db.test_db_name")
	}

	DB, err = gorm.Open(sqlite.Open(db_name), &gorm.Config{})
	if err != nil {
		return err
	}

	DB.AutoMigrate(&md.User{}, &md.Project{}, &md.Sprint{}, &md.Issue{}, &md.Permission{}, &md.Issue{}, &md.UserRole{}, &md.UserAuth{}, &md.Comment{})
	return nil
}

func GetDBInstance(env string) *gorm.DB {
	if err := InitializeConn(env); err != nil {
		return nil
	}
	return DB
}

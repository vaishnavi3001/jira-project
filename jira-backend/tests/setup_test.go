package tests

import (
	"fmt"
	dt "jira-backend/dbutils"
	md "jira-backend/models"
	ut "jira-backend/utils"
	"os"
)

func InitConfig() {
	if err := ut.ConfigReader(true); err != nil {
		fmt.Println("Couldn't read config.. Exiting Tests")
		os.Exit(-1)
	}
}
func DBInitForTest() {
	os.Remove(ut.Vconfig.GetString("test_db_name"))
	dt.GetDBInstance(true)
	if dt.DB == nil {
		fmt.Println("Couldn't create db.. Exiting Tests")
		os.Exit(-1)
	}
}

func DBFree() {
	if dt.DB != nil {
		sqlDB, err := dt.DB.DB()
		if err != nil {
			return
		}
		sqlDB.Close()
		os.Remove(ut.Vconfig.GetString("test_db_name"))
	}
}

func DBInsertRows() {
	user1 := &md.User{Username: "pypalkar23", Password: "dd29b8cb089a56606fca480e137c27c4", Firstname: "Mandar", Lastname: "Palkar", EmailId: "mandar.palkar23@gmail.com"}
	dt.DB.Create(user1)
	user2 := &md.User{Username: "amhaske32", Password: "7b69ad8a8999d4ca7c42b8a729fb0ffd", Firstname: "Ashish", Lastname: "Mhaske", EmailId: "amhaske32@gmail.com"}
	dt.DB.Create(user2)
	user_auth1 := &md.UserAuth{UserRef: user1.UserId, Token: TokenStr}
	dt.DB.Create(&user_auth1)
}

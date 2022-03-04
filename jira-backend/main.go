package main

import (
	"fmt"
	"jira-backend/handlers"
	"jira-backend/models"
	"jira-backend/utils"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error
var config *viper.Viper

func setDbconnInContext(c *gin.Context) {
	c.Set("db", db)
	c.Next()
}

func setConfigInContext(c *gin.Context) {
	c.Set("config", config)
	c.Next()
}

func main() {
	db, err = gorm.Open(sqlite.Open("jira.db"), &gorm.Config{})
	if err != nil {
		fmt.Println("Cannot Connect to the Database Exiting", err)
		os.Exit(1)
	}

	// Migrate the schema
	db.AutoMigrate(&models.User{}, &models.Project{}, &models.Sprint{}, &models.Issue{}, &models.Permission{}, &models.Issue{}, &models.UserRole{})
	//,&models.Project{}, &models.Issue{}, &models.Sprint{},  &models.UserRole{}, &models.Role{}, &models.Permission{}

	//viper config
	config, err = utils.ConfigReader()
	if err != nil {
		fmt.Println("Cannot Read from the Config", err)
		os.Exit(1)
	}

	r := gin.Default()
	r.Use(setConfigInContext)
	r.Use(setDbconnInContext)

	/*user := models.User{Name: "Mandar Palkar", Username: "pypalkar23", EmailId: "mandar.palkar23@gmail.com"}
	db.Create(&user)
	user1 := models.User{Name: "Ashish Mhaske", Username: "amhaske32", EmailId: "amhaske32@gmail.com"}
	db.Create(&user1)*/

	project := r.Group("/project")
	{
		project.POST("/create", handlers.CreateProject)
		project.GET("/list", handlers.ListProjects)
		project.GET("/info", handlers.GetProjectInfo)
		project.POST("/delete", handlers.DeleteProject)
	}

	issue := r.Group("/issues")
	{
		issue.POST("/create", handlers.CreateIssue)
		issue.GET("/list", handlers.ListIssue)
		issue.GET("/info", handlers.GetIssueInfo)
		issue.POST("/delete", handlers.DeleteIssue)
		issue.POST("/update", handlers.UpdateIssue)
	}

	ip_address := fmt.Sprintf("%s:%d", config.GetString("server.ip_address"), config.GetInt("server.port"))
	r.Run(ip_address)

	/*project1 := models.Project{Name: "Jira-Clone", OwnerId: 1, CreatedAt: time.Now()}
	db.Create(&project1)*/

}

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

	//
	db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		fmt.Println("Cannot Connect to the Database Exiting", err)
		os.Exit(1)
	}

	// Migrate the schema
	db.AutoMigrate(&models.Project{}, &models.Issue{}, &models.Sprint{}, &models.User{}, &models.UserRole{}, &models.Role{}, &models.Permission{})

	//viper config
	config, err = utils.ConfigReader()
	if err != nil {
		fmt.Println("Cannot Read from the Config", err)
		os.Exit(1)
	}

	r := gin.Default()
	r.Use(setConfigInContext)
	r.Use(setDbconnInContext)

	project := r.Group("/project")
	{
		project.POST("/create", handlers.CreateProject)
		project.GET("/list", handlers.ListProjects)
		project.GET("/info", handlers.GetProjectInfo)
	}

	ip_address := fmt.Sprintf("%s:%d", config.GetString("server.ip_address"), config.GetInt("server.port"))
	r.Run(ip_address)

}

package main

import (
	"fmt"
	"jira-backend/dbutils"
	"jira-backend/handlers"
	"jira-backend/utils"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error
var config *viper.Viper

func corsMiddleware(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(204)
		return
	}

	c.Next()
}

func setConfigInContext(c *gin.Context) {
	c.Set("config", config)
	c.Next()
}

func main() {

	// Migrate the schema
	dbutils.InitializeConn()
	//,&models.Project{}, &models.Issue{}, &models.Sprint{},  &models.UserRole{}, &models.Role{}, &models.Permission{}

	//viper config
	config, err = utils.ConfigReader()
	if err != nil {
		fmt.Println("Cannot Read from the Config", err)
		os.Exit(1)
	}

	r := gin.Default()
	r.Use(setConfigInContext)
	r.Use(corsMiddleware)
	/*user := models.User{Name: "Mandar Palkar", Username: "pypalkar23", EmailId: "mandar.palkar23@gmail.com"}
	db.Create(&user)
	user1 := models.User{Name: "Ashish Mhaske", Username: "amhaske32", EmailId: "amhaske32@gmail.com"}
	db.Create(&user1)*/

	project := r.Group("/project")
	{
		project.GET("/list", handlers.ListProjects)
		project.GET("/info", handlers.GetProjectInfo)
		project.POST("/delete", handlers.DeleteProject)
		project.POST("/create", handlers.CreateProject)
	}

	sprint := r.Group("/sprint")
	{
		sprint.GET("/list", handlers.ListSprints)
		sprint.GET("/info", handlers.GetSprintInfo)
		//sprint.POST("/create", handlers.CreateSprint)
		sprint.POST("/delete", handlers.DeleteSprint)

	}

	issue := r.Group("/issues")
	{
		issue.GET("/list", handlers.ListIssue)
		issue.GET("/info", handlers.GetIssueInfo)
		issue.POST("/create", handlers.CreateIssue)
		issue.POST("/delete", handlers.DeleteIssue)
		issue.POST("/update", handlers.UpdateIssue)
	}

	ip_address := fmt.Sprintf("%s:%d", config.GetString("server.ip_address"), config.GetInt("server.port"))
	r.Run(ip_address)

}

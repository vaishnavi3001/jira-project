package main

import (
	"fmt"
	dt "jira-backend/dbutils"
	hd "jira-backend/handlers"
	ut "jira-backend/utils"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	// Migrate the schema
	dberr := dt.InitializeConn()
	if dberr != nil {
		fmt.Println("Cannot Connect the DB Details")
		os.Exit(1)
	}

	r := gin.Default()
	err := ut.ConfigReader()
	if err != nil {
		fmt.Println("Cannot Read from the Config", err)
		os.Exit(1)
	}
	r.Use(cors.Default())

	r.POST("/login", hd.Userlogin)
	r.GET("/logout", hd.AuthInterceptor, hd.UserLogout)

	api := r.Group("api")
	api.Use(hd.AuthInterceptor)

	project := api.Group("project")
	project.POST("/list", hd.ListProjects)
	project.POST("/info", hd.GetProjectInfo)
	project.POST("/delete", hd.DeleteProject)
	project.POST("/create", hd.CreateProject)
	project.POST("/members", hd.ListMembers)

	sprint := api.Group("sprint")
	sprint.POST("/list", hd.ListSprints)
	sprint.POST("/info", hd.GetSprintInfo)
	sprint.POST("/create", hd.CreateSprint)
	sprint.POST("/delete", hd.DeleteSprint)
	sprint.POST("/update", hd.UpdateSprint)

	issue := api.Group("issue")
	issue.POST("/info", hd.GetIssueInfo)
	issue.POST("/create", hd.CreateIssue)
	issue.POST("/delete", hd.DeleteIssue)
	issue.POST("/update", hd.UpdateIssue)
	issue.POST("/list", hd.ListIssuesForSprint)
	issue.POST("/move", hd.UpdateIssueStatus)

	ip_address := fmt.Sprintf("%s:%d", ut.Vconfig.GetString("server.ip_address"), ut.Vconfig.GetInt("server.port"))
	r.Run(ip_address)
}

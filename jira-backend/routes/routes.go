package routes

import (
	hd "jira-backend/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter(useAuth bool) *gin.Engine {
	r := gin.Default()

	r.Use(cors.Default())

	r.POST("/login", hd.Userlogin)
	r.POST("/register", hd.UserRegister)
	r.GET("/logout", hd.AuthInterceptor, hd.UserLogout)

	api := r.Group("api")

	if useAuth {
		api.Use(hd.AuthInterceptor)
	}

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

	comment := api.Group("comment")
	comment.POST("/add", hd.AddComment)
	comment.POST("/edit", hd.EditComment)
	comment.POST("/view", hd.ViewComment)
	return r
}

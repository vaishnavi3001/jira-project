package handlers

import (
	dt "jira-backend/dbutils"
	sk "jira-backend/skeletons"
	ut "jira-backend/utils"

	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateProject(c *gin.Context) {
	user_id := ut.GetUserIdFromContext(c)
	var createProjectReq sk.CreateProjectReq
	if err := c.BindJSON(&createProjectReq); err != nil {
		ut.ThrowBadRequest(c)
		return
	}

	//the one who creates the project will be owner of the project
	c.JSON(http.StatusOK, dt.CreateProject(createProjectReq, user_id))
}

func ListProjects(c *gin.Context) {
	user_id := ut.GetUserIdFromContext(c)
	c.JSON(http.StatusOK, dt.GetProjectList(user_id))
}

func GetProjectInfo(c *gin.Context) {
	user_id := ut.GetUserIdFromContext(c)
	var req sk.ProjectInfoReq
	if err := c.BindJSON(&req); err != nil {
		ut.ThrowBadRequest(c)
		return
	}

	c.JSON(http.StatusOK, dt.GetProjectInfo(req, user_id))
}

func DeleteProject(c *gin.Context) {
	user_id := ut.GetUserIdFromContext(c)
	var req sk.BaseProjectIdReq
	if err := c.BindJSON(&req); err != nil {
		ut.ThrowBadRequest(c)
		return
	}

	c.JSON(http.StatusOK, dt.DeleteProject(req, user_id))
}

func ListMembers(c *gin.Context) {
	user_id := ut.GetUserIdFromContext(c)
	var req sk.BaseProjectIdReq

	if err := c.BindJSON(&req); err != nil {
		ut.ThrowBadRequest(c)
		return
	}

	c.JSON(http.StatusOK, dt.ListMembers(req, user_id))
}

func ShowStats(c *gin.Context) {
	user_id := ut.GetUserIdFromContext(c)
	var req sk.BaseProjectIdReq

	if err := c.BindJSON(&req); err != nil {
		ut.ThrowBadRequest(c)
		return
	}

	c.JSON(http.StatusOK, dt.ShowStats(req, user_id))
}
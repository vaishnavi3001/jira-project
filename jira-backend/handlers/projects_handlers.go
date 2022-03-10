package handlers

import (
	"jira-backend/dbutils"
	sk "jira-backend/skeletons"
	ut "jira-backend/utils"

	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateProject(c *gin.Context) {

	var createProjectReq sk.CreateProjectReq
	if err := c.BindJSON(&createProjectReq); err != nil {
		ut.ThrowBadRequest(c)
		return
	}

	//the one who creates the project will be owner of the project
	c.JSON(http.StatusOK, dbutils.CreateProject(createProjectReq))
}

func ListProjects(c *gin.Context) {

	var req sk.UsersBaseReq
	if err := c.BindJSON(&req); err != nil {
		ut.ThrowBadRequest(c)
		return
	}

	c.JSON(http.StatusOK, dbutils.GetProjectList(req))
}

func GetProjectInfo(c *gin.Context) {
	var req sk.ProjectInfoReq
	if err := c.BindJSON(&req); err != nil {
		ut.ThrowBadRequest(c)
		return
	}

	c.JSON(http.StatusOK, dbutils.GetProjectInfo(req))
}

func DeleteProject(c *gin.Context) {

	var req sk.BaseProjectIdReq
	if err := c.BindJSON(&req); err != nil {
		ut.ThrowBadRequest(c)
		return
	}

	c.JSON(http.StatusOK, dbutils.DeleteProject(req))

}
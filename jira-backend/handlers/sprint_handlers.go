package handlers

import (
	dt "jira-backend/dbutils"
	sk "jira-backend/skeletons"
	ut "jira-backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

//list project
//new sprint
//delete project
//add member to project
//transfer ownership

func CreateSprint(c *gin.Context) {
	var req sk.CreateSprintReq

	if err := c.BindJSON(&req); err != nil {
		ut.ThrowBadRequest(c)
		return
	}

	c.JSON(http.StatusOK, dt.AddSprint(req))
}

func ListSprints(c *gin.Context) {

	var req sk.SprintListReq
	if err := c.BindJSON(&req); err != nil {
		ut.ThrowBadRequest(c)
		return
	}

	c.JSON(http.StatusOK, dt.GetSprintList(req))
}

func GetSprintInfo(c *gin.Context) {
	var req sk.SprintInfoReq
	if err := c.BindJSON(&req); err != nil {
		ut.ThrowBadRequest(c)
		return
	}

	c.JSON(http.StatusOK, dt.GetSprintInfo(req))
}

func UpdateSprintInfo(c *gin.Context) {

}

func DeleteSprint(c *gin.Context) {
	var req sk.SprintDeleteReq
	if err := c.BindJSON(&req); err != nil {
		ut.ThrowBadRequest(c)
		return
	}

	c.JSON(http.StatusOK, dt.DeleteSprint(req))
}

func GetSprintsForProject(c *gin.Context) {
	var req sk.SprintListReq
	if err := c.BindJSON(&req); err != nil {
		ut.ThrowBadRequest(c)
		return
	}

	c.JSON(http.StatusOK, dt.GetSprintList(req))
}

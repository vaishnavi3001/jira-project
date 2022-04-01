package handlers

import (
	dt "jira-backend/dbutils"
	sk "jira-backend/skeletons"
	ut "jira-backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateSprint(c *gin.Context) {
	var req sk.CreateSprintReq
	user_id := ut.GetUserIdFromContext(c)
	if err := c.BindJSON(&req); err != nil {
		ut.ThrowBadRequest(c)
		return
	}

	c.JSON(http.StatusOK, dt.AddSprint(req, user_id))
}

func ListSprints(c *gin.Context) {
	user_id := ut.GetUserIdFromContext(c)
	var req sk.SprintListReq
	if err := c.BindJSON(&req); err != nil {
		ut.ThrowBadRequest(c)
		return
	}

	c.JSON(http.StatusOK, dt.GetSprintList(req, user_id))
}

func GetSprintInfo(c *gin.Context) {
	var req sk.SprintInfoReq
	user_id := ut.GetUserIdFromContext(c)
	if err := c.BindJSON(&req); err != nil {
		ut.ThrowBadRequest(c)
		return
	}

	c.JSON(http.StatusOK, dt.GetSprintInfo(req, user_id))
}

func UpdateSprint(c *gin.Context) {
	var req sk.SprintUpdateReq
	user_id := ut.GetUserIdFromContext(c)
	if err := c.BindJSON(&req); err != nil {
		ut.ThrowBadRequest(c)
		return
	}

	c.JSON(http.StatusOK, dt.UpdateSprintInfo(req, user_id))
}

func DeleteSprint(c *gin.Context) {
	var req sk.SprintDeleteReq
	user_id := ut.GetUserIdFromContext(c)
	if err := c.BindJSON(&req); err != nil {
		ut.ThrowBadRequest(c)
		return
	}

	c.JSON(http.StatusOK, dt.DeleteSprint(req, user_id))
}

func GetSprintsForProject(c *gin.Context) {
	var req sk.SprintListReq
	user_id := ut.GetUserIdFromContext(c)
	if err := c.BindJSON(&req); err != nil {
		ut.ThrowBadRequest(c)
		return
	}

	c.JSON(http.StatusOK, dt.GetSprintList(req, user_id))
}

func ListIssuesForSprint(c *gin.Context) {
	var req sk.BaseSprintReq
	user_id := ut.GetUserIdFromContext(c)
	if err := c.BindJSON(&req); err != nil {
		ut.ThrowBadRequest(c)
		return
	}

	c.JSON(http.StatusOK, dt.GetIssuesForSprint(req, user_id))
}

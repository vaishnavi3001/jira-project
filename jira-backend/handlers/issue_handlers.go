package handlers

import (
	dt "jira-backend/dbutils"
	sk "jira-backend/skeletons"
	ut "jira-backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateIssue(c *gin.Context) {
	var req sk.AddIssueReq
	user_id := ut.GetUserIdFromContext(c)
	if err := c.BindJSON(&req); err != nil {
		ut.ThrowBadRequest(c)
		return
	}

	c.JSON(http.StatusOK, dt.CreateIssue(req, user_id))
}

func UpdateIssue(c *gin.Context) {
	user_id := ut.GetUserIdFromContext(c)
	var req sk.UpdateIssueReq
	if err := c.BindJSON(&req); err != nil {
		ut.ThrowBadRequest(c)
		return
	}

	c.JSON(http.StatusOK, dt.UpdateIssue(req, user_id))
}

func DeleteIssue(c *gin.Context) {
	var req sk.IssueBaseReq
	user_id := ut.GetUserIdFromContext(c)
	if err := c.BindJSON(&req); err != nil {
		ut.ThrowBadRequest(c)
		return
	}

	c.JSON(http.StatusOK, dt.DeleteIssue(req, user_id))
}

func GetIssueInfo(c *gin.Context) {
	user_id := ut.GetUserIdFromContext(c)
	var req sk.IssueBaseReq
	if err := c.BindJSON(&req); err != nil {
		ut.ThrowBadRequest(c)
		return
	}

	c.JSON(http.StatusOK, dt.GetIssueInfo(req, user_id))
}

func UpdateIssueStatus(c *gin.Context) {
	user_id := ut.GetUserIdFromContext(c)
	var req sk.UpdateIssueStatusReq
	if err := c.BindJSON(&req); err != nil {
		ut.ThrowBadRequest(c)
		return
	}

	c.JSON(http.StatusOK, dt.UpdateIssueStatus(req, user_id))
}

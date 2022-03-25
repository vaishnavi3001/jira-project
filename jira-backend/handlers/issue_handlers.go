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

	if err := c.BindJSON(&req); err != nil {
		ut.ThrowBadRequest(c)
		return
	}

	c.JSON(http.StatusOK, dt.CreateIssue(req))
}

func UpdateIssue(c *gin.Context) {

	var req sk.UpdateIssueReq
	if err := c.BindJSON(&req); err != nil {
		ut.ThrowBadRequest(c)
		return
	}

	c.JSON(http.StatusOK, dt.UpdateIssue(req))
}

func DeleteIssue(c *gin.Context) {
	var req sk.IssueBaseReq
	if err := c.BindJSON(&req); err != nil {
		ut.ThrowBadRequest(c)
		return
	}

	c.JSON(http.StatusOK, dt.DeleteIssue(req))
}

func GetIssueInfo(c *gin.Context) {
	var req sk.IssueBaseReq
	if err := c.BindJSON(&req); err != nil {
		ut.ThrowBadRequest(c)
		return
	}

	c.JSON(http.StatusOK, dt.GetIssueInfo(req))
}

func UpdateIssueStatus(c *gin.Context) {
	var req sk.UpdateIssueStatusReq
	if err := c.BindJSON(&req); err != nil {
		ut.ThrowBadRequest(c)
		return
	}

	c.JSON(http.StatusOK, dt.UpdateIssueStatus(req))
}

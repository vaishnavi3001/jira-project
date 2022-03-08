package handlers

import (
	dt "jira-backend/dbutils"
	"jira-backend/models"
	"jira-backend/skeletons"
	"jira-backend/utils"
	ut "jira-backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateIssue(c *gin.Context) {
	var req skeletons.AddIssueReq

	if err := c.BindJSON(&req); err != nil {
		ut.ThrowBadRequest(c)
		return
	}

	//sprint id, project id, issue_type, issue description, createdBy, AssignedTo, title
	c.JSON(http.StatusOK, dt.CreateIssue(req))
}

func UpdateIssue(c *gin.Context) {

	var req skeletons.UpdateIssueReq
	if err := c.BindJSON(&req); err != nil {
		ut.ThrowBadRequest(c)
		return
	}

	c.JSON(http.StatusOK, dt.UpdateIssue(req))
}

func DeleteIssue(c *gin.Context) {
	var req skeletons.IssueBaseReq
	if err := c.BindJSON(&req); err != nil {
		ut.ThrowBadRequest(c)
		return
	}

	c.JSON(http.StatusOK, dt.DeleteIssue(req))
}

func ListIssue(c *gin.Context) {
	db, exists := c.Keys["db"].(*gorm.DB)

	if !exists {
		c.JSON(http.StatusInternalServerError, utils.GetResponse(false, "Something went wrong", ""))
		c.AbortWithStatus(http.StatusInternalServerError)
	}

	var req skeletons.IssueListReq
	if err := c.BindJSON(&req); err != nil {

		c.JSON(http.StatusBadRequest, utils.GetResponse(false, "Could not parse the request", ""))
		c.AbortWithStatus(http.StatusBadRequest)
	}

	var issues []models.Issue
	db.Where("project_ref = ? AND sprint_ref= ?", req.ProjectId, req.SprintId).Find(&issues)

	var res []skeletons.IssueEntry
	for _, x := range issues {
		res = append(res, skeletons.IssueEntry{IssueId: x.IssueId, Name: x.Title, Status: x.Status})
	}

	c.JSON(http.StatusOK, utils.GetResponse(true, "", skeletons.IssueListResp{Issues: res}))
}

func GetIssueInfo(c *gin.Context) {
	db, exists := c.Keys["db"].(*gorm.DB)

	if !exists {
		c.JSON(http.StatusInternalServerError, utils.GetResponse(false, "Something went wrong", ""))
		c.AbortWithStatus(http.StatusInternalServerError)
	}

	var req skeletons.IssueBaseReq
	if err := c.BindJSON(&req); err != nil {

		c.JSON(http.StatusBadRequest, utils.GetResponse(false, "Could not parse the request", ""))
		c.AbortWithStatus(http.StatusBadRequest)
	}

	var issue models.Issue
	db.Preload("Creator").Preload("AssignedTo").Where("issue_id", req.IssueId).Find(&issue)

	c.JSON(http.StatusOK, utils.GetResponse(true, "Issue", skeletons.IssueEntryDetailed{IssueId: issue.IssueId, Name: issue.Title, Status: issue.Status, Description: issue.Description, CreatedBy: issue.Creator.Username, AssignedTo: issue.AssignedTo.Username, ModifiedAt: issue.UpdatedAt, CreatedAt: issue.CreatedAt}))
}

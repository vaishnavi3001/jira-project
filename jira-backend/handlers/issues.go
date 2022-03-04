package handlers

import (
	"jira-backend/models"
	"jira-backend/skeletons"
	"jira-backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateIssue(c *gin.Context) {
	db, exists := c.Keys["db"].(*gorm.DB)

	if !exists {
		c.JSON(http.StatusInternalServerError, utils.GetResponse(false, "Something went wrong", ""))
		c.AbortWithStatus(http.StatusInternalServerError)
	}

	var req skeletons.AddIssueReq
	if err := c.BindJSON(&req); err != nil {

		c.JSON(http.StatusBadRequest, utils.GetResponse(false, "Could not parse the request", ""))
		c.AbortWithStatus(http.StatusBadRequest)
	}

	issue := models.Issue{Title: req.IssueTitle, Description: req.IssueText, Type: req.IssueType, CreatedBy: req.Creator, AssigneeId: req.Assignee}
	db.Create(&issue)

	//sprint id, project id, issue_type, issue description, createdBy, AssignedTo, title
	c.JSON(http.StatusOK, utils.GetResponse(true, "Issue Created Successfully", skeletons.BaseIssueResp{IssueName: issue.Title, IssueId: issue.IssueId}))
}

func UpdateIssue(c *gin.Context) {
	db, exists := c.Keys["db"].(*gorm.DB)

	if !exists {
		c.JSON(http.StatusInternalServerError, utils.GetResponse(false, "Something went wrong", ""))
		c.AbortWithStatus(http.StatusInternalServerError)
	}

	var req skeletons.UpdateIssueReq
	if err := c.BindJSON(&req); err != nil {

		c.JSON(http.StatusBadRequest, utils.GetResponse(false, "Could not parse the request", ""))
		c.AbortWithStatus(http.StatusBadRequest)
	}

	db.Model(&models.Issue{}).Where("issue_id = ?", req.IssueId).Updates(models.Issue{Title: req.IssueTitle, Type: req.IssueType, Description: req.IssueText, Status: req.Status})
}

func DeleteIssue(c *gin.Context) {
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

	db.Where("issue_id = ?", req.IssueId).Delete(&models.Issue{})
	c.JSON(http.StatusOK, utils.GetResponse(true, "Issue deleted successfully", ""))
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
	db.Where("project_ref = ? AND sprint_ref=?", req.ProjectId, req.SprintId).Find(&issues)

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

	//var project models.Project
	//db.Preload("User").Where("project_id", req.ProjectId).Find(&project)
	/*IssueId     uint `gorm:"primaryKey;auto_increment;not_null"`
	Status      uint `gorm:"default:1"`
	Type        uint `gorm:"default:1"` //`epic/task/subtask/bug/`
	Title       string
	CreatedBy   uint //user ref
	SprintRef   uint
	AssigneeId  uint //user ref
	ProjectRef  uint
	IsDeleted   sql.NullBool
	UpdatedAt   time.Time
	Description string
	CreatedAt   time.Time `gorm:"autoCreateTime:milli"`
	Sprint      Sprint    `gorm:"foreignKey:SprintRef"`
	AssignedTo  User      `gorm:"foreignKey:AssigneeId"`
	Creator     User      `gorm:"foreignKey:CreatedBy"`
	Project     Project   `gorm:"foreignKey:ProjectRef"`
	DeletedAt   gorm.DeletedAt*/

	var issue models.Issue
	db.Preload("Creator").Preload("AssignedTo").Where("issue_id", req.IssueId).Find(&issue)

	c.JSON(http.StatusOK, utils.GetResponse(true, "Issue", skeletons.IssueEntryDetailed{IssueId: issue.IssueId, Name: issue.Title, Status: issue.Status, Description: issue.Description, CreatedBy: issue.Creator.Username, AssignedTo: issue.AssignedTo.Username}))
}

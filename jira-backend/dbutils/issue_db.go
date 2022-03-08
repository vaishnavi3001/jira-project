package dbutils

import (
	"fmt"
	ct "jira-backend/constants"
	md "jira-backend/models"
	sk "jira-backend/skeletons"
	ut "jira-backend/utils"

	"github.com/gin-gonic/gin"
)

func CreateIssue(data sk.AddIssueReq) gin.H {

	issue := md.Issue{Title: data.IssueTitle, Description: data.IssueText, Type: data.IssueType, CreatedBy: data.Creator, AssigneeId: data.Assignee, SprintRef: data.SprintId, ProjectRef: data.ProjectId}
	db.Create(&issue)

	count := int64(0)
	var creator md.User
	db.Where("user_id = ?", data.Creator).Find(&creator).Count(&count)
	if count == 0 {
		return ut.GetErrorResponse(ct.USER_DOESNT_EXIST)
	}
	var assignee md.User
	db.Where("user_id = ?", data.Assignee).Find(&assignee).Count(&count)
	if count == 0 {
		return ut.GetErrorResponse(ct.USER_DOESNT_EXIST)
	}

	resp := sk.AddIssueResp{IssueTitle: issue.Title, IssueText: issue.Description, IssueType: issue.Type, Creator: creator.Username, Assignee: assignee.Username, CreatedAt: issue.CreatedAt}
	return ut.GetSuccessResponse("", resp)
}

func GetIssueInfo(data sk.ProjectInfoReq) gin.H {
	var res sk.ProjectInfoResp
	var count int64
	var userRole md.UserRole
	db.Where("user_id=? AND project_id=?", data.UserId, data.ProjectId).Find(&userRole).Count(&count)
	if count == 0 {
		return ut.GetErrorResponse(ct.NOT_PART_OF_THE_PROJECT)
	} else {
		db.Preload("User").Where("role_id=1 AND project_id=?", data.ProjectId).Find(&userRole)
		var project md.Project
		fmt.Println(userRole)
		db.Where("project_id", data.ProjectId).Find(&project)
		res = sk.ProjectInfoResp{ProjectId: project.ProjectId, Name: project.ProjectName, CreatedAt: project.CreatedAt, OwnerUName: userRole.User.Username, OwnerFName: userRole.User.Firstname, OwnerLName: userRole.User.Lastname, OwnerId: userRole.User.UserId}
		return ut.GetSuccessResponse("", res)
	}
}

func DeleteIssue(data sk.BaseProjectIdReq) gin.H {
	var count int64
	var userRole md.UserRole
	db.Where("user_id = ? AND project_id = ? AND role_id = 1", data.UserId, data.ProjectId).Find(&userRole).Count(&count)
	if count == 0 {
		return ut.GetErrorResponse(ct.ACTION_NOT_AUTHORIZED)
	}

	project := md.Project{ProjectId: data.ProjectId}
	db.Delete(&project)
	db.Where("project_id=?", data.ProjectId).Delete(&md.UserRole{})
	return ut.GetSuccessResponse(ct.PROJECT_DELETE_SUCCESS, "")
}

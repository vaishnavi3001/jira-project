package dbutils

import (
	ct "jira-backend/constants"
	md "jira-backend/models"
	sk "jira-backend/skeletons"
	ut "jira-backend/utils"

	"github.com/gin-gonic/gin"
)

func CreateIssue(data sk.AddIssueReq, userId uint) gin.H {
	res, creator := IsUserPartOfTheProject(userId, data.ProjectId)
	if !res {
		return ut.GetErrorResponse(ct.USER_DOESNT_EXIST)
	}

	res1, assignee := IsUserPartOfTheProject(data.Assignee, data.ProjectId)
	if !res1 {
		return ut.GetErrorResponse(ct.USER_DOESNT_EXIST)
	}

	issue := md.Issue{Title: data.IssueTitle, Description: data.IssueText, Type: data.IssueType, CreatedBy: userId, AssigneeId: data.Assignee, SprintRef: data.SprintId, ProjectRef: data.ProjectId, Status: ct.Created}
	DB.Create(&issue)

	resp := sk.AddIssueResp{IssueTitle: issue.Title, IssueText: issue.Description, IssueType: issue.Type, Creator: creator.User.Username, Assignee: assignee.User.Username, CreatedAt: issue.CreatedAt, Status: issue.Status}
	return ut.GetSuccessResponse("", resp)
}

func GetIssueInfo(data sk.IssueBaseReq, userId uint) gin.H {
	var count int64
	var issue md.Issue
	DB.Joins("JOIN projects ON projects.project_id = issues.project_ref AND issues.issue_id = ?", data.IssueId).Joins("JOIN user_roles ON user_roles.project_id = projects.project_id AND user_roles.user_id = ?", userId).Find(&issue).Count(&count)

	if count == 0 {
		return ut.GetErrorResponse(ct.ACTION_NOT_AUTHORIZED)
	} else {
		DB.Preload("Sprint").Preload("Project").Preload("Creator").Preload("AssignedTo").Where("issue_id", data.IssueId).Find(&issue)

		res := sk.IssueInfoResp{IssueId: issue.IssueId, Name: issue.Title, Type: issue.Type, SprintId: issue.SprintRef, SprintName: issue.Sprint.SprintName, ProjectId: issue.ProjectRef, ProjectName: issue.Project.ProjectName, Description: issue.Description, CreatorId: issue.CreatedBy, CreatorName: issue.Creator.Username, AssigneeId: issue.AssigneeId, AssigneeName: issue.AssignedTo.Username, CreatedAt: issue.CreatedAt, Status: issue.Status}
		return ut.GetSuccessResponse("", res)
	}
}

func DeleteIssue(data sk.IssueBaseReq, userId uint) gin.H {
	var count int64
	DB.Joins("JOIN projects ON projects.project_id = issues.project_ref AND issues.issue_id = ?", data.IssueId).Joins("JOIN user_roles ON user_roles.project_id = projects.project_id AND user_roles.user_id = ? AND user_roles.role_id = 1", userId).Find(&md.Issue{}).Count(&count)
	//DB.Where("user_id = ? AND project_id = ? AND role_id = 1", data.UserId, data.ProjectId).Find(&userRole).Count(&count)
	if count == 0 {
		return ut.GetErrorResponse(ct.ACTION_NOT_AUTHORIZED)
	}

	DB.Where("issue_id = ?", data.IssueId).Delete(&md.Issue{})
	//DB.Where("project_id=?", data.IssueId).Delete(&md.UserRole{})
	return ut.GetSuccessResponse(ct.ISSUE_DELETE_SUCCESS, "")
}

func UpdateIssue(req sk.UpdateIssueReq, userId uint) gin.H {
	var count int64
	var userRole md.UserRole
	DB.Where("user_id = ? AND project_id = ?", userId, req.ProjectId).Find(&userRole).Count(&count)

	if count == 0 {
		return ut.GetErrorResponse(ct.ACTION_NOT_AUTHORIZED)
	}

	DB.Where("sprint_id = ? AND project_ref = ?", req.SprintId, req.ProjectId).Find(&md.Sprint{}).Count(&count)

	if count == 0 {
		return ut.GetErrorResponse(ct.ACTION_NOT_AUTHORIZED)
	}

	var issue md.Issue
	DB.Model(&issue).Where("issue_id = ?", req.IssueId).Updates(md.Issue{Title: req.IssueTitle, Type: req.IssueType, Description: req.IssueText, Status: req.Status, SprintRef: req.SprintId})
	return ut.GetSuccessResponse(ct.ISSUE_UPDATE_SUCCESS, gin.H{})
}

func UpdateIssueStatus(req sk.UpdateIssueStatusReq, userId uint) gin.H {
	var count int64
	var userRole md.UserRole
	DB.Where("user_id = ? AND project_id = ?", userId, req.ProjectId).Find(&userRole).Count(&count)

	if count == 0 {
		return ut.GetErrorResponse(ct.ACTION_NOT_AUTHORIZED)
	}

	var issue md.Issue
	DB.Model(&issue).Where("issue_id = ?", req.IssueId).Updates(md.Issue{Status: req.Status})
	return ut.GetSuccessResponse(ct.ISSUE_UPDATE_SUCCESS, gin.H{})
}

func ChangeSprint(req sk.IssueTransfer, userId uint) gin.H {
	sprint_to := req.Sprint_to
	sprint_from := req.Sprint_from

	var count int64
	//DB.Joins("JOIN emails ON emails.user_id = users.id AND emails.email = ?", "jinzhu@example.org").Joins("JOIN credit_cards ON credit_cards.user_id = users.id").Where("credit_cards.number = ?", "411111111111").Find(&user)

	if count == 0 {
		return ut.GetErrorResponse(ct.ACTION_NOT_AUTHORIZED)
	}

	//SELECT sprints.sprint_id,sprints.sprint_name,projects.project_name FROM sprints JOIN projects ON projects.project_id = sprints.project_ref AND sprints.sprint_id=3 JOIN user_roles ON user_roles.project_id=projects.project_id AND user_roles.user_id=1
	DB.Joins("JOIN projects ON projects.project_id = sprints.project_ref AND sprints.sprint_id = ?", sprint_from).Joins("JOIN user_roles ON user_roles.project_id = projects.project_id AND users.user_id = ?", userId).Find(&md.Sprint{}).Count(&count)

	if count == 0 {
		return ut.GetErrorResponse(ct.ACTION_NOT_AUTHORIZED)
	}

	DB.Model(&md.Issue{}).Where("issue_id = ?", req.IssueId).Updates(md.Issue{SprintRef: sprint_to})
	return ut.GetSuccessResponse(ct.ISSUE_UPDATE_SUCCESS, gin.H{})
}

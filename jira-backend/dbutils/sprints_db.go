package dbutils

import (
	"fmt"
	ct "jira-backend/constants"
	md "jira-backend/models"
	sk "jira-backend/skeletons"
	ut "jira-backend/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

func AddSprint(data sk.CreateSprintReq) gin.H {
	sprint_name := strings.TrimSpace(data.SprintName)
	//sprint_start_date := strings.TrimSpace(data.StartDate)
	var userRole md.UserRole
	count := int64(0)

	db.Where("user_id = ? AND project_id = ? AND role_id = ?", data.UserId, data.ProjectId, ct.Owner).Find(&userRole).Count(&count)
	fmt.Println(data)
	if count == 0 {
		return ut.GetErrorResponse(ct.ACTION_NOT_AUTHORIZED)
	}

	db.Where("project_ref = ? AND sprint_name = ?", data.ProjectId, data.SprintName).Find(&md.Sprint{}).Count(&count)

	if count != 0 {
		return ut.GetErrorResponse(ct.SPRINT_ALREADY_EXISTS)
	}

	startdate, err1 := ut.DateFormat(data.StartDate)

	if err1 != nil {
		return ut.GetErrorResponse(ct.WRONG_DATE_FORMAT)
	}
	enddate, err2 := ut.DateFormat(data.EndDate)
	if err2 != nil {
		return ut.GetErrorResponse(ct.WRONG_DATE_FORMAT)
	}

	sprint := md.Sprint{SprintName: sprint_name, StartDate: startdate, EndDate: enddate, ProjectRef: data.ProjectId}
	db.Create(&sprint)
	resp := sk.CreateSprintResp{SprintId: sprint.SprintId, SprintName: sprint.SprintName}
	return ut.GetSuccessResponse("", resp)
}

func GetSprintInfo(data sk.SprintInfoReq) gin.H {
	var sprint md.Sprint
	count := int64(0)
	db.Preload("Project").Where("sprint_id = ?", data.SprintId).Find(&sprint).Count(&count)
	if count == 0 {
		return ut.GetErrorResponse(ct.NOT_VALID_SPRINT)
	}

	db.Where("project_id = ? AND user_id = ?", sprint.Project.ProjectId, data.UserId).Find(&md.UserRole{}).Count(&count)
	if count == 0 {
		return ut.GetErrorResponse(ct.ACTION_NOT_AUTHORIZED)
	}

	resp := sk.SprintInfoResp{SprintId: sprint.SprintId, Name: sprint.SprintName, CreatedAt: sprint.CreatedAt, StartDate: sprint.StartDate, EndDate: sprint.EndDate}
	return ut.GetSuccessResponse("", resp)
}

func GetSprintList(data sk.SprintListReq) gin.H {
	var sprintlist []md.Sprint
	count := int64(0)

	db.Where("user_id = ? AND project_id = ?", data.UserId, data.ProjectId).Find(&md.UserRole{}).Count(&count)
	if count == 0 {
		return ut.GetErrorResponse(ct.ACTION_NOT_AUTHORIZED)
	}

	db.Where("project_ref = ? ", data.ProjectId).Find(&sprintlist)
	res := make([]sk.SprintEntry, 0)
	for _, x := range sprintlist {
		res = append(res, sk.SprintEntry{Name: x.SprintName, Id: x.SprintId, StartDate: x.StartDate, EndDate: x.EndDate, CreatedAt: x.CreatedAt, ProjectId: x.ProjectRef})
	}

	return ut.GetSuccessResponse("", sk.SprintListResp{Sprints: res})
}

func DeleteSprint(data sk.SprintDeleteReq) gin.H {
	var sprint md.Sprint
	count := int64(0)
	db.Preload("Project").Where("sprint_id = ?", data.SprintId).Find(&sprint)

	db.Where("project_id = ? AND user_id = ?", sprint.Project.ProjectId, data.UserId).Find(&md.UserRole{}).Count(&count)

	if count == 0 {
		return ut.GetErrorResponse(ct.ACTION_NOT_AUTHORIZED)
	}

	db.Delete(&sprint)

	return ut.GetSuccessResponse(ct.SPRINT_DELETE_SUCCESS, "")
}

func GetIssuesForSprint(data sk.BaseSprintReq) gin.H {
	var sprint md.Sprint
	count := int64(0)
	db.Preload("Project").Where("sprint_id = ?", data.SprintId).Find(&sprint).Count(&count)
	if count == 0 {
		return ut.GetErrorResponse(ct.NOT_VALID_SPRINT)
	}

	db.Where("project_id = ? AND user_id = ?", sprint.Project.ProjectId, data.UserId).Find(&md.UserRole{}).Count(&count)

	if count == 0 {
		return ut.GetErrorResponse(ct.ACTION_NOT_AUTHORIZED)
	}
	var issues []md.Issue
	db.Where("sprint_ref = ?", data.SprintId).Find(&issues)

	var resp = make([]sk.IssueEntry, 0)
	for _, x := range issues {
		resp = append(resp, sk.IssueEntry{IssueId: x.IssueId, Name: x.Title, Status: x.Status, CreatedAt: x.CreatedAt})
	}

	return ut.GetSuccessResponse("", sk.IssueListResp{Issues: resp})
}

func GetSprintsForUser(data sk.BaseProjectIdReq) gin.H {
	var count int64
	db.Where("project_id = ? AND user_id = ?", data.ProjectId, data.UserId).Find(&md.UserRole{}).Count(&count)
	if count == 0 {
		return ut.GetErrorResponse(ct.ACTION_NOT_AUTHORIZED)
	}

	var sprints []md.Sprint

	db.Where("project_ref = ?").Find(&sprints)

	var resp []sk.ShortSprintEntry

	for _, x := range sprints {
		resp = append(resp, sk.ShortSprintEntry{SprintId: x.SprintId, SprintName: x.SprintName})
	}

	return ut.GetSuccessResponse("", sk.ShortSprintList{Sprints: resp})
}

func UpdateSprintInfo(data sk.SprintUpdateReq) gin.H {
	var count int64
	var sprint md.Sprint
	db.Joins("JOIN projects ON projects.project_id = sprints.project_ref AND sprints.sprint_id = ?", data.SprintId).Joins("JOIN user_roles ON user_roles.project_id = projects.project_id AND user_roles.user_id = ?", data.UserId).Find(&sprint).Count(&count)

	if count == 0 {
		return ut.GetErrorResponse(ct.ACTION_NOT_AUTHORIZED)
	}

	startdate, err1 := ut.DateFormat(data.StartDate)

	if err1 != nil {
		return ut.GetErrorResponse(ct.WRONG_DATE_FORMAT)
	}
	enddate, err2 := ut.DateFormat(data.EndDate)
	if err2 != nil {
		return ut.GetErrorResponse(ct.WRONG_DATE_FORMAT)
	}

	db.Model(&md.Sprint{}).Where("sprint_id = ?", data.SprintId).Updates(md.Sprint{SprintName: data.SprintName, StartDate: startdate, EndDate: enddate})
	return ut.GetSuccessResponse(ct.SPRINT_UPDATE_SUCCESS, "")
}

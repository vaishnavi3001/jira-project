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

func CreateProject(data sk.CreateProjectReq, userId uint) gin.H {
	project_name := strings.TrimSpace(data.Name)
	project := md.Project{ProjectName: project_name}
	var count int64
	DB.Where("project_name = ?", project_name).Find(&md.Project{}).Count(&count)
	fmt.Println(count)
	if count > 0 {
		return ut.GetErrorResponse(ct.PROJECT_ALREADY_EXISTS)
	}

	DB.Create(&project)
	//the one who creates the project will be owner of the project
	user_role := md.UserRole{UserId: userId, RoleId: ct.Owner, ProjectId: project.ProjectId}
	DB.Create(&user_role)

	res := sk.CreateProjectResp{ProjectName: project.ProjectName, ProjectId: project.ProjectId, CreatedAt: project.CreatedAt}
	return ut.GetSuccessResponse("", res)
}

func GetProjectInfo(data sk.ProjectInfoReq, userId uint) gin.H {
	var res sk.ProjectInfoResp
	var count int64
	var userRole md.UserRole
	DB.Where("user_id=? AND project_id=?", userId, data.ProjectId).Find(&userRole).Count(&count)
	if count == 0 {
		return ut.GetErrorResponse(ct.USER_DOESNT_EXIST)
	} else {
		DB.Preload("User").Where("role_id=1 AND project_id=?", data.ProjectId).Find(&userRole)
		var project md.Project
		DB.Where("project_id", data.ProjectId).Find(&project)
		res = sk.ProjectInfoResp{ProjectId: project.ProjectId, Name: project.ProjectName, CreatedAt: project.CreatedAt, OwnerUName: userRole.User.Username, OwnerFName: userRole.User.Firstname, OwnerLName: userRole.User.Lastname, OwnerId: userRole.User.UserId}
		return ut.GetSuccessResponse("", res)
	}
}

func GetProjectList(userId uint) gin.H {

	var userRoles []md.UserRole
	var count int64
	DB.Preload("Project").Where("user_id = ?", userId).Find(&userRoles).Count(&count)

	if count == 0 {
		return ut.GetErrorResponse(ct.USER_DOESNT_EXIST)
	}
	list := make([]sk.ProjectEntry, 0)
	for _, x := range userRoles {
		list = append(list, sk.ProjectEntry{Name: x.Project.ProjectName, Id: x.Project.ProjectId, CreatedAt: x.Project.CreatedAt, UserRole: x.RoleId})
	}

	res := sk.ProjectListResp{Projects: list}
	return ut.GetSuccessResponse("", res)
}

func DeleteProject(data sk.BaseProjectIdReq, userId uint) gin.H {
	var count int64
	var userRole md.UserRole
	DB.Where("user_id = ? AND project_id = ? AND role_id = 1", userId, data.ProjectId).Find(&userRole).Count(&count)
	if count == 0 {
		return ut.GetErrorResponse(ct.ACTION_NOT_AUTHORIZED)
	}

	project := md.Project{ProjectId: data.ProjectId}
	DB.Delete(&project)
	DB.Where("project_id=?", data.ProjectId).Delete(&md.UserRole{})
	return ut.GetSuccessResponse(ct.PROJECT_DELETE_SUCCESS, "")
}

func ListMembers(data sk.BaseProjectIdReq, userId uint) gin.H {
	var count int64
	var userRoles []md.UserRole
	DB.Where("user_id = ? AND project_id = ?", userId, data.ProjectId).Find(&md.UserRole{}).Count(&count)

	if count == 0 {
		return ut.GetErrorResponse(ct.ACTION_NOT_AUTHORIZED)
	}

	DB.Preload("User").Where("project_id = ?", data.ProjectId).Find(&userRoles)
	list := make([]sk.ProjectMembersResp, 0)

	for _, x := range userRoles {
		list = append(list, sk.ProjectMembersResp{UserId: x.UserId, FirstName: x.User.Firstname, LastName: x.User.Lastname, UserRole: x.RoleId})
	}

	return ut.GetSuccessResponse("", sk.ProjectMembersListResp{Members: list})
}

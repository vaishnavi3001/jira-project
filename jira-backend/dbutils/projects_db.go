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

func CreateProject(data sk.CreateProjectReq) gin.H {
	project_name := strings.TrimSpace(data.Name)
	project := md.Project{ProjectName: project_name}

	var count int64
	db.Where("project_name = ?", project_name).Find(&md.Project{}).Count(&count)
	fmt.Println(count)
	if count > 0 {
		return ut.GetErrorResponse(ct.PROJECT_ALREADY_EXISTS)
	}

	db.Create(&project)
	//the one who creates the project will be owner of the project
	user_role := md.UserRole{UserId: data.UserId, RoleId: ct.Owner, ProjectId: project.ProjectId}
	db.Create(&user_role)

	res := sk.CreateProjectResp{ProjectName: project.ProjectName, ProjectId: project.ProjectId, CreatedAt: project.CreatedAt}
	return ut.GetSuccessResponse("", res)
}

func GetProjectInfo(data sk.ProjectInfoReq) gin.H {
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

func GetProjectList(data sk.UsersBaseReq) gin.H {

	var userRoles []md.UserRole
	db.Preload("Project").Where("user_id = ?", data.UserId).Find(&userRoles)

	var list []sk.ProjectEntry
	for _, x := range userRoles {
		list = append(list, sk.ProjectEntry{Name: x.Project.ProjectName, Id: x.Project.ProjectId, CreatedAt: x.Project.CreatedAt, UserRole: x.RoleId})
	}

	res := sk.ProjectListResp{Projects: list}
	return ut.GetSuccessResponse("", res)
}

func DeleteProject(data sk.BaseProjectIdReq) gin.H {
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

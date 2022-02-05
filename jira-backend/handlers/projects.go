package handlers

//list project
//new project
//delete project
//add member to project
//transfer ownership
import (
	"jira-backend/constants"
	"jira-backend/models"
	"jira-backend/skeletons"
	"jira-backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateProject(c *gin.Context) {

	db, exists := c.Keys["db"].(*gorm.DB)

	if !exists {
		c.JSON(http.StatusInternalServerError, utils.GetResponse(false, "Something went wrong", ""))
		c.AbortWithStatus(http.StatusInternalServerError)
	}

	var createProjectReq skeletons.CreateProjectReq
	if err := c.BindJSON(&createProjectReq); err != nil {

		c.JSON(http.StatusBadRequest, utils.GetResponse(false, "Could not parse the request", ""))
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	project := models.Project{Name: createProjectReq.Name, OwnerId: createProjectReq.UserId}
	db.Create(&project)

	//the one who creates the project will be owner of the project
	user_role := models.UserRole{UserId: createProjectReq.UserId, RoleId: constants.Owner, ProjectId: project.ProjectId}
	db.Create(&user_role)
	c.JSON(http.StatusOK, utils.GetResponse(true, "Project Created Successfully", skeletons.CreateProjectResp{ProjectName: project.Name, ProjectId: project.ProjectId}))
}

func ListProjects(c *gin.Context) {
	db, exists := c.Keys["db"].(*gorm.DB)
	if !exists {
		c.JSON(http.StatusInternalServerError, utils.GetResponse(false, "Something went wrong", ""))
		c.AbortWithStatus(http.StatusInternalServerError)
	}
	var req skeletons.UserIdReq
	if err := c.BindJSON(&req); err != nil {

		c.JSON(http.StatusBadRequest, utils.GetResponse(false, "Could not parse the request", ""))
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	var userRoles []models.UserRole
	db.Preload("Project").Where("user_id = ?", req.UserId).Find(&userRoles)

	var res []skeletons.ProjectEntry
	for _, x := range userRoles {
		res = append(res, skeletons.ProjectEntry{Name: x.Project.Name, Id: x.Project.ProjectId, CreatedAt: x.Project.CreatedAt, UserRole: x.RoleId})
	}

	c.JSON(http.StatusOK, utils.GetResponse(true, "", skeletons.ProjectListResp{Projects: res}))

}

func GetProjectInfo(c *gin.Context) {
	db, exists := c.Keys["db"].(*gorm.DB)

	if !exists {
		c.JSON(http.StatusInternalServerError, utils.GetResponse(false, "Something went wrong", ""))
		c.AbortWithStatus(http.StatusInternalServerError)
	}

	var req skeletons.ProjectInfoReq
	if err := c.BindJSON(&req); err != nil {

		c.JSON(http.StatusBadRequest, utils.GetResponse(false, "Could not parse the request", ""))
		c.AbortWithStatus(http.StatusBadRequest)
	}

	var result int64
	var userRole models.UserRole
	db.Where("user_id=? AND project_id=?", req.UserId, req.ProjectId).Find(&userRole).Count(&result)
	if result > 0 {
		c.JSON(http.StatusUnauthorized, utils.GetResponse(false, "User does not belong to this project", ""))
	} else {
		var project models.Project
		db.Preload("User").Where("project_id", req.ProjectId).Find(&project)
		resBody := skeletons.ProjectInfoResp{ProjectId: project.ProjectId, OwnerName: project.User.Username, Name: project.Name, CreatedAt: project.CreatedAt}
		c.JSON(http.StatusOK, utils.GetResponse(true, "", resBody))
	}

}

func DeleteProject(c *gin.Context) {
	db, exists := c.Keys["db"].(*gorm.DB)
	if !exists {
		c.JSON(http.StatusInternalServerError, utils.GetResponse(false, "Something went wrong", ""))
		c.AbortWithStatus(http.StatusInternalServerError)
	}

}

//delete project
//transfer the ownership
//add member to the project

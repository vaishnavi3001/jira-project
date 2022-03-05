package handlers

//list project
//new sprint
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

func CreateSprint(c *gin.Context) {

	db, exists := c.Keys["db"].(*gorm.DB)

	if !exists {
		c.JSON(http.StatusInternalServerError, utils.GetResponse(false, "Something went wrong", ""))
		c.AbortWithStatus(http.StatusInternalServerError)
	}
	
	var createSprintReq skeletons.CreateSprintReq
	if err := c.BindJSON(&createSprintReq); err != nil {
		c.JSON(http.StatusBadRequest, utils.GetResponse(false, "Could not parse the request", ""))
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	sprint := models.Sprint{SprintName: createSprintReq.Name, ProjectId: createSprintReq.ProjectId, StartDate: createSprintReq.StartDate, EndDate: createSprintReq.EndDate}
	db.Create(&sprint)

	//the owner of the project under which the sprint is created is the owner of the sprint
	var project models.Project
	db.Preload("User").Where("project_id", createSprintReq.ProjectId).Find(&project)
	user_role := models.UserRole{UserId: project.OwnerId, RoleId: constants.Owner}
	db.Create(&user_role)
	c.JSON(http.StatusOK, utils.GetResponse(true, "Sprint Created Successfully", skeletons.CreateSprintResp{SprintName: sprint.SprintName, SprintId: sprint.SprintId}))
}

func ListSprints(c *gin.Context) {
	db, exists := c.Keys["db"].(*gorm.DB)
	if !exists {
		c.JSON(http.StatusInternalServerError, utils.GetResponse(false, "Something went wrong", ""))
		c.AbortWithStatus(http.StatusInternalServerError)
	}
	var req skeletons.SprintListReq
	if err := c.BindJSON (&req); err != nil {

		c.JSON(http.StatusBadRequest, utils.GetResponse(false, "Could not parse the request", ""))
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	var res []skeletons.SprintEntry
	db.Preload("Sprint").Where("project_id = ?", req.ProjectId).Find(&res)
	c.JSON(http.StatusOK, utils.GetResponse(true, "", skeletons.SprintListResp{Sprints: res}))

}

func GetSprintInfo(c *gin.Context) {
	db, exists := c.Keys["db"].(*gorm.DB)

	if !exists {
		c.JSON(http.StatusInternalServerError, utils.GetResponse(false, "Something went wrong", ""))
		c.AbortWithStatus(http.StatusInternalServerError)
	}

	var req skeletons.SprintInfoReq
	if err := c.BindJSON(&req); err != nil {

		c.JSON(http.StatusBadRequest, utils.GetResponse(false, "Could not parse the request", ""))
		c.AbortWithStatus(http.StatusBadRequest)
	}	
	var sprint models.Sprint
	db.Preload("Sprint").Where("sprint_id", req.SprintId).Find(&sprint)
	resBody := skeletons.SprintInfoResp{SprintId: sprint.SprintId, Name: sprint.SprintName,   CreatedAt: sprint.CreatedAt }
	c.JSON(http.StatusOK, utils.GetResponse(true, "", resBody))
	

}

// func UpdateSprintInfo(c *gin.Context) {
// 	db, exists := c.Keys["db"].(*gorm.DB)

// 	if !exists {
// 		c.JSON(http.StatusInternalServerError, utils.GetResponse(false, "Something went wrong", ""))
// 		c.AbortWithStatus(http.StatusInternalServerError)
// 	}

// 	var req skeletons.SprintInfoReq
// 	if err := c.BindJSON(&req); err != nil {

// 		c.JSON(http.StatusBadRequest, utils.GetResponse(false, "Could not parse the request", ""))
// 		c.AbortWithStatus(http.StatusBadRequest)
// 	}	

// 	var sprint models.Sprint
// 	db.Preload("Sprint").Where("sprint_id", req.SprintId).Find(&sprint)
// 	resBody := skeletons.SprintInfoResp{SprintId: sprint.SprintId, Name: sprint.SprintName,   CreatedAt: sprint.CreatedAt }
// 	c.JSON(http.StatusOK, utils.GetResponse(true, "", resBody))
	

// }

func DeleteSprint(c *gin.Context) {
	db, exists := c.Keys["db"].(*gorm.DB)
	if !exists {
		c.JSON(http.StatusInternalServerError, utils.GetResponse(false, "Something went wrong", ""))
		c.AbortWithStatus(http.StatusInternalServerError)
	}

	var req skeletons.SprintInfoReq
	if err := c.BindJSON(&req); err != nil {

		c.JSON(http.StatusBadRequest, utils.GetResponse(false, "Could not parse the request", ""))
		c.AbortWithStatus(http.StatusBadRequest)
	}

	// var result int64
	
	db.Preload("Sprint").Where("sprint_id", req.SprintId).Update("isDeleted", true)
	c.JSON(http.StatusOK, utils.GetResponse(true, "Sprint deleted successfully", ""))
	// sprint = &models.Sprint{
	// 	sprint.IsDeleted: 1
	// }

	// db.Where("user_id=? AND project_id=? AND role_id=1", req.UserId, req.ProjectId).Update("isDeleted", true)
	// if result == 0 {
	// 	c.JSON(http.StatusUnauthorized, utils.GetResponse(false, "You do not belong to this project/ You are not the owner of the project", ""))
	// 	c.AbortWithStatus(http.StatusInternalServerError)
	// } else {
	// 	var project models.Project
	// 	db.Delete(&project, req.ProjectId)
	// 	c.JSON(http.StatusOK, utils.GetResponse(true, "Project deleted successfully", ""))
	// }
}


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

	sprint := models.Sprint{Name: createSprintReq.Name, OwnerId: createSprintReq.UserId}
	db.Create(&sprint)

	//the one who creates the project will be owner of the project
	user_role := models.UserRole{UserId: createSprintReq.UserId, RoleId: constants.Owner, SprintId: sprint.SprintId}
	db.Create(&user_role)
	c.JSON(http.StatusOK, utils.GetResponse(true, "Sprint Created Successfully", skeletons.CreateSprintResp{SprintName: sprint.Name, SprintId: sprint.SprintId}))
}

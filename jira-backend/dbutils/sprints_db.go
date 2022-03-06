package dbutils

import (
	sk "jira-backend/skeletons"

	"github.com/gin-gonic/gin"
)

func AddSprint(data sk.CreateSprintReq) {
	/*sprint_name := strings.TrimSpace(data.SprintName)
	sprint_start_date := strings.TrimSpace(data.StartDate)
	count := 0
	db.Where("spring_name = ? & project_id = ?", sprint_name, data.ProjectId).Find(&count)

	if count != 0 {
		return utils.GetErrorResponse(ct.SPRINT_ALREADY_EXISTS)
	}

	//mm-dd-yyyy
	sprint := md.Sprint{SprintName: sprint_name, StartDate: data.StartDate, EndDate: data.EndDate}
	db.Create()
	return*/
}

func GetSprintInfo(c *gin.Context) {

}

func GetSprintList(c *gin.Context) {

}

func DeleteSprint(c *gin.Context) {

}

func GetIssuesForSprint(c *gin.Context) {

}

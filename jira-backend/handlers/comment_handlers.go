package handlers

import (
	dt "jira-backend/dbutils"
	sk "jira-backend/skeletons"
	ut "jira-backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddComment(c *gin.Context) {
	user_id := ut.GetUserIdFromContext(c)
	var req sk.CommentAddReq
	if err := c.BindJSON(&req); err != nil {
		ut.ThrowBadRequest(c)
		return
	}

	c.JSON(http.StatusOK, dt.AddComments(req, user_id))

}

func EditComment(c *gin.Context) {

}
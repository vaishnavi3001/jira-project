package dbutils

import (
	// ct "jira-backend/constants"

	md "jira-backend/models"
	sk "jira-backend/skeletons"
	ut "jira-backend/utils"

	"github.com/gin-gonic/gin"
)

func AddComments(data sk.CommentAddReq, userId uint) gin.H {

	comment := md.Comment{CommentText: data.CommentText, UserId: userId, IssueId: data.IssueId}
	DB.Create(&comment)

	resp := sk.AddIssueCommentResp{Comment: data.CommentText}
	return ut.GetSuccessResponse("", resp)
}

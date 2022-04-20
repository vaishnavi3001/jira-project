package dbutils

import (
	md "jira-backend/models"
	sk "jira-backend/skeletons"
	ut "jira-backend/utils"

	"github.com/gin-gonic/gin"
)

func AddComments(data sk.CommentAddReq, userId uint) gin.H {

	comment := md.Comment{CommentText: data.CommentText, UserRef: userId, IssueRef: data.IssueId}
	DB.Create(&comment)

	resp := sk.AddIssueCommentResp{Comment: data.CommentText}
	return ut.GetSuccessResponse("", resp)
}

func ViewComments(data sk.CommentViewReq, userId uint) gin.H {

	var commentlist []md.Comment

	DB.Where("issue_id = ? ", data.IssueId).Find(&commentlist)
	res := make([]sk.AddIssueCommentResp, 0)
	for _, x := range commentlist {
		res = append(res, sk.AddIssueCommentResp{CommentId: x.CommentId, Comment: x.CommentText})
	}

	resp := sk.CommentListResp{Comments: res}
	return ut.GetSuccessResponse("", resp)
}

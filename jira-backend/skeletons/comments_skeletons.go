package skeletons

type CommentAddReq struct {
	IssueId     uint   `json:"issue_id"`
	CommentText string `json:"comment"`
}

type AddIssueCommentResp struct {
	Comment string `json:"comment"`
}

type CommentEditReq struct {
	CommentId   uint
	CommentText string
}

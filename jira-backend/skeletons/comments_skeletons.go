package skeletons

type CommentAddReq struct {
	IssueId     uint   `json:"issue_id"`
	CommentText string `json:"comment"`
}

type AddIssueCommentResp struct {
	CommentId uint   `json:"comment_id"`
	Comment   string `json:"comment"`
}

type CommentEditReq struct {
	CommentId   uint
	CommentText string
}

type CommentViewReq struct {
	IssueId uint `json:"issue_id"`
}

type CommentListResp struct {
	Comments []AddIssueCommentResp `json:"comments"`
}

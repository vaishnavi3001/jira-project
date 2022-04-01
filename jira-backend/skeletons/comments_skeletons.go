package skeletons

type CommentAddReq struct {
	IssueId     uint
	CommentText string
}

type CommentEditReq struct {
	CommentId   uint
	CommentText string
}

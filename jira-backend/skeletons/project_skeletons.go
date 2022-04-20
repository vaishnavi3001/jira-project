package skeletons

import "time"

type BaseProjectIdReq struct {
	ProjectId uint `json:"project_id"`
}
type CreateProjectReq struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Key         string `json:"key"`
}

type ProjectEntry struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Key         string    `json:"key"`
	Id          uint      `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UserRole    uint      `json:"user_role"`
}

type ProjectListResp struct {
	Projects []ProjectEntry `json:"projects"`
}

type CreateProjectResp struct {
	ProjectName        string    `json:"project_name"`
	ProjectKey         string    `json:"project_key"`
	ProjectDescription string    `json:"project_desc"`
	ProjectId          uint      `json:"project_id"`
	CreatedAt          time.Time `json:"created_at"`
}

type ProjectInfoReq struct {
	ProjectId uint `json:"project_id"`
}

type ProjectInfoResp struct {
	ProjectId   uint      `json:"project_id"`
	Name        string    `json:"project_name"`
	Description string    `json:"project_desc"`
	Key         string    `json:"project_key"`
	OwnerUName  string    `json:"owner_uname"`
	OwnerId     uint      `json:"owner_id"`
	OwnerFName  string    `json:"owner_fname"`
	OwnerLName  string    `json:"owner_lname"`
	CreatedAt   time.Time `json:"created_at"`
}

type ProjectUsers struct {
	Users []UserEntry `json:"users"`
}

type ProjectMembersResp struct {
	UserId    uint   `json:"user_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	UserRole  uint   `json:"user_role"`
}

type ProjectMembersListResp struct {
	Members []ProjectMembersResp `json:"members"`
}

type Stats struct {
	MemberCount  int64 `json:"member_count"`
	IssueCount   int64 `json:"issue_count"`
	CommentCount int64 `json:"comment_count"`
}

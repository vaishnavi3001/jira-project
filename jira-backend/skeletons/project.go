package skeletons

import "time"

type CreateProjectReq struct {
	Name   string `json:"name"`
	UserId uint   `json:"user_id"`
}

type UserIdReq struct {
	UserId int `json:"user_id"`
}

type ProjectEntry struct {
	Name      string    `json:"name"`
	Id        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UserRole  uint      `json:"user_role"`
}

type ProjectListResp struct {
	Projects []ProjectEntry `json:"projects"`
}

type CreateProjectResp struct {
	ProjectName string `json:"project_name"`
	ProjectId   uint   `json:"project_id"`
}

type ProjectInfoReq struct {
	ProjectId uint `json:"project_id"`
	UserId    uint `json:"user_id"`
}

type UserEntry struct {
	UserId   uint `json:"user_id"`
	UserName uint `json:"user_name"`
	UserRole uint `json:"user_role"`
}

type ProjectInfoResp struct {
	ProjectId  uint      `json:"project_id"`
	OwnerUName string    `json:"owner_uname"`
	OwnerId    uint      `json:"owner_id"`
	OwnerName  string    `json:"owner_name"`
	Name       string    `json:"project_name"`
	CreatedAt  time.Time `json:"created_at"`
}

type ProjectUsers struct {
	Users []UserEntry `json:"users"`
}

package skeletons

import "time"

type CreateSprintReq struct {
	Name   string `json:"name"`
	UserId uint   `json:"user_id"`
}

type UserIdReq struct {
	UserId int `json:"user_id"`
}

type SprintEntry struct {
	Name      string    `json:"name"`
	Id        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UserRole  uint      `json:"user_role"`
}

type SprintListResp struct {
	Projects []SprintEntry `json:"sprints"`
}

type CreateSprintResp struct {
	ProjectName string `json:"sprint_name"`
	ProjectId   uint   `json:"sprint_id"`
}

type SprintInfoReq struct {
	ProjectId uint `json:"sprint_id"`
	UserId    uint `json:"user_id"`
}

type UserEntry struct {
	UserId   uint `json:"user_id"`
	UserName uint `json:"user_name"`
	UserRole uint `json:"user_role"`
}

type SprintInfoResp struct {
	SprintId  uint      `json:"sprint_id"`
	OwnerUName string    `json:"owner_uname"`
	OwnerId    uint      `json:"owner_id"`
	OwnerName  string    `json:"owner_name"`
	Name       string    `json:"sprint_name"`
	CreatedAt  time.Time `json:"created_at"`
}

type SprintUsers struct {
	Users []UserEntry `json:"users"`
}

package skeletons

import (
	"time"
)

type BaseSprintReq struct {
	SprintId uint `json:"sprint_id"`
}

type CreateSprintReq struct {
	SprintName string `json:"sprint_name"`
	ProjectId  uint   `json:"project_id"`
	StartDate  string `json:"start_date"`
	EndDate    string `json:"end_date"`
}

type SprintEntry struct {
	Name      string    `json:"name"`
	Id        uint      `json:"id"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
	CreatedAt time.Time `json:"created_at"`
	ProjectId uint      `json:"project_id"`
}
type SprintListReq struct {
	ProjectId uint `json:"project_id"`
}
type SprintListResp struct {
	Sprints []SprintEntry `json:"sprints"`
}

type CreateSprintResp struct {
	SprintName string `json:"sprint_name"`
	SprintId   uint   `json:"sprint_id"`
}

type SprintInfoReq struct {
	SprintId uint `json:"sprint_id"`
}

type SprintInfoResp struct {
	SprintId  uint      `json:"sprint_id"`
	Name      string    `json:"sprint_name"`
	CreatedAt time.Time `json:"created_at"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
}

type SprintUsers struct {
	Users []UserEntry `json:"users"`
}

type SprintDeleteReq struct {
	SprintId uint `json:"sprint_id"`
}

type ShortSprintEntry struct {
	SprintId   uint   `json:"sprint_id"`
	SprintName string `json:"sprint_name"`
}

type ShortSprintList struct {
	Sprints []ShortSprintEntry `json:"sprints"`
}

type SprintUpdateReq struct {
	SprintId   uint   `json:"sprint_id"`
	SprintName string `json:"sprint_name"`
	StartDate  string `json:"start_date"`
	EndDate    string `json:"end_date"`
}

package skeletons

import (
	"time"
)

type BaseSprintReq struct {
	SprintId uint `json:"sprint_id"`
}

type CreateSprintReq struct {
	SprintName string `json:"sprint_name" binding:"required"`
	ProjectId  uint   `json:"project_id" binding:"required"`
	StartDate  string `json:"start_date" binding:"required"`
	EndDate    string `json:"end_date" binding:"required"`
}

type SprintEntry struct {
	Name      string    `json:"name" binding:"required"`
	Id        uint      `json:"id" binding:"required"`
	StartDate time.Time `json:"start_date" binding:"required"`
	EndDate   time.Time `json:"end_date" binding:"required"`
	CreatedAt time.Time `json:"created_at" binding:"required"`
	ProjectId uint      `json:"project_id" binding:"required"`
}
type SprintListReq struct {
	ProjectId uint `json:"project_id" binding:"required"`
}
type SprintListResp struct {
	Sprints []SprintEntry `json:"sprints"`
}

type CreateSprintResp struct {
	SprintName string `json:"sprint_name"`
	SprintId   uint   `json:"sprint_id"`
}

type SprintInfoReq struct {
	SprintId uint `json:"sprint_id" binding:"required"`
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
	SprintId   uint   `json:"sprint_id" binding:"required"`
	SprintName string `json:"sprint_name" binding:"required"`
	StartDate  string `json:"start_date" binding:"required"`
	EndDate    string `json:"end_date" binding:"required"`
}

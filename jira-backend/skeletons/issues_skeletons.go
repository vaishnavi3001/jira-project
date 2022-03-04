package skeletons

import "time"

type IssueBaseReq struct {
	UserId  uint `json:"user_id"`
	IssueId uint `json:"issue_id"`
}

type AddIssueReq struct {
	UserId     uint   `json:"user_id"`
	IssueTitle string `json:"issue_name"`
	IssueText  string `json:"issue_text"`
	IssueType  uint   `json:"issue_type"`
	Creator    uint   `json:"creator"`
	Assignee   uint   `json:"assignee"`
	SprintId   uint   `json:"sprint_id"`
	ProjectId  uint   `json:"project_id"`
}

type UpdateIssueReq struct {
	UserId     uint   `json:"user_id"`
	IssueId    uint   `json:"issue_id"`
	IssueTitle string `json:"issue_name"`
	IssueText  string `json:"issue_text"`
	IssueType  uint   `json:"issue_type"`
	Creator    uint   `json:"creator"`
	Assignee   uint   `json:"Assignee"`
	Status     uint   `json:"Status"`
	SprintId   uint   `json:"sprint_id"`
	ProjectId  uint   `json:"project_id"`
}

type BaseIssueResp struct {
	UserId    uint   `json:"user_id"`
	IssueId   uint   `json:"issue_id"`
	IssueName string `json:"issue_name"`
}

type IssueListReq struct {
	UserId    uint `json:"user_id"`
	ProjectId uint `json:"project_id"`
	SprintId  uint `json:"sprint_id"`
}

type IssueEntry struct {
	IssueId uint   `json:"issue_id"`
	Name    string `json:"title"`
	Status  uint   `json:"status"`
}

type IssueListResp struct {
	Issues []IssueEntry `json:"issues"`
}

/*IssueId     uint `gorm:"primaryKey;auto_increment;not_null"`
Status      uint `gorm:"default:1"`
Type        uint `gorm:"default:1"` //`epic/task/subtask/bug/`
Title       string
CreatedBy   uint //user ref
SprintRef   uint
AssigneeId  uint //user ref
ProjectRef  uint
IsDeleted   sql.NullBool
UpdatedAt   time.Time
Description string
CreatedAt   time.Time `gorm:"autoCreateTime:milli"`
Sprint      Sprint    `gorm:"foreignKey:SprintRef"`
AssignedTo  User      `gorm:"foreignKey:AssigneeId"`
Creator     User      `gorm:"foreignKey:CreatedBy"`
Project     Project   `gorm:"foreignKey:ProjectRef"`
DeletedAt   gorm.DeletedAt*/
type IssueEntryDetailed struct {
	IssueId     uint      `json:"issue_id"`
	Name        string    `json:"title"`
	Status      uint      `json:"status"`
	Description string    `json:"description"`
	AssignedTo  string    `json:"assignee"`
	CreatedBy   string    `json:"created_by"`
	CreatedAt   time.Time `json:"created_at"`
	ModifiedAt  time.Time `json:"modified_at"`
}

//sprint id, project id, issue_type, issue description, createdBy, AssignedTo, title

package skeletons

import "time"

type IssueBaseReq struct {
	UserId  uint `json:"user_id"`
	IssueId uint `json:"issue_id"`
}

type AddIssueReq struct {
	UserId     uint   `json:"user_id"`
	IssueTitle string `json:"issue_title"`
	IssueText  string `json:"issue_description"`
	IssueType  uint   `json:"issue_type"`
	Creator    uint   `json:"creator"`
	Assignee   uint   `json:"assignee"`
	SprintId   uint   `json:"sprint_id"`
	ProjectId  uint   `json:"project_id"`
}

type AddIssueResp struct {
	IssueId    uint      `json:"issue_id"`
	IssueTitle string    `json:"issue_title"`
	IssueText  string    `json:"issue_description"`
	IssueType  uint      `json:"issue_type"`
	Creator    string    `json:"creator_name"`
	Assignee   string    `json:"assignee_name"`
	CreatedAt  time.Time `json:"created_at"`
	Status     uint      `json:"issue_status"`
}
type UpdateIssueReq struct {
	UserId     uint   `json:"user_id"`
	IssueId    uint   `json:"issue_id"`
	IssueTitle string `json:"issue_title"`
	IssueText  string `json:"issue_description"`
	IssueType  uint   `json:"issue_type"`
	Creator    uint   `json:"creator"`
	Assignee   uint   `json:"assignee"`
	Status     uint   `json:"status"`
	SprintId   uint   `json:"sprint_id"`
	ProjectId  uint   `json:"project_id"`
}

type UpdateIssueStatusReq struct {
	IssueId   uint `json:"issue_id"`
	ProjectId uint `json:"project_id"`
	Status    uint `json:"status"`
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
	IssueId   uint      `json:"issue_id"`
	Name      string    `json:"title"`
	Status    uint      `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}

type IssueListResp struct {
	Issues []IssueEntry `json:"issues"`
}

type IssueTransfer struct {
	UserId      uint `json:"user_id"`
	IssueId     uint `json:"issue_id"`
	Sprint_from uint `json:"sprint_from"`
	Sprint_to   uint `json:"sprint_to"`
}

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

type IssueInfoResp struct {
	IssueId      uint      `json:"issue_id"`
	Name         string    `json:"title"`
	Type         uint      `json:"type"`
	SprintId     uint      `json:"sprint_id"`
	SprintName   string    `json:"sprint_name"`
	ProjectId    uint      `json:"project_id"`
	Description  string    `json:"description"`
	AssigneeId   uint      `json:"assignee_id"`
	AssigneeName string    `json:"assignee_name"`
	CreatorId    uint      `json:"creator_id"`
	CreatorName  string    `json:"creator_name"`
	CreatedAt    time.Time `json:"created_at"`
	ProjectName  string    `json:"project_name"`
	Status       uint      `json:"issue_status"`
}

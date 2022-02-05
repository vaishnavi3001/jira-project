package models

import (
	"time"

	"database/sql"
)

type Project struct {
	ProjectId uint `gorm:"primaryKey;auto_increment"`
	Name      string
	IsDeleted bool
	OwnerId   uint      //UserRoleId
	CreatedAt time.Time `gorm:"autoUpdateTime:milli"`
	UpdatedAt time.Time
	User      User `gorm:"foreignKey:OwnerId"`
}

type Issue struct {
	IssueId    uint
	Status     string
	Type       string //`epic/task/subtask/bug/`
	Title      string
	CreatedBy  uint //user ref
	SprintId   uint
	AssigneeId uint //user ref
	IsDeleted  sql.NullBool
	UpdatedAt  time.Time
	CreatedAt  time.Time `gorm:"autoUpdateTime:milli"`
	Sprint     Sprint    `gorm:"foreignKey:SprintId"`
	Creator    User      `gorm:"foreignKey:CreatedBy"`
	AssignedTo User      `gorm:"foreignKey:AssigneeId"`
}

type Sprint struct {
	ID        uint `gorm:"primaryKey;auto_increment;not_null"`
	Name      string
	ProjectId uint
	CreatedAt time.Time `gorm:"autoUpdateTime:milli"`
	UpdatedAt time.Time
	StartDate time.Time
	EndDate   time.Time
	Project   Project `gorm:"foreignKey:ProjectId"`
}

type User struct {
	UserId uint `gorm:"primaryKey;auto_increment;not_null"`
	//RoleId    uint
	Username  string
	Name      string
	IsDeleted sql.NullBool
	EmailId   string
	CreatedAt time.Time `gorm:"autoUpdateTime:milli"`
}

type UserRole struct {
	UserId     uint `gorm:"primaryKey"`
	RoleId     uint `gorm:"primaryKey"`
	ProjectId  uint `gorm:"primaryKey"`
	Membership uint
	Project    Project `gorm:"foreignKey:ProjectId"`
	User       User    `gorm:"foreignKey:UserId"`
	Role       Role    `gorm:"foreignKey:RoleId"`
}

type Role struct {
	RoleId       uint `gorm:"primaryKey"`
	RoleName     string
	PermissionId uint
	Permission   Permission `gorm:"foreignKey:PermissionId"`
}

type Permission struct {
	PermissionId   uint `gorm:"primaryKey"`
	Permissionname string
}

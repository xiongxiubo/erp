package model

import "time"

type User struct {
	BaseModel
	Username     string     `gorm:"size:64;uniqueIndex;not null" json:"username"`
	PasswordHash string     `gorm:"size:255;not null" json:"-"`
	Name         string     `gorm:"size:80;not null" json:"name"`
	Email        string     `gorm:"size:120" json:"email"`
	Phone        string     `gorm:"size:40" json:"phone"`
	DepartmentID uint       `gorm:"index" json:"departmentId"`
	Status       string     `gorm:"size:24;index;default:enabled" json:"status"`
	LastLoginAt  *time.Time `json:"lastLoginAt"`
	Roles        []Role     `gorm:"many2many:sys_user_roles" json:"roles"`
}

func (User) TableName() string { return "sys_users" }

type Role struct {
	BaseModel
	Code        string       `gorm:"size:80;uniqueIndex;not null" json:"code"`
	Name        string       `gorm:"size:80;not null" json:"name"`
	Description string       `gorm:"size:255" json:"description"`
	Status      string       `gorm:"size:24;index;default:enabled" json:"status"`
	Permissions []Permission `gorm:"many2many:sys_role_permissions" json:"permissions"`
	Menus       []Menu       `gorm:"many2many:sys_role_menus" json:"menus"`
}

func (Role) TableName() string { return "sys_roles" }

type Permission struct {
	BaseModel
	Code         string `gorm:"size:120;uniqueIndex;not null" json:"code"`
	Name         string `gorm:"size:120;not null" json:"name"`
	ResourceType string `gorm:"size:40" json:"resourceType"`
	Method       string `gorm:"size:16" json:"method"`
	Path         string `gorm:"size:255" json:"path"`
	Description  string `gorm:"size:255" json:"description"`
	Status       string `gorm:"size:24;index;default:enabled" json:"status"`
}

func (Permission) TableName() string { return "sys_permissions" }

type Menu struct {
	BaseModel
	ParentID       uint   `gorm:"index" json:"parentId"`
	Name           string `gorm:"size:80;not null" json:"name"`
	Path           string `gorm:"size:160" json:"path"`
	Component      string `gorm:"size:160" json:"component"`
	Icon           string `gorm:"size:80" json:"icon"`
	Sort           int    `gorm:"index" json:"sort"`
	PermissionCode string `gorm:"size:120;index" json:"permissionCode"`
	Visible        bool   `gorm:"default:true" json:"visible"`
	Status         string `gorm:"size:24;index;default:enabled" json:"status"`
}

func (Menu) TableName() string { return "sys_menus" }

type Department struct {
	BaseModel
	ParentID uint   `gorm:"index" json:"parentId"`
	Code     string `gorm:"size:80;uniqueIndex;not null" json:"code"`
	Name     string `gorm:"size:120;not null" json:"name"`
	Sort     int    `gorm:"index" json:"sort"`
	Status   string `gorm:"size:24;index;default:enabled" json:"status"`
}

func (Department) TableName() string { return "sys_departments" }

type OperationLog struct {
	BaseModel
	UserID   uint   `gorm:"index" json:"userId"`
	Module   string `gorm:"size:80;index" json:"module"`
	Action   string `gorm:"size:80;index" json:"action"`
	BizType  string `gorm:"size:80;index" json:"bizType"`
	BizID    uint   `gorm:"index" json:"bizId"`
	Method   string `gorm:"size:16" json:"method"`
	Path     string `gorm:"size:255" json:"path"`
	IP       string `gorm:"size:80" json:"ip"`
	Snapshot string `gorm:"type:text" json:"snapshot"`
}

func (OperationLog) TableName() string { return "sys_operation_logs" }

type LoginLog struct {
	BaseModel
	Username string `gorm:"size:80;index" json:"username"`
	UserID   uint   `gorm:"index" json:"userId"`
	IP       string `gorm:"size:80" json:"ip"`
	Success  bool   `json:"success"`
	Message  string `gorm:"size:255" json:"message"`
}

func (LoginLog) TableName() string { return "sys_login_logs" }

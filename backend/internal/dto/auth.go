package dto

import "time"

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	Token     string       `json:"token"`
	ExpiresAt time.Time    `json:"expiresAt"`
	User      UserResponse `json:"user"`
}

type UserResponse struct {
	ID           uint           `json:"id"`
	Username     string         `json:"username"`
	Name         string         `json:"name"`
	Email        string         `json:"email"`
	Phone        string         `json:"phone"`
	DepartmentID uint           `json:"departmentId"`
	Status       string         `json:"status"`
	Roles        []RoleResponse `json:"roles"`
}

type RoleResponse struct {
	ID   uint   `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
}

type MenuResponse struct {
	ID             uint           `json:"id"`
	ParentID       uint           `json:"parentId"`
	Name           string         `json:"name"`
	Path           string         `json:"path"`
	Icon           string         `json:"icon"`
	Sort           int            `json:"sort"`
	PermissionCode string         `json:"permissionCode"`
	Children       []MenuResponse `json:"children"`
}

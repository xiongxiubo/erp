package service

import (
	"erp/internal/dto"
	"erp/internal/model"
	"erp/internal/response"
	"erp/internal/utils"
	"time"

	"gorm.io/gorm"
)

type AuthService struct {
	db        *gorm.DB
	secret    string
	expiresIn time.Duration
}

func NewAuthService(db *gorm.DB, secret string, expiresIn time.Duration) *AuthService {
	return &AuthService{db: db, secret: secret, expiresIn: expiresIn}
}

func (s *AuthService) Login(req dto.LoginRequest, ip string) (dto.LoginResponse, error) {
	var user model.User
	if err := s.db.Preload("Roles").Where("username = ?", req.Username).First(&user).Error; err != nil {
		s.writeLoginLog(req.Username, 0, ip, false, "用户名或密码错误")
		return dto.LoginResponse{}, response.Unauthorized("用户名或密码错误")
	}
	if user.Status != model.StatusEnabled || !utils.VerifyPassword(user.PasswordHash, req.Password) {
		s.writeLoginLog(req.Username, user.ID, ip, false, "用户名或密码错误")
		return dto.LoginResponse{}, response.Unauthorized("用户名或密码错误")
	}
	token, err := utils.SignJWT(user.ID, user.Username, s.secret, s.expiresIn)
	if err != nil {
		return dto.LoginResponse{}, err
	}
	now := time.Now()
	s.db.Model(&user).Update("last_login_at", now)
	s.writeLoginLog(req.Username, user.ID, ip, true, "登录成功")
	return dto.LoginResponse{Token: token, ExpiresAt: now.Add(s.expiresIn), User: toUserResponse(user)}, nil
}

func (s *AuthService) Profile(userID uint) (dto.UserResponse, error) {
	var user model.User
	if err := s.db.Preload("Roles").First(&user, userID).Error; err != nil {
		return dto.UserResponse{}, response.NotFound("用户不存在")
	}
	return toUserResponse(user), nil
}

func (s *AuthService) Permissions(userID uint) ([]string, error) {
	var codes []string
	err := s.db.Table("sys_permissions").
		Select("distinct sys_permissions.code").
		Joins("join sys_role_permissions on sys_role_permissions.permission_id = sys_permissions.id").
		Joins("join sys_user_roles on sys_user_roles.role_id = sys_role_permissions.role_id").
		Where("sys_user_roles.user_id = ? and sys_permissions.status = ?", userID, model.StatusEnabled).
		Pluck("sys_permissions.code", &codes).Error
	return codes, err
}

func (s *AuthService) Menus(userID uint) ([]dto.MenuResponse, error) {
	var menus []model.Menu
	err := s.db.Table("sys_menus").
		Select("distinct sys_menus.*").
		Joins("join sys_role_menus on sys_role_menus.menu_id = sys_menus.id").
		Joins("join sys_user_roles on sys_user_roles.role_id = sys_role_menus.role_id").
		Where("sys_user_roles.user_id = ? and sys_menus.status = ? and sys_menus.visible = ?", userID, model.StatusEnabled, true).
		Order("sys_menus.sort asc, sys_menus.id asc").
		Find(&menus).Error
	if err != nil {
		return nil, err
	}
	return buildMenuTree(menus, 0), nil
}

func (s *AuthService) HasPermission(userID uint, code string) bool {
	if code == "" {
		return true
	}
	var count int64
	s.db.Table("sys_permissions").
		Joins("join sys_role_permissions on sys_role_permissions.permission_id = sys_permissions.id").
		Joins("join sys_user_roles on sys_user_roles.role_id = sys_role_permissions.role_id").
		Where("sys_user_roles.user_id = ? and sys_permissions.code = ? and sys_permissions.status = ?", userID, code, model.StatusEnabled).
		Count(&count)
	return count > 0
}

func (s *AuthService) writeLoginLog(username string, userID uint, ip string, success bool, message string) {
	s.db.Create(&model.LoginLog{Username: username, UserID: userID, IP: ip, Success: success, Message: message})
}

func toUserResponse(user model.User) dto.UserResponse {
	roles := make([]dto.RoleResponse, 0, len(user.Roles))
	for _, role := range user.Roles {
		roles = append(roles, dto.RoleResponse{ID: role.ID, Code: role.Code, Name: role.Name})
	}
	return dto.UserResponse{
		ID:           user.ID,
		Username:     user.Username,
		Name:         user.Name,
		Email:        user.Email,
		Phone:        user.Phone,
		DepartmentID: user.DepartmentID,
		Status:       user.Status,
		Roles:        roles,
	}
}

func buildMenuTree(menus []model.Menu, parentID uint) []dto.MenuResponse {
	items := make([]dto.MenuResponse, 0)
	for _, menu := range menus {
		if menu.ParentID != parentID {
			continue
		}
		items = append(items, dto.MenuResponse{
			ID:             menu.ID,
			ParentID:       menu.ParentID,
			Name:           menu.Name,
			Path:           menu.Path,
			Icon:           menu.Icon,
			Sort:           menu.Sort,
			PermissionCode: menu.PermissionCode,
			Children:       buildMenuTree(menus, menu.ID),
		})
	}
	return items
}

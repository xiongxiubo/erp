package controller

import (
	"erp/internal/dto"
	"erp/internal/middleware"
	"erp/internal/response"
	"erp/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	svc *service.AuthService
}

func NewAuthController(svc *service.AuthService) *AuthController {
	return &AuthController{svc: svc}
}

func (ctl *AuthController) Login(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, response.BadRequest("请输入用户名和密码"))
		return
	}
	data, err := ctl.svc.Login(req, c.ClientIP())
	if err != nil {
		response.Fail(c, err)
		return
	}
	response.OK(c, data)
}

func (ctl *AuthController) Profile(c *gin.Context) {
	data, err := ctl.svc.Profile(middleware.UserID(c))
	if err != nil {
		response.Fail(c, err)
		return
	}
	response.OK(c, data)
}

func (ctl *AuthController) Permissions(c *gin.Context) {
	data, err := ctl.svc.Permissions(middleware.UserID(c))
	if err != nil {
		response.Fail(c, err)
		return
	}
	response.OK(c, data)
}

func (ctl *AuthController) Menus(c *gin.Context) {
	data, err := ctl.svc.Menus(middleware.UserID(c))
	if err != nil {
		response.Fail(c, err)
		return
	}
	response.OK(c, data)
}

func ParseID(c *gin.Context) (uint, bool) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil || id == 0 {
		response.Fail(c, response.BadRequest("无效的 ID"))
		return 0, false
	}
	return uint(id), true
}

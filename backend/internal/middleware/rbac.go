package middleware

import (
	"erp/internal/response"
	"erp/internal/service"

	"github.com/gin-gonic/gin"
)

func RequirePermission(auth *service.AuthService, code string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if !auth.HasPermission(UserID(c), code) {
			response.Fail(c, response.Forbidden("无访问权限"))
			c.Abort()
			return
		}
		c.Next()
	}
}

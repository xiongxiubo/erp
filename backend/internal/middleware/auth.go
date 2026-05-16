package middleware

import (
	"erp/internal/response"
	"erp/internal/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

func Auth(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		if !strings.HasPrefix(header, "Bearer ") {
			response.Fail(c, response.Unauthorized("未登录"))
			c.Abort()
			return
		}
		claims, err := utils.ParseJWT(strings.TrimPrefix(header, "Bearer "), secret)
		if err != nil {
			response.Fail(c, response.Unauthorized("登录已过期"))
			c.Abort()
			return
		}
		c.Set("userID", claims.UserID)
		c.Set("username", claims.Username)
		c.Next()
	}
}

func UserID(c *gin.Context) uint {
	value, ok := c.Get("userID")
	if !ok {
		return 0
	}
	id, ok := value.(uint)
	if !ok {
		return 0
	}
	return id
}

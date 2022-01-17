package jwt

import (
	enum "crm/app/enums/jwt"
	"crm/app/response"
	"crm/app/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// Jwt jwt中间件
func Jwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code enum.Type = enum.SUCCESS

		token := c.GetHeader("Authorization")

		if token == "" {
			code = enum.InvalidParams
		} else {
			service := &services.Jwt{}

			claims, err := service.ParseToken(token)

			if err != nil {
				code = enum.TokenFail
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = enum.TokenTimeout
			}
		}

		if  code != enum.SUCCESS {
			var response response.Response

			response.SetMessage(string(code)).SetCode(http.StatusUnauthorized).ResponseError(c)

			c.Abort()
		}

		c.Next()
	}
}
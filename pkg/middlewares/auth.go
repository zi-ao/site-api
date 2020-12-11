package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/zi-ao/site-api/app/models/users"
	"github.com/zi-ao/site-api/pkg/response"
	"net/http"
	"strings"
)

const bearerLength = len("Bearer ")

// Auth 用户认证中间件
func Auth() gin.HandlerFunc {
	return func(context *gin.Context) {
		token := context.GetHeader("Authorization")
		if len(token) < bearerLength {
			response.FAIL(context, http.StatusUnauthorized, nil)
			return
		}
		token = strings.TrimSpace(token[bearerLength:])
		user, err := users.ParseToken(token)
		if err != nil {
			response.FAIL(context, http.StatusPreconditionFailed, err.Error())
			return
		}

		context.Set("user", user)
		context.Next()
	}
}

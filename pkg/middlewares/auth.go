package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/zi-ao/site-api/pkg/auth"
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
			response.Abort(context, http.StatusUnauthorized, nil)
			return
		}
		token = strings.TrimSpace(token[bearerLength:])
		user, err := auth.ParseToken(token)
		if err != nil {
			response.Abort(context, http.StatusPreconditionFailed, err.Error())
			return
		}
		if user.ID == 0 {
			response.Abort(context, http.StatusUnauthorized, nil)
			return
		}

		context.Set("user", user)
		context.Next()
	}
}

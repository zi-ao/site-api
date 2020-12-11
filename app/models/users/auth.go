package users

import (
	"github.com/gin-gonic/gin"
)

// Auth 获取认证用户
func Auth(context *gin.Context) *User {
	value, exist := context.Get("user")

	if exist {
		user, ok := value.(*User)

		if ok {
			return user
		}
	}
	return nil
}

// AuthID 获取认证用户 ID
func AuthID(context *gin.Context) uint {
	user := Auth(context)
	if user == nil {
		return 0
	}
	return user.ID
}

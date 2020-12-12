package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/zi-ao/site-api/app/models"
)

// Auth 获取认证用户
func User(context *gin.Context) *models.User {
	value, exist := context.Get("user")

	if exist {
		user, ok := value.(*models.User)

		if ok {
			return user
		}
	}
	return nil
}

// AuthID 获取认证用户 ID
func ID(context *gin.Context) uint {
	user := User(context)
	if user == nil {
		return 0
	}
	return user.ID
}

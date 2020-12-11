package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/zi-ao/site-api/app/models/users"
	"github.com/zi-ao/site-api/app/validation"
	"github.com/zi-ao/site-api/pkg/response"
	"net/http"
	"time"
)

var tokenActivateDuration = time.Hour * 24 * 365

// LoginEndpoint 用户登录 API
func LoginEndpoint(context *gin.Context) {
	var form validation.Login
	if err := context.ShouldBind(&form); err != nil {
		response.FAIL(context, http.StatusBadRequest, err.Error())
		return
	}

	// 用户登录
	user := users.Login(&form)
	if user == nil {
		response.FAIL(context, http.StatusBadRequest, "用户名或密码错误")
		return
	}

	// jwt
	token, err := users.GenerateToken(user, tokenActivateDuration)
	if err != nil {
		response.FAIL(context, http.StatusInternalServerError, "服务器错误")
		return
	}

	response.SUCCESS(context, &map[string]interface{}{
		"token":       token,
		"activate_at": tokenActivateDuration / time.Second,
	})
}

// RegisterEndpoint 用户注册 API
func RegisterEndpoint(context *gin.Context) {
	var form validation.Register
	if err := context.ShouldBind(&form); err != nil {
		response.FAIL(context, http.StatusBadRequest, err.Error())
		return
	}
	_, err := users.Register(&form)
	if err != nil {
		response.FAIL(context, http.StatusInternalServerError, err.Error())
		return
	}
	response.SUCCESS(context, nil)
}

// UpdatePasswordEndpoint 更新用户密码
func UpdatePasswordEndpoint(context *gin.Context) {
	var form validation.UpdatePassword
	if err := context.ShouldBind(&form); err != nil {
		response.FAIL(context, http.StatusBadRequest, err.Error())
		return
	}
	var user *users.User
	id := users.AuthID(context)
	if id != 0 {
		user = users.Find(id)

		if user != nil && users.CheckPassword(user.Password, form.OldPassword) {
			if users.UpdatePassword(user, form.NewPassword) == nil {
				response.SUCCESS(context, nil)
				return
			}
			response.FAIL(context, http.StatusInternalServerError, nil)
			return
		}
		response.FAIL(context, http.StatusBadRequest, "密码错误")
		return
	}
	response.FAIL(context, http.StatusUnauthorized, nil)
}

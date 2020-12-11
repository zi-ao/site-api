package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Basic 响应 JSON 数据
func Basic(context *gin.Context, code int, message, data interface{}) {
	if data == nil {
		data = make(map[string]string)
	}
	context.JSON(code, gin.H{
		"data":    data,
		"message": message,
	})
}

// SUCCESS 返回成功响应
func SUCCESS(context *gin.Context, data interface{}) {
	Basic(context, http.StatusOK, "success", data)
}

// Abort 中止运行并返回 JSON 数据
func Abort(context *gin.Context, code int, data, message interface{}) {
	if data == nil {
		data = make(map[string]string)
	}
	if message == nil {
		var ok bool
		if message, ok = respMessages[code]; !ok {
			message = "error"
		}
	}
	context.AbortWithStatusJSON(code, gin.H{
		"data":    data,
		"message": message,
	})
}

// FAIL 返回失败响应
func FAIL(context *gin.Context, code int, message interface{}) {
	Abort(context, code, nil, message)
}

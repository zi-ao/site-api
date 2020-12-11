package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zi-ao/site-api/pkg/response"
	"net/http"
)

// NotFoundEndpoint 404 页面
func NotFoundEndpoint(context *gin.Context) {
	response.FAIL(context, http.StatusNotFound, nil)
}

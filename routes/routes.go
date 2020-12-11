package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/zi-ao/site-api/app/api"
	v1 "github.com/zi-ao/site-api/app/api/v1"
	"github.com/zi-ao/site-api/pkg/middlewares"
)

func SetupRoutes(engine *gin.Engine) {
	{
		route := engine.Group("/v1")
		route.POST("/login", v1.LoginEndpoint)
		route.POST("/register", v1.RegisterEndpoint)

		route.Use(middlewares.Auth())
		route.POST("update-password", v1.UpdatePasswordEndpoint)
	}

	// 设置 404 路由
	engine.NoRoute(api.NotFoundEndpoint)
	engine.NoMethod(api.NotFoundEndpoint)
}

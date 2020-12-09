package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zi-ao/site-api/bootstrap"
	"github.com/zi-ao/site-api/pkg/logger"
	"github.com/zi-ao/site-api/pkg/middlewares"
	"net/http"
	"strconv"
)

func main() {
	conf := bootstrap.SetupConfig()
	bootstrap.SetupDatabase()

	// 设置 Debug
	if conf.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	// 启动 Gin
	engine := gin.New()
	engine.Use(middlewares.Logger(), gin.Recovery())
	engine.GET("/", func(context *gin.Context) {
		logger.Info(2222)
		context.JSON(http.StatusOK, gin.H{
			"message": "success",
		})
	})
	logger.Info(1111)
	err := engine.Run(":" + strconv.Itoa(int(conf.Port)))
	if err != nil {
		panic(err)
	}
}

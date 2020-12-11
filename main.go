package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zi-ao/site-api/bootstrap"
	"github.com/zi-ao/site-api/pkg/middlewares"
	"github.com/zi-ao/site-api/routes"
	"strconv"
)

func main() {
	conf := bootstrap.SetupConfig()
	bootstrap.SetupDatabase()
	bootstrap.SetupValidator()

	// 设置 Debug
	if conf.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	// 启动 Gin
	engine := gin.New()
	engine.Use(middlewares.Logger(), gin.Recovery())
	routes.SetupRoutes(engine)
	err := engine.Run(":" + strconv.Itoa(int(conf.Port)))
	if err != nil {
		panic(err)
	}
}

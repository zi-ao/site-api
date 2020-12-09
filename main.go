package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zi-ao/site-api/bootstrap"
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
	engine := gin.Default()
	engine.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "success",
		})
	})
	err := engine.Run(":" + strconv.Itoa(int(conf.Port)))
	if err != nil {
		panic(err)
	}
}

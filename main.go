package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	engine := gin.Default()
	engine.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "success",
		})
	})
	err := engine.Run(":9090")
	if err != nil {
		panic(err)
	}
}

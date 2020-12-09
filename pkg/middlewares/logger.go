package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/zi-ao/site-api/pkg/config"
	"github.com/zi-ao/site-api/pkg/logger"
	"time"
)

// Logger 日志中间件
func Logger() gin.HandlerFunc {
	log := logger.Initialization(config.Global.LogPath)

	return func(context *gin.Context) {
		//开始时间
		startTime := time.Now()
		//处理请求
		context.Next()
		//结束时间
		endTime := time.Now()
		// 执行时间
		latencyTime := endTime.Sub(startTime)
		//请求方式
		reqMethod := context.Request.Method
		//请求路由
		reqUrl := context.Request.RequestURI
		//状态码
		statusCode := context.Writer.Status()
		//请求ip
		clientIP := context.ClientIP()

		var msg string
		if statusCode < 300 {
			msg = "SUCCESS"
		} else {
			msg = "ERROR"
		}
		log.WithFields(logrus.Fields{
			"code":         statusCode,
			"ip":           clientIP,
			"url":          reqUrl,
			"method":       reqMethod,
			"latency_time": latencyTime,
		}).Info(msg)
	}
}

package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(context *gin.Context) {
		startTime := time.Now()
		context.Next()
		endTime := time.Now()
		latencyTime := endTime.Sub(startTime)
		reqMethod := context.Request.Method
		reqUri := context.Request.RequestURI
		statusCode := context.Writer.Status()
		clientIp := context.ClientIP()
		// 格式化日志输出
		log.Printf("[GIN] %v | %3d | %13v | %15s | %-7s %s\n",
			endTime.Format("2006/01/02 - 15:04:05"),
			statusCode,
			latencyTime,
			clientIp,
			reqMethod,
			reqUri,
		)
	}
}

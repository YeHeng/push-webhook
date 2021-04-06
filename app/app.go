package app

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

func CreateApp() *gin.Engine {

	r := gin.New()
	r.Use(func(c *gin.Context) {
		// 开始时间
		startTime := time.Now()

		// 处理请求
		c.Next()

		// 结束时间
		endTime := time.Now()

		// 执行时间
		latencyTime := endTime.Sub(startTime)

		// 请求方式
		reqMethod := c.Request.Method

		// 请求路由
		reqUri := c.Request.RequestURI

		// 状态码
		statusCode := c.Writer.Status()

		// 请求IP
		clientIP := c.ClientIP()

		// 日志格式
		Logger.Infow("",
			zap.String("client_ip", clientIP),
			zap.Int("status_code", statusCode),
			zap.Duration("latency_time", latencyTime),
			zap.String("request_method", reqMethod),
			zap.String("request_uri", reqUri),
			zap.Int("response_size", c.Writer.Size()),
		)

	}, gin.Recovery())

	return r
}

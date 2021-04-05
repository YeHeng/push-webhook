package app

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
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
		Logger.WithFields(logrus.Fields{
			"client_ip":      clientIP,
			"status_code":    statusCode,
			"latency_time":   latencyTime,
			"request_method": reqMethod,
			"request_uri":    reqUri,
			"response_size":  c.Writer.Size(),
		}).Infof("%d %s %s", statusCode, reqMethod, reqUri)

	}, gin.Recovery())

	return r
}

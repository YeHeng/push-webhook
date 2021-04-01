package common

import (
	"github.com/YeHeng/qy-wexin-webhook/common/util"
	"github.com/YeHeng/qy-wexin-webhook/internal/alertmanager"
	"github.com/YeHeng/qy-wexin-webhook/internal/grafana"
	"github.com/gin-gonic/gin"
	"time"
)

type Option func(*gin.Engine)

// 初始化
func config(r *gin.Engine, opts ...Option) {
	for _, opt := range opts {
		opt(r)
	}
}

func Init(r *gin.Engine) *gin.Engine {
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
		util.Logger.Infof("| %3d | %13v | %15s | %s | %s |",
			statusCode,
			latencyTime,
			clientIP,
			reqMethod,
			reqUri,
		)
	}, gin.Recovery())

	config(r, alertmanager.Routers, grafana.Routers)

	config := util.AppConfig

	if err := r.Run(":" + config.Port); err != nil {
		util.Logger.Fatalf("Gin start fail. %v", err)
	}

	return r
}

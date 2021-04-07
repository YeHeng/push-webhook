package middleware

import (
	"bytes"
	"github.com/YeHeng/push-webhook/app"
	"github.com/YeHeng/push-webhook/common/util"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	if n, err := w.body.Write(b); err != nil {
		app.Logger.Errorf("%v", err)
		return n, err
	}
	return w.ResponseWriter.Write(b)
}

func Logger() func(c *gin.Context) {
	return func(c *gin.Context) {
		// 开始时间
		startTime := time.Now()

		buffer := util.Borrow()

		blw := &bodyLogWriter{body: buffer, ResponseWriter: c.Writer}
		c.Writer = blw

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

		switch {
		case statusCode >= 400 && statusCode <= 499:
			app.Logger.Warnw(c.Errors.String(), zap.String("client_ip", clientIP),
				zap.Int("status_code", statusCode),
				zap.Duration("latency_time", latencyTime),
				zap.String("request_method", reqMethod),
				zap.String("request_uri", reqUri),
				zap.Int("response_size", c.Writer.Size()))
		case statusCode >= 500:
			app.Logger.Errorw(c.Errors.String(), zap.String("client_ip", clientIP),
				zap.Int("status_code", statusCode),
				zap.Duration("latency_time", latencyTime),
				zap.String("request_method", reqMethod),
				zap.String("request_uri", reqUri),
				zap.Int("response_size", c.Writer.Size()))
		default:

			if blw.body.Len() < 1024 {
				app.Logger.Infow(string(blw.body.Bytes()), zap.String("client_ip", clientIP),
					zap.Int("status_code", statusCode),
					zap.Duration("latency_time", latencyTime),
					zap.String("request_method", reqMethod),
					zap.String("request_uri", reqUri),
					zap.Int("response_size", c.Writer.Size()))
			} else {
				app.Logger.Infow(string(blw.body.Bytes()), zap.String("client_ip", clientIP),
					zap.Int("status_code", statusCode),
					zap.Duration("latency_time", latencyTime),
					zap.String("request_method", reqMethod),
					zap.String("request_uri", reqUri),
					zap.Int("response_size", c.Writer.Size()))
			}
		}

		util.Return(buffer)

	}
}

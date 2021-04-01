package routers

import "github.com/gin-gonic/gin"

type Option func(*gin.Engine)

// 初始化
func Init(r *gin.Engine, opts ...Option) {
	for _, opt := range opts {
		opt(r)
	}
}

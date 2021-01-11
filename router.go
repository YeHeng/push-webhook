package main

import (
	"github.com/YeHeng/qy-wexin-webhook/handler"
	"github.com/gin-gonic/gin"
)

func configRoute(engine *gin.Engine) {
	engine.POST("/alertmanager", handler.AlertManagerHandler())
	engine.POST("/grafana", handler.GrafanaManagerHandler())
}

package grafana

import (
	"github.com/YeHeng/push-webhook/internal/grafana/service"
	"github.com/gin-gonic/gin"
)

func Routers(e *gin.Engine) {
	e.POST("/grafana", service.GrafanaManagerHandler)
}

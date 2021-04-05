package alertmanager

import (
	"github.com/YeHeng/push-webhook/internal/alertmanager/service"
	"github.com/gin-gonic/gin"
)

func Routers(e *gin.Engine) {
	e.POST("/alertmanager", service.AlertManagerHandler)
}

package grafana

import (
	"github.com/gin-gonic/gin"
)

func Routers(e *gin.Engine) {
	e.POST("/webhook/grafana", routeHandler)
}

package alertmanager

import (
	"fmt"
	"github.com/YeHeng/push-webhook/app"
	"github.com/gin-gonic/gin"
	"net/http"
)

func routerHandler(c *gin.Context) {

	transformer, e := app.GetTransformer(AlertManager)
	if e != nil {
		app.Logger.Errorf("%v", e)
		c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("%v", e), "Code": -1})
		return
	}

	msg, e := transformer.Transform(c)
	if e != nil {
		app.Logger.Errorf("%v", e)
		c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("%v", e), "Code": -1})
		return
	}

	pushService, e := app.GetPushStrategy(msg.PushChannel)
	if e != nil {
		app.Logger.Errorf("%v", e)
		c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("%v", e), "Code": -1})
		return
	}

	result, e := pushService.Push(msg)
	if e != nil {
		app.Logger.Errorf("%v", e)
		c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("推送失败，原因:%v", e), "Code": -1})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": result.Message, "Code": result.Code})
}
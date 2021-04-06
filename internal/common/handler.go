package common

import (
	"fmt"
	"net/http"

	api "github.com/YeHeng/push-webhook/api"
	"github.com/YeHeng/push-webhook/app"
	"github.com/gin-gonic/gin"
)

func routeHandler(c *gin.Context) {

	transformer, e := api.GetTransformer(Common)
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

	pushService, e := api.GetPushStrategy(msg.PushChannel)
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

package handler

import (
	"github.com/YeHeng/qy-wexin-webhook/model"
	"github.com/YeHeng/qy-wexin-webhook/notifier"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

func AlertManagerHandler(c *gin.Context, sugar *zap.SugaredLogger) {

	var notification model.Notification

	err := c.BindJSON(&notification)
	key := c.Params.ByName("key")

	sugar.Debug("received alertmanager json: %s, robot key: %s", notification, key)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = notifier.Send(notification, "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key="+key)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	}

	c.JSON(http.StatusOK, gin.H{"message": " successful receive alert notification message!"})

}

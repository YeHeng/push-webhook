package handler

import (
	"encoding/json"
	"github.com/YeHeng/qy-wexin-webhook/model"
	"github.com/YeHeng/qy-wexin-webhook/notifier"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

func AlertManagerHandler(logger *logrus.Logger) gin.HandlerFunc {

	return func(c *gin.Context) {
		var notification model.Notification

		key := c.Query("key")
		err := c.BindJSON(&notification)

		bolB, _ := json.Marshal(notification)

		logger.Debugf("received alertmanager json: %s, robot key: %s", string(bolB), key)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		result, e := notifier.Send(notification, "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key="+key, logger)
		if e != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": e.Error()})
			return

		}

		c.JSON(http.StatusOK, gin.H{"message": result.Message, "Code": result.Code})
	}

}

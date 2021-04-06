package alertmanager

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/YeHeng/push-webhook/app"
	common "github.com/YeHeng/push-webhook/common/model"
	"github.com/YeHeng/push-webhook/internal/push/qywx"
	"github.com/gin-gonic/gin"
	"net/http"
)

const AlertManager string = "ALERT_MANAGER"

type alertManagerTransform struct {
}

func init() {
	app.RegisterTransformer(AlertManager, &alertManagerTransform{})
}

func (s *alertManagerTransform) Transform(c *gin.Context) (*common.PushMessage, error) {
	var notification Notification
	var buffer bytes.Buffer
	err := c.BindJSON(&notification)

	key := c.Query("key")
	bolB, _ := json.Marshal(notification)

	app.Logger.Infof("received alertmanager json: %s, robot key: %s", string(bolB), key)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		app.Logger.Errorf("序列化json异常，原因：%v", err)
		return nil, fmt.Errorf("序列化json异常，原因：%v", err)
	}

	status := notification.Status
	commonLabels := notification.CommonAnnotations

	annotations := notification.CommonAnnotations
	key = annotations["key"]

	buffer.WriteString("## 告警项:\n")

	for _, alert := range notification.Alerts {
		annotations := alert.Annotations
		labels := alert.Labels
		buffer.WriteString(fmt.Sprintf("### 【%s】%s\n", annotations["summary"], annotations["description"]))
		if len(annotations["quantile"]) > 0 && len(annotations["metrics"]) > 0 {
			buffer.WriteString(fmt.Sprintf("### %s超过阀值<font color=\\\"warning\\\">【%s】</font>\n", annotations["metrics"], annotations["quantile"]))
		}
		buffer.WriteString(fmt.Sprintf("\n> NAMESPACES: %s, POD:%s, CONTAINER:%s, IP:%s\n", labels["namespace"], labels["pod"], labels["container"], labels["ip"]))
		buffer.WriteString(fmt.Sprintf("\n> 开始时间：%s\n", alert.StartsAt.Format("2006-01-02 15:04:05-0700")))
		buffer.WriteString(fmt.Sprintf("\n> 结束时间：%s\n", alert.EndsAt.Format("2006-01-02 15:04:05-0700")))
	}

	markdown := &qywx.MarkdownMessage{
		MsgType: "markdown",
		Markdown: &qywx.Markdown{
			Content: fmt.Sprintf("# 【%s】告警(当前状态:%s)\n%s", commonLabels["alertname"], status, buffer.String()),
		},
	}

	content, err := json.Marshal(markdown)
	if err != nil {
		app.Logger.Errorf("序列化json异常，原因：%v", err)
		return nil, err
	}
	return &common.PushMessage{
		Content:     content,
		Key:         key,
		PushChannel: qywx.EnterpriseWechat,
		Params:      nil,
	}, nil
}

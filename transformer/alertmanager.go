package transformer

import (
	"bytes"
	"fmt"
	"github.com/YeHeng/qy-wexin-webhook/model"
)

func AlertManagerToMarkdown(notification model.AlertManagerNotification) (markdown *model.MarkdownMessage, robotURL string, err error) {
	status := notification.Status
	commonLabels := notification.CommonAnnotations

	annotations := notification.CommonAnnotations
	robotURL = annotations["robotUrl"]

	var buffer bytes.Buffer

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

	markdown = &model.MarkdownMessage{
		MsgType: "markdown",
		Markdown: &model.Markdown{
			Content: fmt.Sprintf("# 【%s】告警(当前状态:%s)\n%s", commonLabels["alertname"], status, buffer.String()),
		},
	}
	return
}

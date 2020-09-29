package transformer

import (
	"bytes"
	"fmt"
	"github.com/YeHeng/qy-wexin-webhook/model"
)

func TransformToMarkdown(notification model.AlertManagerNotification) (markdown *model.MarkdownMessage, robotURL string, err error) {
	groupKey := notification.GroupKey
	status := notification.Status

	annotations := notification.CommonAnnotations
	robotURL = annotations["robotUrl"]

	var buffer bytes.Buffer

	buffer.WriteString(fmt.Sprintf("### 通知组%s(当前状态:%s) \n", groupKey, status))

	buffer.WriteString("#### 告警项:\n")

	for _, alert := range notification.Alerts {
		annotations := alert.Annotations
		buffer.WriteString(fmt.Sprintf("##### %s\n > %s\n", annotations["summary"], annotations["description"]))
		buffer.WriteString(fmt.Sprintf("\n> 开始时间：%s\n", alert.StartsAt.Format("15:04:05")))
	}

	markdown = &model.MarkdownMessage{
		MsgType: "markdown",
		Markdown: &model.Markdown{
			Content: fmt.Sprintf("通知组：%s(当前状态:%s)\n%s", groupKey, status, buffer.String()),
		},
		MentionedList: []string{"yeheng@zuzuche.com"},
	}
	return
}

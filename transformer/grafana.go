package transformer

import (
	"bytes"
	"fmt"
	"github.com/YeHeng/qy-wexin-webhook/model"
)

func GrafanaToMarkdown(notification model.GrafanaAlert) (news *model.News, err error) {
	status := notification.State

	var buffer bytes.Buffer

	buffer.WriteString(fmt.Sprintf("## 【%s】告警项:【%s】\n", status, notification.RuleName))

	for _, alert := range notification.EvalMatches {
		buffer.WriteString(fmt.Sprintf("### 指标：【%s】当前值为：%s\n", alert.Metric, alert.Value))
	}

	article := &model.Article{
		Title:       notification.Title,
		Description: buffer.String(),
		URL:         notification.RuleUrl,
		PicURL:      notification.ImageUrl,
	}

	news = &model.News{
		Articles: []model.Article{*article},
	}

	return
}

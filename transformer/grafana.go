package transformer

import (
	"bytes"
	"fmt"
	"github.com/YeHeng/qy-wexin-webhook/model"
)

func GrafanaToMarkdown(notification model.GrafanaAlert) (newsMessage *model.NewMessage, err error) {

	var buffer bytes.Buffer

	buffer.WriteString(fmt.Sprintf("告警项:【%s】\n", notification.RuleName))

	for _, alert := range notification.EvalMatches {
		buffer.WriteString(fmt.Sprintf("指标：【%s】当前值为：%d\n", alert.Metric, alert.Value))
	}

	article := &model.Article{
		Title:       notification.Title,
		Description: buffer.String(),
		URL:         notification.RuleUrl,
		PicURL:      notification.ImageUrl,
	}

	news := &model.News{
		Articles: []model.Article{*article},
	}

	newsMessage = &model.NewMessage{
		News:    news,
		MsgType: "news",
	}

	return
}

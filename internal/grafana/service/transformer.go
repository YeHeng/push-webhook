package service

import (
	"bytes"
	"fmt"
	"github.com/YeHeng/push-webhook/common/model"
	model2 "github.com/YeHeng/push-webhook/internal/grafana/model"
	"strings"
)

func grafanaToMarkdown(notification model2.Alert) (newsMessage *model.NewMessage, qyWxUrl string, err error) {

	var buffer bytes.Buffer
	qyWxUrl = ""
	ruleUrl := notification.RuleUrl

	for _, alert := range notification.EvalMatches {
		buffer.WriteString(fmt.Sprintf("实例：【%s】当前值为：%f\n", alert.Metric, alert.Value))
	}

	if len(notification.Tags) > 0 {
		tags := notification.Tags
		if len(tags["qyweixin_key"]) > 0 {
			qyWxUrl = "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=" + tags["qyweixin_key"]
		}
		if len(tags["base_url"]) > 0 {
			ruleUrl = strings.ReplaceAll(ruleUrl, "http://localhost:3000/", tags["base_url"])
		}
	}

	article := &model.Article{
		Title:       notification.Title,
		Description: buffer.String(),
		URL:         ruleUrl,
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

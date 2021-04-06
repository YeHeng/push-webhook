package grafana

import (
	"bytes"
	"encoding/json"
	"fmt"

	api "github.com/YeHeng/push-webhook/api"
	"github.com/YeHeng/push-webhook/app"
	"github.com/YeHeng/push-webhook/common/model"
	"github.com/YeHeng/push-webhook/internal/push/qywx"
	"github.com/gin-gonic/gin"
)

const Grafana string = "GRAFANA"

type grafanaTransform struct {
}

func init() {
	api.RegisterTransformer(Grafana, &grafanaTransform{})
}

func (s *grafanaTransform) Transform(c *gin.Context) (*model.PushMessage, error) {
	var alert Alert

	var buffer bytes.Buffer
	var param map[string]string
	err := c.BindJSON(&alert)
	param = make(map[string]string)

	key := c.Query("key")
	if len(key) < 0 {
		key = app.Config.Key
	}
	param["key"] = key
	bolB, _ := json.Marshal(alert)

	app.Logger.Infof("received alertmanager json: %s, robot key: %s", string(bolB), key)

	if err != nil {
		app.Logger.Errorf("序列化json异常，原因：%v", err)
		return nil, fmt.Errorf("序列化json异常，原因：%v", err)
	}

	ruleUrl := alert.RuleUrl

	for _, alert := range alert.EvalMatches {
		buffer.WriteString(fmt.Sprintf("实例：【%s】当前值为：%f\n", alert.Metric, alert.Value))
	}

	if len(alert.Tags) > 0 {
		tags := alert.Tags
		if len(tags["qyweixin_key"]) > 0 {
			param["key"] = tags["qyweixin_key"]
		}
	}

	article := &qywx.Article{
		Title:       alert.Title,
		Description: buffer.String(),
		URL:         ruleUrl,
		PicURL:      alert.ImageUrl,
	}

	news := &qywx.News{
		Articles: []qywx.Article{*article},
	}

	newsMessage := &qywx.NewMessage{
		News:    news,
		MsgType: "news",
	}

	content, err := json.Marshal(newsMessage)
	if err != nil {
		app.Logger.Errorf("序列化json异常，原因：%v", err)
		return nil, err
	}
	return &model.PushMessage{
		Content:     string(content),
		Key:         key,
		PushChannel: app.Config.Channel,
		Params:      nil,
	}, nil
}

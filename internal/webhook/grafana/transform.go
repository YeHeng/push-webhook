package grafana

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

const Grafana string = "GRAFANA"

type grafanaTransform struct {
}

func init() {
	app.RegisterTransformer(Grafana, &grafanaTransform{})
}

func (s *grafanaTransform) Transform(c *gin.Context) (*common.PushMessage, error) {
	var alert Alert

	var buffer bytes.Buffer
	var param map[string]string
	err := c.BindJSON(&alert)
	param = make(map[string]string)

	key := c.Query("key")
	param["key"] = key
	bolB, _ := json.Marshal(alert)

	app.Logger.Infof("received alertmanager json: %s, robot key: %s", string(bolB), key)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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
	return &common.PushMessage{
		Content:     content,
		Key:         key,
		PushChannel: qywx.EnterpriseWechat,
		Params:      nil,
	}, nil
}

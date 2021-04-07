package qywx

import (
	"bytes"
	"io/ioutil"
	"net/http"

	api "github.com/YeHeng/push-webhook/api"
	"github.com/YeHeng/push-webhook/app"
	"github.com/YeHeng/push-webhook/common/model"
)

const (
	EnterpriseWechat string = "ENTERPRISE_WECHAT"
)

type EnterpriseWechatPushService struct{}

func init() {
	api.RegisterPushStrategy(EnterpriseWechat, &EnterpriseWechatPushService{})
}

// Save Save
func (s *EnterpriseWechatPushService) Push(msg *model.PushMessage) (model.CommonResult, error) {
	var key string
	key = msg.Key

	if len(key) == 0 {
		return model.CommonResult{
				Code:    404,
				Message: "robot url is nil",
			},
			nil
	}

	req, err := http.NewRequest(
		"POST",
		"https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key="+key,
		bytes.NewBufferString(msg.Content))

	if err != nil {
		return model.CommonResult{
				Code:    400,
				Message: "request robot url fail " + err.Error(),
			},
			nil
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return model.CommonResult{
				Code:    404,
				Message: "request wx api url fail " + err.Error(),
			},
			nil
	}

	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		app.Logger.Fatal(err)
	}
	bodyString := string(bodyBytes)
	app.Logger.Debugf("response: %s, header: %s", bodyString, resp.Header)

	return model.CommonResult{
		Code:    resp.StatusCode,
		Message: bodyString,
	}, nil

}

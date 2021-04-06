package common

import (
	"encoding/json"
	"fmt"

	api "github.com/YeHeng/push-webhook/api"
	"github.com/YeHeng/push-webhook/app"
	"github.com/YeHeng/push-webhook/common/model"
	"github.com/gin-gonic/gin"
	"github.com/hoisie/mustache"
)

const Common string = "COMMON"

type commonTransform struct {
}

func init() {
	api.RegisterTransformer(Common, &commonTransform{})
}

func (s *commonTransform) Transform(c *gin.Context) (*model.PushMessage, error) {

	var pushRequest model.PushRequest
	err := c.BindJSON(&pushRequest)
	if err != nil {
		app.Logger.Errorf("序列化json异常，原因：%v", err)
		return nil, fmt.Errorf("序列化json异常，原因：%v", err)
	}

	if len(pushRequest.Key) < 0 {
		app.Logger.Errorf("序列化json异常，原因：%v", err)
		return nil, fmt.Errorf("序列化json异常，原因：%v", err)
	}

	bolB, _ := json.Marshal(pushRequest)
	app.Logger.Infof("received json: %s, robot key: %s", string(bolB), pushRequest.Key)

	content := mustache.Render(pushRequest.Content, pushRequest.Params)

	return &model.PushMessage{
		Content:     content,
		Key:         pushRequest.Key,
		PushChannel: pushRequest.PushChannel,
		Params:      pushRequest.Params,
	}, nil

}

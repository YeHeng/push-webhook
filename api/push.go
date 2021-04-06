package app

import (
	"fmt"
	"strings"

	"github.com/YeHeng/push-webhook/common/model"
)

type PushStrategy interface {
	Push(content *model.PushMessage) (model.CommonResult, error)
}

var pushServices = map[string]PushStrategy{}

func GetPushStrategy(channel string) (PushStrategy, error) {
	s, ok := pushServices[strings.ToUpper(channel)]
	if !ok {
		return nil, fmt.Errorf("找不到推送渠道: %s", channel)
	}

	return s, nil
}

func RegisterPushStrategy(channel string, service PushStrategy) {
	pushServices[strings.ToUpper(channel)] = service
}

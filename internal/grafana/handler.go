package grafana

import (
	"bytes"
	"encoding/json"
	"github.com/YeHeng/push-webhook/common/app"
	common "github.com/YeHeng/push-webhook/common/model"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func routeHandler(c *gin.Context) {
	var alert Alert

	key := c.Query("key")
	err := c.BindJSON(&alert)

	bolB, _ := json.Marshal(alert)

	app.Logger.Infof("received alertmanager json: %s, robot key: %s", string(bolB), key)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		app.Logger.Errorf("序列化json异常，原因：%v", err)
		return
	}
	result, e := grafanaSend2Wx(alert, "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key="+key)
	if e != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": e.Error()})
		app.Logger.Errorf("推送企业微信异常，原因：%v", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": result.Message, "Code": result.Code})
}

func grafanaSend2Wx(notification Alert, defaultRobot string) (common.CommonResult, error) {

	markdown, qyWxUrl, err := grafanaToMarkdown(notification)

	if err != nil {
		return common.CommonResult{
				Code:    400,
				Message: "marshal json fail " + err.Error(),
			},
			nil
	}

	data, err := json.Marshal(markdown)
	if err != nil {
		return common.CommonResult{
				Code:    400,
				Message: "marshal json fail " + err.Error(),
			},
			nil
	}

	url := qyWxUrl
	if len(url) == 0 {
		url = defaultRobot
	}

	if len(url) == 0 {
		return common.CommonResult{
				Code:    404,
				Message: "robot url is nil",
			},
			nil
	}

	req, err := http.NewRequest(
		"POST",
		url,
		bytes.NewBuffer(data))

	if err != nil {
		return common.CommonResult{
				Code:    400,
				Message: "request robot url fail " + err.Error(),
			},
			nil
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return common.CommonResult{
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

	return common.CommonResult{
		Code:    resp.StatusCode,
		Message: bodyString,
	}, nil
}

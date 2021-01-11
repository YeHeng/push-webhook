package handler

import (
	"bytes"
	"encoding/json"
	. "github.com/YeHeng/qy-wexin-webhook/model"
	. "github.com/YeHeng/qy-wexin-webhook/transformer"
	. "github.com/YeHeng/qy-wexin-webhook/util"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func GrafanaManagerHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var alert GrafanaAlert

		key := c.Query("key")
		err := c.BindJSON(&alert)

		bolB, _ := json.Marshal(alert)

		Logger.Infof("received alertmanager json: %s, robot key: %s", string(bolB), key)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		result, e := grafanaSend2Wx(alert, "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key="+key)
		if e != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": e.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": result.Message, "Code": result.Code})
	}
}

func grafanaSend2Wx(notification GrafanaAlert, defaultRobot string) (ResultVo, error) {

	markdown, err := GrafanaToMarkdown(notification)

	if err != nil {
		return ResultVo{
				Code:    400,
				Message: "marshal json fail " + err.Error(),
			},
			nil
	}

	data, err := json.Marshal(markdown)
	if err != nil {
		return ResultVo{
				Code:    400,
				Message: "marshal json fail " + err.Error(),
			},
			nil
	}

	if len(defaultRobot) == 0 {
		return ResultVo{
				Code:    404,
				Message: "robot url is nil",
			},
			nil
	}

	req, err := http.NewRequest(
		"POST",
		defaultRobot,
		bytes.NewBuffer(data))

	if err != nil {
		return ResultVo{
				Code:    400,
				Message: "request robot url fail " + err.Error(),
			},
			nil
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return ResultVo{
				Code:    404,
				Message: "request wx api url fail " + err.Error(),
			},
			nil
	}

	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		Logger.Fatal(err)
	}
	bodyString := string(bodyBytes)
	Logger.Debugf("response: %s, header: %s", bodyString, resp.Header)

	return ResultVo{
		Code:    resp.StatusCode,
		Message: bodyString,
	}, nil
}

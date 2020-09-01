package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/yunlzheng/alertmanaer-dingtalk-webhook/model"
	"io/ioutil"
	"net/http"
	"strings"
	"webhook/json"
)

func httpCall(notification model.Notification, key string) {
	client := &http.Client{}

	req, err := http.NewRequest("POST", "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key="+key, strings.NewReader("name=cjb"))
	if err != nil {
		// handle error
	}

	var newMessage json.NewsMessage
	newMessage.MsgType = "news"

	var article json.Article
	article.Title = "线上Alert Manager告警"
	article.Url = notification.ExternalURL

	resp, err := client.Do(req)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))
}

func main() {
	router := gin.Default()
	router.POST("/alertmanaer", func(c *gin.Context) {
		var notification model.Notification

		err := c.BindJSON(&notification)
		key := c.Params.ByName("key")

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		go httpCall(notification, key)

		c.JSON(http.StatusOK, gin.H{"message": " successful receive alert notification message!"})

	})

	err := router.Run()
	if err != nil {
		panic(err)
	}

}

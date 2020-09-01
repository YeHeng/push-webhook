package notifier

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/YeHeng/qy-wexin-webhook/model"
	"github.com/YeHeng/qy-wexin-webhook/transformer"
	"log"
	"net/http"
)

// Send send markdown message to dingtalk
func Send(notification model.Notification, defaultRobot string) (err error) {

	markdown, robotURL, err := transformer.TransformToMarkdown(notification)

	if err != nil {
		return
	}

	data, err := json.Marshal(markdown)
	if err != nil {
		return
	}

	var qywxRobotURL string

	if robotURL != "" {
		qywxRobotURL = robotURL
	} else {
		qywxRobotURL = defaultRobot
	}

	if len(qywxRobotURL) == 0 {
		return nil
	}

	req, err := http.NewRequest(
		"POST",
		qywxRobotURL,
		bytes.NewBuffer(data))

	if err != nil {
		log.Fatal("qywx robot url not found ignore:")
		return
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return
	}

	defer resp.Body.Close()
	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)

	return
}

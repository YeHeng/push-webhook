package model

type CommonResult struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type PushMessage struct {
	Key         string
	PushChannel string
	Params      map[string]string
	Content     string
}

type PushRequest struct {
	Key          string            `json:"key"`
	PushChannel  string            `json:"pushChannel"`
	Params       map[string]string `json:"params"`
	TemplateCode string            `json:"templateCode"`
	TemplateType string            `json:"templateType"`
	Mentions     []string          `json:"mentions"`
	Content      string            `json:"content"`
}

package model

type CommonResult struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type PushMessage struct {
	Key         string
	PushChannel string
	Params      map[string]string
	Content     []byte
}

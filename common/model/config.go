package model

type Configuration struct {
	LogConfig LogConfig
	Port      string `default:":9092"`
}

type LogConfig struct {
	Folder   string `default:"./logs/"`
	Filename string `default:"webhook.log"`
}

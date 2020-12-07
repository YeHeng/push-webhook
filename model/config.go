package model

type Configuration struct {
	Log LogConfig
}

type LogConfig struct {
	Folder   string `default:"./logs/"`
	Filename string `default:"webhook.log"`
}

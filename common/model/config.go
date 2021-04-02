package model

import "time"

type Configuration struct {
	LogConfig LogConfig
	Port      string `default:"9092"`
}

type LogConfig struct {
	Folder      string        `default:"./logs/"`
	Filename    string        `default:"webhook.log"`
	Level       string        `default:"info"`
	RollingTime time.Duration `default:"24h"`
	MaxAge      time.Duration `default:"168h"`
}

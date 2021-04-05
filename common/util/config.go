package util

import (
	common "github.com/YeHeng/push-webhook/common/model"
	"github.com/spf13/viper"
	"log"
)

var AppConfig common.Configuration

func init() {
	viper.SetConfigName("config")
	viper.AddConfigPath("./etc/")
	viper.AddConfigPath("/etc/webhook")
	viper.AddConfigPath("$HOME/.webhook")
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	err := viper.Unmarshal(&AppConfig)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}
}

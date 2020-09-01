package main

import (
	"github.com/YeHeng/qy-wexin-webhook/handler"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	router := gin.Default()

	sugar := zap.NewExample().Sugar()
	defer sugar.Sync()

	/*vp := viper.New()
	vp.SetConfigName("wx-robot")
	vp.AddConfigPath(".")
	vp.AddConfigPath("/etc")
	vp.AddConfigPath("/root")
	vp.SetConfigType("yaml")
	err := vp.ReadInConfig()
	if err != nil {
		sugar.Errorf("Failed to read config file. %v", err)
	}
	vp.WatchConfig()
	vp.OnConfigChange(func(e fsnotify.Event) {
		sugar.Infof("log file changed: %s.", e.Name)
	})*/

	router.POST("/alertmanager",
		func(c *gin.Context) {
			handler.AlertManagerHandler(c, sugar)
		})

	err := router.Run()
	if err != nil {
		sugar.Errorf("Gin start fail. %s", err)
	}

}

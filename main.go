package main

import (
	"github.com/YeHeng/push-webhook/app"
	"github.com/YeHeng/push-webhook/internal/alertmanager"
	"github.com/YeHeng/push-webhook/internal/grafana"
)

func main() {

	r := app.CreateApp()
	app.InitRouter(r, alertmanager.Routers, grafana.Routers)

	app.Logger.Infof("开始启动APP!")

	config := app.Config
	if err := r.Run(":" + config.Port); err != nil {
		app.Logger.Fatalf("Gin start fail. %v", err)
	}

}

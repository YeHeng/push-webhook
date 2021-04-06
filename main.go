package main

import (
	"github.com/YeHeng/push-webhook/app"
	"github.com/YeHeng/push-webhook/internal/common"
	"github.com/YeHeng/push-webhook/internal/middleware/alertmanager"
	"github.com/YeHeng/push-webhook/internal/middleware/grafana"
)

func main() {

	r := app.CreateApp()
	app.InitRouter(r, common.Routers, alertmanager.Routers, grafana.Routers)

	app.Logger.Infof("开始启动APP!")

	config := app.Config
	if err := r.Run(":" + config.Port); err != nil {
		app.Logger.Fatalf("Gin start fail. %v", err)
	}

}

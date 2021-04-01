package main

import (
	app "github.com/YeHeng/qy-wexin-webhook/common"
	"github.com/YeHeng/qy-wexin-webhook/common/util"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	util.Logger.Infof("开始启动APP!")
	app.Init(gin.New())
}

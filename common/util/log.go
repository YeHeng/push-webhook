package util

import (
	"fmt"
	nested "github.com/antonfisher/nested-logrus-formatter"
	rotateLogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"time"
)

var Logger *logrus.Logger

func init() {

	if err := os.MkdirAll(AppConfig.LogConfig.Folder, 0777); err != nil {
		fmt.Println(err.Error())
	}

	// 实例化
	Logger = logrus.New()

	//设置日志格式
	Logger.SetFormatter(&nested.Formatter{
		HideKeys:        true,
		TimestampFormat: time.RFC3339,
		FieldsOrder:     []string{"component", "category"},
	})

	// 设置输出
	Logger.Out = getWriter(AppConfig.LogConfig.Folder, AppConfig.LogConfig.Filename)

	// 设置日志级别
	Logger.SetLevel(logrus.InfoLevel)
}

func getWriter(outputDir string, filename string) io.Writer {
	// 生成rotatelogs的Logger 实际生成的文件名 demo.log.YYmmddHH
	// demo.log是指向最新日志的链接
	// 保存7天内的日志，每24小时(整点)分割一次日志
	hook, err := rotateLogs.New(
		// 没有使用go风格反人类的format格式
		outputDir+filename+".%Y%m%d",
		rotateLogs.WithLinkName(outputDir+filename),
		rotateLogs.WithMaxAge(time.Hour*24*7),
		rotateLogs.WithRotationTime(time.Hour*24),
	)
	if err != nil {
		panic(err)
	}
	return hook
}

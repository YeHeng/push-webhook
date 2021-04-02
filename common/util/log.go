package util

import (
	"fmt"
	nested "github.com/antonfisher/nested-logrus-formatter"
	rotateLogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

var Logger *logrus.Logger

func init() {

	if err := os.MkdirAll(AppConfig.LogConfig.Folder, 0777); err != nil {
		fmt.Println(err.Error())
	}

	Logger = logrus.New()

	//设置日志格式
	Logger.SetFormatter(&nested.Formatter{
		TimestampFormat: time.RFC3339,
	})

	Logger.AddHook(newLfsHook())

	level, err := logrus.ParseLevel(AppConfig.LogConfig.Level)

	if err != nil {
		Logger.SetLevel(logrus.WarnLevel)
	} else {
		Logger.SetLevel(level)
	}
}

func newLfsHook() logrus.Hook {

	// 生成rotatelogs的Logger 实际生成的文件名 demo.log.YYmmddHH
	// demo.log是指向最新日志的链接
	// 保存7天内的日志，每24小时(整点)分割一次日志
	writer, err := rotateLogs.New(
		// 没有使用go风格反人类的format格式
		AppConfig.LogConfig.Folder+AppConfig.LogConfig.Filename+".%Y%m%d",
		rotateLogs.WithLinkName(AppConfig.LogConfig.Folder+AppConfig.LogConfig.Filename),
		rotateLogs.WithMaxAge(AppConfig.LogConfig.MaxAge),
		rotateLogs.WithRotationTime(AppConfig.LogConfig.RollingTime),
	)
	if err != nil {
		panic(err)
	}

	lfsHook := lfshook.NewHook(lfshook.WriterMap{
		logrus.DebugLevel: writer,
		logrus.InfoLevel:  writer,
		logrus.WarnLevel:  writer,
		logrus.ErrorLevel: writer,
		logrus.FatalLevel: writer,
		logrus.PanicLevel: writer,
	}, &logrus.JSONFormatter{
		TimestampFormat: time.RFC3339,
	})
	return lfsHook
}

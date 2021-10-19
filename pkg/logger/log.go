package logger

import (
	"fmt"
	"strconv"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/medivh-jay/lfshook"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	// terminal 默认终端log输出
	terminal = logrus.New()
	logger   = terminal

	// 每天切割一次
	rotationTime = rotatelogs.WithRotationTime(time.Hour * 24)
	writer       *rotatelogs.RotateLogs

	logLevel = map[string]logrus.Level{
		"panic": logrus.PanicLevel,
		"fatal": logrus.FatalLevel,
		"error": logrus.ErrorLevel,
		"warn":  logrus.WarnLevel,
		"info":  logrus.InfoLevel,
		"debug": logrus.DebugLevel,
		"trace": logrus.TraceLevel,
	}
)

// Initiate 启动日志服务
func Initiate() {
	terminal.SetFormatter(&logrus.TextFormatter{ForceColors: true})
	terminal.SetNoLock()

	terminal.Level = logLevel[viper.GetString("log.level")]
	terminal.AddHook(lfshook.NewHook(
		lfshook.WriterMap{
			logrus.InfoLevel:  newWriter("info"),
			logrus.ErrorLevel: newWriter("error"),
			logrus.DebugLevel: newWriter("debug"),
			logrus.WarnLevel:  newWriter("warn"),
			logrus.TraceLevel: newWriter("trace"),
		}, &JSONFormatter{}))

	terminal.SetOutput(nilWriter{})
	terminal.SetReportCaller(true)
}
func newWriter(level string) *rotatelogs.RotateLogs {
	var (
		err        error
		FileFormat = viper.GetString("log.file_format")
		Path       = viper.GetString("log.path")
		MaxAge     = "7"
	)
	// 保留天数
	if viper.IsSet("log.max_save_days") {
		MaxAge = viper.GetString("log.max_save_days")
	}
	saveDays, _ := strconv.Atoi(MaxAge)

	if viper.IsSet("log.path") {
		Path = viper.GetString("log.path")
	}
	if viper.IsSet("log.file_format") {
		FileFormat = viper.GetString("log.file_format")
	}

	writer, err = rotatelogs.New(
		fmt.Sprintf("%s/%s.%s", Path, level, FileFormat)+".log", rotatelogs.WithMaxAge(time.Hour*24*time.Duration(saveDays)), rotationTime)

	if err != nil {
		panic(err)
	}

	return writer
}

// Logger 对外日志操作
//  filter 过滤 key, 在日志记录中将会存在一个 key 为 filter 值为传入的值的元素, 供筛选日志使用
func Logger(filter string) *logrus.Entry {
	return logger.WithField("filter", filter)
}

// 当添加了文件写入的 hook 之后禁用终端输出
type nilWriter struct {
}

func (nilWriter) Write(_ []byte) (n int, err error) {
	return 0, nil
}

package log

import (
	"github.com/rs/zerolog"
	l "github.com/rs/zerolog/log"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

type Config struct {
	Level      string // 调试模式，默认仅记录错误
	Output     string // 日志打印方式。none不打印日志，console打印到控制台，file输出到文件。默认为console
	File       string // 日志文件路径
	MaxSize    int    // 单个日志文件的大小限制，单位MB
	MaxBackups int    // 日志文件数据的限制
	MaxAge     int    // 日志文件的保存天数
	Compress   bool   // 日志文件压缩开关
}

func InitLogger(config Config) {
	timeFormat := "2006-01-02 15:04:05"
	zerolog.TimeFieldFormat = timeFormat
	// 日志级别
	switch config.Level {
	case "debug":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	case "info":
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	case "warn":
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	case "error":
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	case "fatal":
		zerolog.SetGlobalLevel(zerolog.FatalLevel)
	case "panic":
		zerolog.SetGlobalLevel(zerolog.PanicLevel)
	case "none":
		zerolog.SetGlobalLevel(zerolog.Disabled)
	}

	switch config.Output {
	case "none":
		zerolog.SetGlobalLevel(zerolog.Disabled)
	case "console":
		consoleWriter := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: timeFormat}
		l.Logger = zerolog.New(consoleWriter).With().Timestamp().Logger()
	case "file":
		logRate := &lumberjack.Logger{
			Filename:   config.File,
			MaxSize:    config.MaxSize,
			MaxBackups: config.MaxBackups,
			MaxAge:     config.MaxAge,
			Compress:   config.Compress,
		}
		l.Logger = l.Output(logRate)
	}

}

func SetTraceId(traceId string) {
	l.Logger = l.With().Str("trace_id", traceId).Logger()
}

func Err(err error) *zerolog.Event {
	return l.Logger.Err(err)
}

func Debug() *zerolog.Event {
	return l.Logger.Debug()

}

func Info() *zerolog.Event {
	return l.Logger.Info()
}

func Warn() *zerolog.Event {
	return l.Logger.Warn()
}

func Error() *zerolog.Event {
	return l.Logger.Error()
}

func Fatal() *zerolog.Event {
	return l.Logger.Fatal()
}

func Panic() *zerolog.Event {
	return l.Logger.Panic()
}

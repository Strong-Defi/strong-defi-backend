package config

import (
	"github.com/ethereum/go-ethereum/log"
	"github.com/spf13/viper"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	c "strong-defi-backend/resources"
)

var (
	Log    log.Logger
	Config = &ConfigInfo{}
)

type ConfigInfo struct {
	Logging c.Logging
	App     c.App
	Db      c.Db
}

func InitConfig() {

	logging := c.Logging{}
	app := c.App{}
	db := c.Db{}

	logging.IsStdout = viper.GetBool("log.isStdout")
	logging.Filename = viper.GetString("log.filename")
	logging.Prefix = viper.GetString("log.prefix")
	logging.MaxAge = viper.GetInt("log.maxAge")
	logging.MaxSize = viper.GetInt("log.maxSize")
	logging.MaxBackups = viper.GetInt("log.maxBackups")
	logging.Compress = viper.GetBool("log.compress")
	//启动地址
	app.Port = viper.GetString("app.port")
	//管理员部署地址
	app.AdminAddress = viper.GetString("app.adminAddress")
	//部署地址
	app.DeployAddress = viper.GetString("app.deployAddress")

	//db的配置
	db.Dsn = viper.GetString("mysql.dsn")
	db.Idle = viper.GetInt("mysql.idle")
	db.IdleTimeout = viper.GetInt("mysql.idleTimeout")
	db.Active = viper.GetInt("mysql.active")

	Config.App = app
	Config.Logging = logging
	Config.Db = db
}

/**
 * 初始化日志
 */
func InitLog() {
	/*通过lumberjack，引入日志文件*/
	filename := ""                      // 日志文件名
	maxSize := 1                        // 每个日志文件的最大大小（MB）
	maxBackups := 5                     // 保留旧日志文件的最大数量
	maxAge := 30                        // 保留旧日志文件的最大天数
	compress := Config.Logging.Compress //设置 compress 为 true 以启用压缩

	if Config.Logging.Filename != "" {
		if Config.Logging.Prefix != "" {
			filename = Config.Logging.Prefix + Config.Logging.Filename
		} else {
			filename = Config.Logging.Filename
		}
	}

	if Config.Logging.MaxSize > 0 {
		maxSize = Config.Logging.MaxSize
	}

	if Config.Logging.MaxBackups > 0 {
		maxBackups = Config.Logging.MaxBackups
	}
	if Config.Logging.MaxAge > 0 {
		maxAge = Config.Logging.MaxAge
	}

	//if Config.Logging.Compress > 0 {
	//	maxAge = Config.Logging.MaxAge
	//}
	var logHandle *log.TerminalHandler
	if Config.Logging.IsStdout {
		logHandle = log.NewTerminalHandlerWithLevel(os.Stdout, log.LvlInfo, true)

	} else {
		logFile := &lumberjack.Logger{
			Filename:   filename,
			MaxSize:    maxSize,
			MaxBackups: maxBackups,
			MaxAge:     maxAge,
			Compress:   compress,
		}
		logHandle = log.NewTerminalHandlerWithLevel(logFile, log.LvlInfo, true)
	}

	Log = log.NewLogger(logHandle)
}

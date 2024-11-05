package config

import (
	"github.com/ethereum/go-ethereum/log"
	"github.com/spf13/viper"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

type App struct {
	Port string
	//管理员钱包地址
	AdminAddress string
	//链地址
	DeployAddress string
	//合约地址
	ContractAddress string
}
type Db struct {
	Dsn         string // 是否输出到控制台
	Active      int    //链接池数量
	Idle        int    //存货
	IdleTimeout int    //超时时间
}

type Config struct {
	Logging Logging
	App     App
	Db      Db
}

type Logging struct {
	IsStdout   bool   // 是否输出到控制台
	Filename   string //日志路径
	Prefix     string //日志前缀
	MaxSize    int
	MaxBackups int
	MaxAge     int
	Compress   bool
}

var Conf Config
var Log log.Logger

func initConfig() {
	logging := Logging{}
	app := App{}
	db := Db{}

	logging.IsStdout = viper.GetBool("log.isStdout")
	logging.Filename = viper.GetString("log.filename")
	logging.Prefix = viper.GetString("log.prefix")
	logging.MaxAge = viper.GetInt("log.maxAge")
	logging.MaxSize = viper.GetInt("log.maxSize")
	logging.MaxBackups = viper.GetInt("log.maxBackups")
	logging.Compress = viper.GetBool("log.compress")
	//启动地址
	app.Port = viper.GetString("app.port")
	//管理员钱包地址
	app.AdminAddress = viper.GetString("app.adminAddress")
	//部署地址
	app.DeployAddress = viper.GetString("app.deployAddress")
	//合约地址
	app.ContractAddress = viper.GetString("app.contractAddress")

	//db的配置
	db.Dsn = viper.GetString("mysql.dsn")
	db.Idle = viper.GetInt("mysql.idle")
	db.IdleTimeout = viper.GetInt("mysql.idleTimeout")
	db.Active = viper.GetInt("mysql.active")

	Conf.App = app
	Conf.Logging = logging
	Conf.Db = db
}
func init() {

	initConfig()
	initLog()
}

// TODO::此处待优化
func initLog() {
	/*通过lumberjack，引入日志文件*/
	filename := ""                    // 日志文件名
	maxSize := 1                      // 每个日志文件的最大大小（MB）
	maxBackups := 5                   // 保留旧日志文件的最大数量
	maxAge := 30                      // 保留旧日志文件的最大天数
	compress := Conf.Logging.Compress //设置 compress 为 true 以启用压缩

	if Conf.Logging.Filename != "" {
		if Conf.Logging.Prefix != "" {
			filename = Conf.Logging.Prefix + Conf.Logging.Filename
		} else {
			filename = Conf.Logging.Filename
		}
	}

	if Conf.Logging.MaxSize > 0 {
		maxSize = Conf.Logging.MaxSize
	}

	if Conf.Logging.MaxBackups > 0 {
		maxBackups = Conf.Logging.MaxBackups
	}
	if Conf.Logging.MaxAge > 0 {
		maxAge = Conf.Logging.MaxAge
	}

	//if Conf.Logging.Compress > 0 {
	//	maxAge = Conf.Logging.MaxAge
	//}
	var logHandle *log.TerminalHandler
	if Conf.Logging.IsStdout {
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

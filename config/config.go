package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"strong-defi-backend/pkg/log"
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

type AppConfig struct {
	Log log.Config
	App App
	Db  Db
}

var Config AppConfig

func init() {
	// 读取配置
	workDir, _ := os.Getwd()
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(workDir)
	fmt.Println(workDir)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	// 应用配置
	app := App{
		Port:            viper.GetString("app.port"),
		AdminAddress:    viper.GetString("app.adminAddress"),    //管理员钱包地址
		DeployAddress:   viper.GetString("app.deployAddress"),   //部署地址
		ContractAddress: viper.GetString("app.contractAddress"), //合约地址
	}
	// 数据库配置
	db := Db{
		Dsn:         viper.GetString("mysql.dsn"),
		Active:      viper.GetInt("mysql.active"),
		Idle:        viper.GetInt("mysql.idle"),
		IdleTimeout: viper.GetInt("mysql.idleTimeout"),
	}
	// 日志配置
	logConfig := log.Config{
		Level:      viper.GetString("log.debug"),
		Output:     viper.GetString("log.output"),
		File:       viper.GetString("log.filename"),
		MaxSize:    viper.GetInt("log.maxSize"),
		MaxBackups: viper.GetInt("log.maxBackups"),
		MaxAge:     viper.GetInt("log.maxAge"),
		Compress:   viper.GetBool("log.compress"),
	}

	Config.App = app
	Config.Log = logConfig
	Config.Db = db
}

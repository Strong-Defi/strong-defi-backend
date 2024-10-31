package main

import (
	"github.com/strong-defi/strong-defi-backend/pkg/db"
	"github.com/strong-defi/strong-defi-backend/pkg/ethereum"
	"github.com/strong-defi/strong-defi-backend/pkg/log"
)

// App 应用，目前没用起来，之后想办法扩展
type App struct {
	Logger    log.Logger
	SysConfig *SysConfig
}

type SysConfig struct {
	Server *Server          `yaml:"server"`
	Logger *log.Config      `yaml:"logger"`
	Mysql  *db.Config       `yaml:"mysql"`
	Chains *ethereum.Config `yaml:"chains"`
}

type Server struct {
	Host          string `yaml:"host"`
	Port          string `yaml:"port"`
	AdminAddress  string `yaml:"adminAddress"`
	DeployAddress string `yaml:"deployAddress"`
}

func InitApp(sysConfig *SysConfig) *App {
	app := &App{
		Logger:    InitLog(sysConfig.Logger),
		SysConfig: sysConfig,
	}
	return app
}

// InitLog 初始化日志
func InitLog(config *log.Config) log.Logger {
	logger, err := log.NewLogger(config)
	if err != nil {
		panic(err)
	}
	return logger
}

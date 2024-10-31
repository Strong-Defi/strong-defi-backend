package main

import (
	"github.com/ethereum/go-ethereum/log"
	"github.com/spf13/viper"
	"github.com/strong-defi/strong-defi-backend/internal/server"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	app *App
)

func init() {
	config := initConfig()
	app = InitApp(config)
}

func main() {

	c := make(chan os.Signal, 1)
	server.NewHttpServer(app.SysConfig).Run()
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		log.Info("get a signal %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			//closeFunc()
			log.Info("auth exit")
			time.Sleep(time.Second)
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}

func initConfig() *SysConfig {
	workDir, err := os.Getwd()

	if err != nil {
		panic(err)
	}
	viper.SetConfigName("conf")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(workDir + "/configs")
	err = viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	config := SysConfig{}
	err = viper.Unmarshal(&config)
	if err != nil {
		panic(err)
	}
	return &config
}

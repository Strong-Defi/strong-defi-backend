package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"strong-defi-backend/config"
	"strong-defi-backend/model"
	server "strong-defi-backend/router"
	"strong-defi-backend/service"
	"time"
)

func main() {
	/*启动*/
	gin.SetMode(gin.DebugMode)
	engine := gin.Default()
	/*初始化日志一定放到最上面，下面的init需要用到日志*/
	initLogAndConf()

	db := dbInit()
	/*服务初始化*/
	serverInit(engine, db)
	/*初始化数据库*/

	readPortErr := os.Setenv("PORT", config.Config.App.Port)
	if readPortErr != nil {
		panic(readPortErr)
	}
	err := engine.Run()
	if err != nil {
		return
	}
}

// 初始化日志，并且引入conf.json文件
func initLogAndConf() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("conf")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(workDir + "/resources")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	config.InitConfig()
	config.InitLog()
}

// 数据库初始化
func dbInit() *gorm.DB {
	dbConfig := config.Config.Db
	db, err := gorm.Open(mysql.Open(dbConfig.Dsn), &gorm.Config{})

	s, err := db.DB()
	if err != nil {
		panic(err)
	}
	/*最大空闲连接数*/
	s.SetMaxIdleConns(dbConfig.Idle)
	/*最大活跃数*/
	s.SetMaxOpenConns(dbConfig.Active)
	s.SetConnMaxLifetime(time.Duration(dbConfig.IdleTimeout) * time.Hour)

	if err != nil {
		return nil
	}
	return db
}

// 初始化业务server
func serverInit(engine *gin.Engine, db *gorm.DB) {
	server.AllInterface(engine)
	service.New(&model.Dao{Orm: db})
}

package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"strong-defi-backend/config"
	"strong-defi-backend/model"
	"strong-defi-backend/pkg/log"
	server "strong-defi-backend/router"
	"strong-defi-backend/service"
	"time"
)

func main() {

	// 初始化日志
	log.InitLogger(config.Config.Log)

	/*启动*/
	gin.SetMode(gin.DebugMode)
	engine := gin.Default()

	db := dbInit()
	/*服务初始化*/
	serverInit(engine, db)
	/*初始化数据库*/

	readPortErr := os.Setenv("PORT", config.Config.App.Port)
	if readPortErr != nil {
		panic(readPortErr)
	}

	err := engine.Run(fmt.Sprintf("%s:%s", config.Config.App.Host, config.Config.App.Port))
	if err != nil {
		return
	}
}

// 数据库初始化
func dbInit() *gorm.DB {
	dbConfig := config.Config.Db
	db, err := gorm.Open(mysql.Open(dbConfig.Dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
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
	server.Route(engine)
	dao := &model.Dao{Orm: db}
	service.New(dao)
}

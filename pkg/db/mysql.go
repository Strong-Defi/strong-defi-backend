package db

import (
	"errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

type Config struct {
	DSN         string `yaml:"dsn"`
	Active      int    `yaml:"active"`
	Idle        int    `yaml:"idle"`
	IdleTimeout int    `yaml:"idleTimeout"`
}

// NewMysql 数据库初始化
func NewMysql(config *Config) (db *gorm.DB, err error) {
	if config == nil {
		return nil, errors.New("db config is nil")
	}
	db, err = gorm.Open(mysql.Open(config.DSN), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	s, err := db.DB()
	if err != nil {
		return nil, err
	}
	/*最大空闲连接数*/
	s.SetMaxIdleConns(config.Idle)
	/*最大活跃数*/
	s.SetMaxOpenConns(config.Active)
	s.SetConnMaxLifetime(time.Duration(config.IdleTimeout) * time.Hour)
	return db, nil
}

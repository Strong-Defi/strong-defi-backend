package dao

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	logs zap.Logger
)

type Dao struct {
	Orm *gorm.DB
}

func New(logger zap.Logger, db *gorm.DB) *Dao {
	logs = logger
	return &Dao{Orm: db}
}

/*get grom db*/
func (d *Dao) ORM() *gorm.DB {
	return d.Orm
}

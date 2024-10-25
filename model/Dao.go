package model

import (
	"github.com/ethereum/go-ethereum/log"
	"gorm.io/gorm"
	"strong-defi-backend/config"
)

var (
	logs log.Logger
)

func init() {
	logs = config.Log
}

type Dao struct {
	Orm *gorm.DB
}

/*get grom db*/
func (d *Dao) ORM() *gorm.DB {
	return d.Orm
}

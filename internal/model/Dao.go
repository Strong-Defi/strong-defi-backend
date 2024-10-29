package model

import (
	"github.com/Strong-Defi/strong-defi-backend/cmd"
	"github.com/ethereum/go-ethereum/log"
	"gorm.io/gorm"
)

var (
	logs log.Logger
)

func init() {
	logs = main.Log
}

type Dao struct {
	Orm *gorm.DB
}

/*get grom db*/
func (d *Dao) ORM() *gorm.DB {
	return d.Orm
}

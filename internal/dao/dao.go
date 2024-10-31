package dao

import (
	pdb "github.com/strong-defi/strong-defi-backend/pkg/db"
	plog "github.com/strong-defi/strong-defi-backend/pkg/log"
	"gorm.io/gorm"
)

var (
	log plog.Logger
)

type Dao struct {
	db *gorm.DB
}

// New 实例化Dao
func New(logger plog.Logger, config *pdb.Config) *Dao {
	log = logger
	db, err := pdb.NewMysql(config)
	if err != nil {
		panic(err)
	}
	return &Dao{db: db}
}

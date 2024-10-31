package dao

import (
	pdb "github.com/strong-defi/strong-defi-backend/pkg/db"
	"github.com/strong-defi/strong-defi-backend/pkg/ethereum"
	plog "github.com/strong-defi/strong-defi-backend/pkg/log"
	"gorm.io/gorm"
)

var (
	log plog.Logger
)

type Dao struct {
	db        *gorm.DB
	ethClient *ethereum.Client
}

// New 实例化Dao
func New(logger plog.Logger, config *pdb.Config, chainConfig *ethereum.Config) *Dao {
	log = logger
	db, err := pdb.NewMysql(config)
	if err != nil {
		panic(err)
	}
	sepoliaConfig, err := ethereum.NewClientFromConfig("sepolia", chainConfig)
	if err != nil {
		panic(err)
	}
	return &Dao{db: db, ethClient: sepoliaConfig}
}

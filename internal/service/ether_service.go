package service

import (
	"github.com/strong-defi/strong-defi-backend/internal/dao"
)

type EtherService struct {
	dao *dao.Dao
}

func NewEtherService(dao *dao.Dao) *EtherService {
	log.Info("NewEtherService is called")
	return &EtherService{
		dao: dao,
	}
}

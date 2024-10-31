package service

import (
	"github.com/strong-defi/strong-defi-backend/internal/dao"
)

type EtherService struct {
	dao *dao.Dao
}

func NewEtherService(dao *dao.Dao) *EtherService {
	return &EtherService{
		dao: dao,
	}
}

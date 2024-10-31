//go:build wireinject
// +build wireinject

package service

import (
	"github.com/google/wire"
	"github.com/strong-defi/strong-defi-backend/internal/dao"
	pdb "github.com/strong-defi/strong-defi-backend/pkg/db"
	plog "github.com/strong-defi/strong-defi-backend/pkg/log"
)

// ProviderSet is dao providers.
var ProviderSet = wire.NewSet(New, NewEtherService, NewStakeService, NewUserService)

func InitService(logger plog.Logger, config *pdb.Config) (*Service, func(), error) {
	panic(wire.Build(dao.New, ProviderSet))
	return new(Service), func() {}, nil
}

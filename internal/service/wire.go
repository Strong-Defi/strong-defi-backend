//go:build wireinject
// +build wireinject

package service

import (
	"github.com/strong-defi/strong-defi-backend/internal/dao"
)

// ProviderSet is dao providers.
var ProviderSet = wire.NewSet(New)

//go:generate go generate ./...
func InitService() (*service.Service, func(), error) {
	panic(wire.Build(dao.Provider, ProviderSet))
}

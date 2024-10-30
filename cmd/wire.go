//go:build wireinject
// +build wireinject

package main

import (
	"github.com/strong-defi/strong-defi-backend/internal/server"
)

var ProviderSet = wire.NewSet(NewApp, server.NewHttpServer)

func CreateApp(config *Config, server *server.HttpServer) (*App, func(), error) {
	panic(wire.Build(ProviderSet))
	return nil, nil, nil
}

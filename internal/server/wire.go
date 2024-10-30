package server

var ProviderSet = wire.NewSet(NewApp, server.NewHttpServer)

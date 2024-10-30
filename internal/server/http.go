package server

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var (
	log *zap.Logger
)

type HttpServer struct {
	engine *gin.Engine
}

func NewHttpServer(logger zap.Logger) (server *HttpServer, cancel func(), err error) {
	log = logger.With(zap.String("server", "http-server"))

	engine := gin.Default()

	server = &HttpServer{
		engine: engine,
	}

	return server, func() {
	}, nil
}

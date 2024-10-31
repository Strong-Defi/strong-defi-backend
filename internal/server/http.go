package server

import (
	"github.com/gin-gonic/gin"
	"github.com/strong-defi/strong-defi-backend/internal/router"
	"github.com/strong-defi/strong-defi-backend/internal/service"
	plog "github.com/strong-defi/strong-defi-backend/pkg/log"
)

var (
	log plog.Logger
)

type HttpServer struct {
	engine *gin.Engine
}

func (s *HttpServer) Start() {
	s.engine.Run()
}

// NewHttpServer creates a new http server
func NewHttpServer(log plog.Logger, engine *gin.Engine, service *service.Service) (server *HttpServer, cancel func(), err error) {
	// TODO cors
	server = &HttpServer{
		engine: engine,
	}
	RegisterRouter(engine, service.UserService, service.StakeService, service.EtherService)
	return server, func() {
		log.Info("http server shutdown")
	}, nil
}

func RegisterRouter(r *gin.Engine, userService *service.UserService,
	stakeService *service.StakeService, etherService *service.EtherService) *gin.Engine {
	router.InjectEtherRouter(r, etherService)
	router.InjectStakeRouter(r, stakeService)
	router.InjectUserRouter(r, userService)
	return r
}

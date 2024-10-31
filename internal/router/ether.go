package router

import (
	"github.com/gin-gonic/gin"
	"github.com/strong-defi/strong-defi-backend/internal/service"
)

func InjectEtherRouter(engine *gin.Engine, service *service.EtherService) *gin.Engine {

	return engine

}

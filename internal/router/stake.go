package server

import (
	"github.com/gin-gonic/gin"
	"github.com/strong-defi/strong-defi-backend/internal/service"
)

func InjectStakeRouter(r *gin.Engine, service service.StakeService) {
	e := r.Group("/v1/stake")
	{
		e.POST("/getAccountBalance", service.GetAccountBalance)
		e.POST("/addPool", service.AddPool)
		e.POST("/getPoolBalance", service.GetPoolBalance)
	}
}

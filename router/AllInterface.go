package server

import (
	"github.com/gin-gonic/gin"
	"strong-defi-backend/service"
)

func AllInterface(r *gin.Engine) *gin.Engine {
	e := r.Group("/gin")
	{
		e.GET("/login", service.GetContractDetail)
		e.GET("/register", service.GetContractInfo)

	}
	r.POST("/basicDemo", service.BasicDemo)
	return r
}

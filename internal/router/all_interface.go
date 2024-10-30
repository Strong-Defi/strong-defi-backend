package server

import (
	"github.com/gin-gonic/gin"
	"github.com/strong-defi/strong-defi-backend/internal/middleware"
	service2 "github.com/strong-defi/strong-defi-backend/internal/service"
)

func AllInterface(r *gin.Engine) *gin.Engine {
	e := r.Group("/gin/user", middleware.CheckToken)
	{
		e.GET("/login", service2.UserLogin)
		e.POST("/UserRegister", service2.UserRegister)
		e.GET("/register", service2.GetContractInfo)
		e.GET("/queryIsExist", service2.QueryUserIsExist)
		e.GET("/queryByWalletAddress", service2.QueryByWalletAddress)
		e.GET("/queryByTelephone", service2.QueryByTelephone)
		e.POST("/modifyUser", service2.ModifyUser)
		e.GET("/deleteUser", service2.DeleteUser)
	}

	i := r.Group("/gin/stake", middleware.CheckToken)
	{
		i.POST("/getAccountBalance", service2.GetAccountBalance)
		i.POST("/addPool", service2.AddPool)
		i.POST("/getPoolBalance", service2.GetPoolBalance)
	}

	return r
}

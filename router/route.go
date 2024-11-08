package server

import (
	"github.com/gin-gonic/gin"
	"strong-defi-backend/middleware"
	"strong-defi-backend/service"
)

func Route(r *gin.Engine) *gin.Engine {

	e := r.Group("/gin/user", middleware.Logging(), middleware.Auth())
	{
		e.GET("/login", service.UserLogin)
		e.POST("/UserRegister", service.UserRegister)
		e.GET("/register", service.GetContractInfo)
		e.GET("/queryIsExist", service.QueryUserIsExist)
		e.GET("/queryByWalletAddress", service.QueryByWalletAddress)
		e.GET("/queryByTelephone", service.QueryByTelephone)
		e.POST("/modifyUser", service.ModifyUser)
		e.GET("/deleteUser", service.DeleteUser)
	}

	i := r.Group("/gin/stake", middleware.Logging(), middleware.Auth())
	{
		i.POST("/getAccountBalance", service.GetAccountBalance)
		i.POST("/addPool", service.AddPool)
		i.POST("/userApprove", service.UserApprove)
		i.POST("/userStake", service.UserStake)
		i.POST("/getUserInfo", service.GetUserInfo)
		i.POST("/getAllPool", service.GetAllPool)
		i.GET("/getPoolDetail", service.GetPoolDetail)
		i.POST("/userTransfer", service.UserTransfer)
		i.POST("/tokenBalance", service.TokenBalance)
	}

	return r
}

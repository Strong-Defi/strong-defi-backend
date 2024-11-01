package server

import (
	"github.com/gin-gonic/gin"
	"strong-defi-backend/service"
)

func AllInterface(r *gin.Engine) *gin.Engine {
	e := r.Group("/gin/user", service.CheckToken)
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

	i := r.Group("/gin/stake", service.CheckToken)
	{
		i.POST("/getAccountBalance", service.GetAccountBalance)
		i.POST("/addPool", service.AddPool)
		i.POST("/userStake", service.UserStake)
		i.POST("/getUserInfo", service.GetUserInfo)
		i.POST("/getAllPool", service.HandleRequestPage, service.GetAllPool)
		i.GET("/getPoolDetail", service.GetPoolDetail)
	}

	return r
}

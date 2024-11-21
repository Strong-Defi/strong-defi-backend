package server

import (
	"github.com/gin-gonic/gin"
	"strong-defi-backend/middleware"
	"strong-defi-backend/service"
)

func Route(r *gin.Engine) *gin.Engine {

	lr := r.Group("/gin", middleware.Logging())
	{
		lr.GET("/user/login", service.UserLogin)
		lr.POST("/user/UserRegister", service.UserRegister)

		ar := lr.Group("/user")
		ar.Use(middleware.Auth())
		{
			lr.GET("/register", service.GetContractInfo)
			lr.GET("/queryIsExist", service.QueryUserIsExist)
			lr.GET("/queryByWalletAddress", service.QueryByWalletAddress)
			lr.GET("/queryByTelephone", service.QueryByTelephone)
			lr.POST("/modifyUser", service.ModifyUser)
			lr.GET("/deleteUser", service.DeleteUser)
		}

		//质押相关接口
		sr := lr.Group("/stake")
		sr.Use(middleware.Auth())
		{
			sr.POST("/getAccountBalance", service.GetAccountBalance)
			sr.POST("/addPool", service.AddPool)
			sr.POST("/userApprove", service.UserApprove)
			sr.POST("/userStake", service.UserStake)
			sr.POST("/getUserInfo", service.GetUserInfo)
			sr.POST("/getAllPool", service.GetAllPool)
			sr.GET("/getPoolDetail", service.GetPoolDetail)
			sr.GET("/userPauseStake", service.UserPauseStake)
		}

		//代币相关接口
		token := lr.Group("/token")
		token.Use(middleware.Auth())
		{
			token.POST("/userTransfer", service.UserTransfer)
			token.GET("/tokenBalance", service.TokenBalance)
		}

		//代币相关接口
		transLog := lr.Group("/trans/logs")
		transLog.Use(middleware.Auth())
		{
			transLog.POST("/queryTransLogs", service.QueryTransLogs)
		}

	}

	return r
}

package router

import (
	"github.com/gin-gonic/gin"
	"github.com/strong-defi/strong-defi-backend/internal/middleware"
	"github.com/strong-defi/strong-defi-backend/internal/service"
)

func InjectUserRouter(r *gin.Engine, service *service.UserService) {
	e := r.Group("/v1/user", middleware.CheckToken)
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
}

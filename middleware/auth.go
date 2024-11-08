package middleware

import (
	"github.com/gin-gonic/gin"
	"strings"
	API "strong-defi-backend/common"
	"strong-defi-backend/service"
	"strong-defi-backend/utils"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		myCtx := &service.CustomContext{Context: c}
		url := myCtx.Request.URL.String()
		if strings.Contains(url, "/login") {
			return
		}
		//校验token

		token := myCtx.GetHeader("Authorization")

		err := utils.ValidateToken(token)

		if err != nil {
			myCtx.JSON(API.NO_AUTH_ERROR, "")
			return
		}
		info, _ := utils.GetUserInfo(token)
		myCtx.Set(API.TOKEN_KEY, info)
		myCtx.Next()
	}
}

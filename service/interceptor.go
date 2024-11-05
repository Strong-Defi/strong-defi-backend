package service

import (
	"github.com/gin-gonic/gin"
	"strings"
	API "strong-defi-backend/common"
	"strong-defi-backend/utils"
)

func CheckIsAdmin(c *gin.Context) {

	myCtx := &CustomContext{c}
	myCtx.GetHeader("")

}

// CheckToken 检查是否具有登陆权限
func CheckToken(c *gin.Context) {

	myCtx := &CustomContext{c}
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
}

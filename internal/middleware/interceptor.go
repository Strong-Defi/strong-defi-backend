package middleware

import (
	"github.com/gin-gonic/gin"
	API "github.com/strong-defi/strong-defi-backend/common"
	"github.com/strong-defi/strong-defi-backend/internal/service"
	"github.com/strong-defi/strong-defi-backend/pkg/authentication"
	"strings"
)

func CheckIsAdmin(c *gin.Context) {

	myCtx := &service.CustomContext{c}
	myCtx.GetHeader("")

}

// CheckToken 检查是否具有登陆权限
func CheckToken(c *gin.Context) {

	myCtx := &service.CustomContext{c}
	url := myCtx.Request.URL.String()
	if strings.Contains(url, "/login") {
		return
	}
	//校验token

	token := myCtx.GetHeader("Authorization")

	err := authentication.ValidateToken(token)

	if err != nil {
		myCtx.JSON(API.NO_AUTH_ERROR, "")
		return
	}
	info, _ := authentication.GetUserInfo(token)

	myCtx.Set(API.TOKEN_KEY, info)
}

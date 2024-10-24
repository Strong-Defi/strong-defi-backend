package service

import "github.com/gin-gonic/gin"

func CheckIsAdmin(c *gin.Context) {

	myCtx := &CustomContext{c}
	myCtx.GetHeader("")

}

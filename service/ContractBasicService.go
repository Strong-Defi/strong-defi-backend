package service

import (
	"github.com/gin-gonic/gin"
	API "strong-defi-backend/common"
	"strong-defi-backend/model"
	"strong-defi-backend/req"
)

func GetContractDetail(c *gin.Context) {

	//切换自定义上下文
	myCtx := &CustomContext{c}

	/*查询数据库*/
	user, _ := model.SelectUser(dao.ORM(), 23)

	//
	//return user
	/*获取数据，然后返回*/
	myCtx.JSON(API.SUCCESS, user)
}

func BasicDemo(c *gin.Context) {
	myCtx := &CustomContext{c}
	var demo req.BasicDemo // 创建绑定用的结构体
	err := c.ShouldBindJSON(&demo)
	/*数据校验*/
	if errorMsg := dataValidate(demo); errorMsg != nil {
		logs.Error("Error binding JSON:", errorMsg)
		myCtx.JSON(API.DATA_VALIDATE_ERROR, "")
		return
	}
	// 绑定 JSON 请求体
	if err != nil {
		// 处理绑定错误
		logs.Error("Error binding JSON:", err.Error())
		myCtx.JSON(API.DATA_ERROR, err.Error())
		return
	}

	// 处理成功
	myCtx.JSON(API.SUCCESS, demo)
}

//切换自定义上下文

//basicDemo := c.ShouldBindJSON(&req.BasicDemo{})
//if basicDemo.Error() != "" {
//fmt.Println(basicDemo.Error())
//}
//fmt.Println(basicDemo)
//myCtx.JSON(API.SUCCESS, basicDemo)

package service

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	API "strong-defi-backend/common"
	"strong-defi-backend/model"
	"strong-defi-backend/req"
	"strong-defi-backend/utils"
)

func UserLogin(c *gin.Context) {

	myCtx := &CustomContext{c}
	var userLoginReq req.UserLoginReq
	if err := myCtx.ShouldBindJSON(&userLoginReq); err != nil {
		logs.Error("Error binding JSON:", err.Error())
		myCtx.JSON(API.DATA_VALIDATE_ERROR, "")
		return
	}
	if errorMsg := dataValidate(userLoginReq); errorMsg != nil {
		logs.Error("入参校验失败:", errorMsg)
		myCtx.JSON(API.DATA_ERROR, "")
		return
	}

	scUser, _ := dao.SelectUserByWalletAddress(userLoginReq.WalletAddress)

	if scUser == nil {
		scUser = &model.ScUser{}
		scUser.UserPassword = "123456"
		scUser.UserUID = "123456"
		scUser.UserWalletAddress = userLoginReq.WalletAddress
		scUser.UserName = utils.Generate16CharString("123")
		err := dao.ORM().Transaction(func(tx *gorm.DB) (err error) {
			dTx := NewDao(tx)
			err = dTx.SaveScUser(scUser)
			if err != nil {
				return err
			}
			return nil
		})
		if err != nil {
			return
		}
		/*转jwt*/
		str, _ := utils.ToJson(scUser)

		//var userInfo *model.ScUser
		userInfo1 := new(model.ScUser)
		err = json.Unmarshal([]byte(str), userInfo1)
		myCtx.JSON(API.SUCCESS, str)
	} else {

	}

}

func UserRegister(c *gin.Context) {

}

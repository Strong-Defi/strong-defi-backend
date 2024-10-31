package service

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
	API "strong-defi-backend/common"
	"strong-defi-backend/model"
	"strong-defi-backend/request"
	"strong-defi-backend/utils"
)

func UserLogin(c *gin.Context) {

	myCtx := &CustomContext{c}
	var userLoginReq request.UserLoginReq
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
		token, _ := utils.CreateToken(str)

		myCtx.JSON(API.SUCCESS, token)
	} else {
		str, _ := utils.ToJson(scUser)

		token, _ := utils.CreateToken(str)
		myCtx.JSON(API.SUCCESS, token)
	}

}

// QueryUserIsExist 查询用户是否已经注册
func QueryUserIsExist(c *gin.Context) {

	myCtx := &CustomContext{c}
	userName := myCtx.Query("userName")

	user, _ := dao.SelectUser("user_name=? and is_deleted=0", userName)

	if user == nil {
		myCtx.JSON(API.SUCCESS, "")
	} else {
		myCtx.JSON(API.USER_EXIST, "用户名已存在")
	}
}

func QueryByWalletAddress(c *gin.Context) {
	myCtx := &CustomContext{c}

	walletAddress := myCtx.Query("walletAddress")

	logs.Warn("通过地址查询用户，入参为：", walletAddress)

	user, _ := dao.SelectUser("user_wallet_address=? and is_deleted=0", walletAddress)

	if user == nil {
		myCtx.JSON(API.SUCCESS, "")
	} else {
		myCtx.JSON(API.USER_EXIST, "该地址已创建用户")
	}
}

func QueryByTelephone(c *gin.Context) {
	myCtx := &CustomContext{c}

	telephone := myCtx.Query("telephone")
	logs.Warn("通过电话号码用户，入参为：", telephone)
	user, _ := dao.SelectUser("user_telephone=? and is_deleted=0", telephone)

	if user == nil {
		myCtx.JSON(API.SUCCESS, "")
	} else {
		myCtx.JSON2(API.USER_EXIST, "该电话已创建用户")
	}
}

func UserRegister(c *gin.Context) {
	myCtx := &CustomContext{c}
	var userRegister request.UserRegister
	if err := myCtx.ShouldBindJSON(&userRegister); err != nil {
		logs.Error("Error binding JSON:", err.Error())
		myCtx.JSON(API.DATA_VALIDATE_ERROR, "")
		return
	}
	if errorMsg := dataValidate(userRegister); errorMsg != nil {
		logs.Error("入参校验失败:", errorMsg)
		myCtx.JSON(API.DATA_ERROR, "")
		return
	}

	var scUser model.ScUser
	err := copier.Copy(&scUser, &userRegister)

	if err != nil {
		logs.Error("数据转化异常。", err)
		myCtx.JSON(API.DATA_ERROR, "服务异常")
		return
	}
	/*首先查询有没有用户，如果有就返回，没有就创建*/
	user, _ := dao.SelectUser(" (user_telephone=? or user_wallet_address=? or user_name=? ) and is_deleted=0 ",
		scUser.UserTelephone, scUser.UserWalletAddress, scUser.UserName)

	if user != nil {
		myCtx.JSON(API.USER_EXIST, "")
		return
	}

	scUser.UserUID = utils.Generate16CharString("123")

	err = dao.SaveScUser(&scUser)

	if err != nil {
		logs.Error("用户保存失败。", err)
		myCtx.JSON(API.DATA_ERROR, "")
		return
	} else {
		myCtx.JSON(API.SUCCESS, "")
	}
}

func ModifyUser(c *gin.Context) {

	myCtx := &CustomContext{c}
	var modifyUser request.ModifyUserReq
	if err := myCtx.ShouldBindJSON(&modifyUser); err != nil {
		logs.Error("Error binding JSON:", err.Error())
		myCtx.JSON(API.DATA_VALIDATE_ERROR, "")
		return
	}
	if errorMsg := dataValidate(modifyUser); errorMsg != nil {
		logs.Error("入参校验失败:", errorMsg)
		myCtx.JSON(API.DATA_ERROR, "")
		return
	}

	logs.Info("修改用户信息入参为：", modifyUser)

	user, _ := dao.SelectUser("id=? and is_deleted=0", modifyUser.UserId)

	if user == nil {
		myCtx.JSON2(API.DATA_ERROR, "用户不存在")
	}
	copier.Copy(user, &modifyUser)
	updateValue := map[string]interface{}{}
	if modifyUser.ModifyType == nil {
		if len(modifyUser.UserName) > 0 {
			updateValue["user_name"] = modifyUser.UserName
		} else if len(modifyUser.UserTelephone) > 0 {
			updateValue["user_telephone"] = modifyUser.UserTelephone
		} else if len(modifyUser.UserAddress) > 0 {
			updateValue["user_address"] = modifyUser.UserAddress
		} else if len(modifyUser.UserCountry) > 0 {
			updateValue["user_country"] = modifyUser.UserCountry
		} else if len(modifyUser.UserEmail) > 0 {
			updateValue["user_email"] = modifyUser.UserEmail
		}
		if len(updateValue) <= 0 {
			myCtx.JSON2(API.COMMOM_ERROR, "用户没有内容需要更新")
			return
		}
	} else if *modifyUser.ModifyType == 1 {
		if len(modifyUser.UserTelephone) > 0 {
			updateValue["user_telephone"] = modifyUser.UserTelephone
		} else {
			myCtx.JSON2(API.COMMOM_ERROR, "修改电话号码入参错误")
			return
		}
	} else if *modifyUser.ModifyType == 2 {
		if len(modifyUser.UserEmail) > 0 {
			updateValue["user_email"] = modifyUser.UserEmail
		} else {
			myCtx.JSON2(API.COMMOM_ERROR, "修改邮箱入参错误")
			return
		}
	} else if *modifyUser.ModifyType == 3 {
		if len(modifyUser.UserWalletAddress) > 0 {
			updateValue["user_wallet_address"] = modifyUser.UserWalletAddress
		} else {
			myCtx.JSON2(API.COMMOM_ERROR, "修改钱包地址入参错误")
			return
		}
	} else if *modifyUser.ModifyType == 4 {
		if len(modifyUser.UserWalletAddress) > 0 {
			updateValue["user_password"] = modifyUser.UserPassword
		} else {
			myCtx.JSON2(API.COMMOM_ERROR, "修改密码入参错误")
			return
		}
	} else {
		myCtx.JSON2(API.COMMOM_ERROR, "修改类型错误")
		return
	}

	rowNum, err := dao.Update("id=?", updateValue, modifyUser.UserId)
	if err != nil || rowNum == 0 {
		logs.Error("用户修改信息失败，", err)
		myCtx.JSON2(API.COMMOM_ERROR, "用户修改信息失败")
		return
	}
	myCtx.JSON(API.SUCCESS, "")
}

// DeleteUser 注销用户
func DeleteUser(c *gin.Context) {
	myCtx := &CustomContext{c}
	value, exists := myCtx.Get(API.TOKEN_KEY)
	var userInfo model.ScUser
	if exists {
		err := json.Unmarshal([]byte(value.(string)), &userInfo)
		if err != nil {
			logs.Error("解析用户信息失败。", err)
			myCtx.JSON(API.COMMOM_ERROR, "解析用户信息失败")
			return
		}
	}

	userId := myCtx.Query("uid")
	logs.Info("注销用户，入参为：", userId)

	if userId != userInfo.UserUID {
		myCtx.JSON(API.COMMOM_ERROR, "无权注销别人的账号")
		return
	}
	err := dao.DeleteScUser("user_uid=?", userId)

	if err != nil {
		logs.Error("注销失败。", err)
		myCtx.JSON2(API.COMMOM_ERROR, "注销失败")
		return
	}
	myCtx.JSON(API.SUCCESS, "")
}

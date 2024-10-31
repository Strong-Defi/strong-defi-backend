package service

import (
	"encoding/json"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	API "github.com/strong-defi/strong-defi-backend/common"
	"github.com/strong-defi/strong-defi-backend/internal/dao"
	"github.com/strong-defi/strong-defi-backend/internal/model"
	"github.com/strong-defi/strong-defi-backend/internal/req"
	"github.com/strong-defi/strong-defi-backend/pkg/authentication"
	string2 "github.com/strong-defi/strong-defi-backend/pkg/string"
)

type UserService struct {
	dao *dao.Dao
}

func NewUserService(d *dao.Dao) *UserService {
	return &UserService{dao: d}
}

func (svc *UserService) UserLogin(c *gin.Context) {
	var err error
	myCtx := &Context{c}
	var userLoginReq req.UserLoginReq
	if err := myCtx.ShouldBindJSON(&userLoginReq); err != nil {
		log.Error("Error binding JSON:", err.Error())
		myCtx.JSON(API.DATA_VALIDATE_ERROR, "")
		return
	}
	if errorMsg := dataValidate(userLoginReq); errorMsg != nil {
		log.Error("入参校验失败:", errorMsg)
		myCtx.JSON(API.DATA_ERROR, "")
		return
	}

	scUser, _ := svc.dao.SelectUserByWalletAddress(userLoginReq.WalletAddress)

	if scUser == nil {
		scUser = &model.ScUser{}
		scUser.UserPassword = "123456"
		scUser.UserUID = "123456"
		scUser.UserWalletAddress = userLoginReq.WalletAddress
		scUser.UserName = string2.Generate16CharString("123")
		err = svc.dao.SaveScUser(scUser)
		if err != nil {
			log.Error("用户保存失败。", err)
			return
		}
		/*转jwt*/
		str, _ := string2.ToJson(scUser)
		token, _ := authentication.CreateToken(str)

		myCtx.JSON(API.SUCCESS, token)
	} else {
		str, _ := string2.ToJson(scUser)

		token, _ := authentication.CreateToken(str)
		myCtx.JSON(API.SUCCESS, token)
	}

}

// QueryUserIsExist 查询用户是否已经注册
func (svc *UserService) QueryUserIsExist(c *gin.Context) {

	myCtx := &Context{c}
	userName := myCtx.Query("userName")

	user, _ := svc.dao.SelectUser("user_name=? and is_deleted=0", userName)

	if user == nil {
		myCtx.JSON(API.SUCCESS, "")
	} else {
		myCtx.JSON(API.USER_EXIST, "用户名已存在")
	}
}

func (svc *UserService) QueryByWalletAddress(c *gin.Context) {
	myCtx := &Context{c}

	walletAddress := myCtx.Query("walletAddress")

	log.Warn("通过地址查询用户，入参为：", walletAddress)

	user, _ := svc.dao.SelectUser("user_wallet_address=? and is_deleted=0", walletAddress)

	if user == nil {
		myCtx.JSON(API.SUCCESS, "")
	} else {
		myCtx.JSON(API.USER_EXIST, "该地址已创建用户")
	}
}

func (svc *UserService) QueryByTelephone(c *gin.Context) {
	myCtx := &Context{c}

	telephone := myCtx.Query("telephone")
	log.Warn("通过电话号码用户，入参为：", telephone)
	user, _ := svc.dao.SelectUser("user_telephone=? and is_deleted=0", telephone)

	if user == nil {
		myCtx.JSON(API.SUCCESS, "")
	} else {
		myCtx.JSON2(API.USER_EXIST, "该电话已创建用户")
	}
}

func (svc *UserService) UserRegister(c *gin.Context) {
	myCtx := &Context{c}
	var userRegister req.UserRegister
	if err := myCtx.ShouldBindJSON(&userRegister); err != nil {
		log.Error("Error binding JSON:", err.Error())
		myCtx.JSON(API.DATA_VALIDATE_ERROR, "")
		return
	}
	if errorMsg := dataValidate(userRegister); errorMsg != nil {
		log.Error("入参校验失败:", errorMsg)
		myCtx.JSON(API.DATA_ERROR, "")
		return
	}

	var scUser model.ScUser
	err := copier.Copy(&scUser, &userRegister)

	if err != nil {
		log.Error("数据转化异常。", err)
		myCtx.JSON(API.DATA_ERROR, "服务异常")
		return
	}
	/*首先查询有没有用户，如果有就返回，没有就创建*/
	user, _ := svc.dao.SelectUser(" (user_telephone=? or user_wallet_address=? or user_name=? ) and is_deleted=0 ",
		scUser.UserTelephone, scUser.UserWalletAddress, scUser.UserName)

	if user != nil {
		myCtx.JSON(API.USER_EXIST, "")
		return
	}

	scUser.UserUID = string2.Generate16CharString("123")

	err = svc.dao.SaveScUser(&scUser)

	if err != nil {
		log.Error("用户保存失败。", err)
		myCtx.JSON(API.DATA_ERROR, "")
		return
	} else {
		myCtx.JSON(API.SUCCESS, "")
	}
}

func (svc *UserService) ModifyUser(c *gin.Context) {

	myCtx := &Context{c}
	var modifyUser req.ModifyUserReq
	if err := myCtx.ShouldBindJSON(&modifyUser); err != nil {
		log.Error("Error binding JSON:", err.Error())
		myCtx.JSON(API.DATA_VALIDATE_ERROR, "")
		return
	}
	if errorMsg := dataValidate(modifyUser); errorMsg != nil {
		log.Error("入参校验失败:", errorMsg)
		myCtx.JSON(API.DATA_ERROR, "")
		return
	}

	log.Info("修改用户信息入参为：", modifyUser)

	user, _ := svc.dao.SelectUser("id=? and is_deleted=0", modifyUser.UserId)

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

	rowNum, err := svc.dao.Update("id=?", updateValue, modifyUser.UserId)
	if err != nil || rowNum == 0 {
		log.Error("用户修改信息失败，", err)
		myCtx.JSON2(API.COMMOM_ERROR, "用户修改信息失败")
		return
	}
	myCtx.JSON(API.SUCCESS, "")
}

// DeleteUser 注销用户
func (svc *UserService) DeleteUser(c *gin.Context) {
	myCtx := &Context{c}
	value, exists := myCtx.Get(API.TOKEN_KEY)
	var userInfo model.ScUser
	if exists {
		err := json.Unmarshal([]byte(value.(string)), &userInfo)
		if err != nil {
			log.Error("解析用户信息失败。", err)
			myCtx.JSON(API.COMMOM_ERROR, "解析用户信息失败")
			return
		}
	}

	userId := myCtx.Query("uid")
	log.Info("注销用户，入参为：", userId)

	if userId != userInfo.UserUID {
		myCtx.JSON(API.COMMOM_ERROR, "无权注销别人的账号")
		return
	}
	err := svc.dao.DeleteScUser("user_uid=?", userId)

	if err != nil {
		log.Error("注销失败。", err)
		myCtx.JSON2(API.COMMOM_ERROR, "注销失败")
		return
	}
	myCtx.JSON(API.SUCCESS, "")
}
func (svc *UserService) GetContractInfo(ctx *gin.Context) {

	log.Info("Info to Ethereum")
	log.Error("Error to Ethereum")
	log.Debug("Debug to Ethereum")
	log.Warn("Warn to Ethereum")
	client, _ := ethclient.Dial("https://cloudflare-eth.com")
	if client == nil {
		log.Error("Connected to Ethereum")
	}
}

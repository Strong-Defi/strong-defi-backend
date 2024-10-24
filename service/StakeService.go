package service

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	API "strong-defi-backend/common"
	"strong-defi-backend/req"
	"strong-defi-backend/utils"
)

//func GetContractInfo(c *gin.Context) {
//	logs.Info("Info to Ethereum")
//	logs.Error("Error to Ethereum")
//	logs.Debug("Debug to Ethereum")
//	logs.Warn("Warn to Ethereum")
//	client, _ := ethclient.Dial("https://cloudflare-eth.com")
//	if client == nil {
//		logs.Error("Connected to Ethereum")
//	}
//}

// 获取账号余额
func GetAccountBalance(c *gin.Context) {
	myCtx := &CustomContext{c}

	var accountReq req.AccountBalanceReq
	err := myCtx.ShouldBindJSON(&accountReq)

	if err != nil {
		logs.Error("Error binding JSON:", err.Error())
		myCtx.JSON(API.DataError, err.Error())
		return
	}

	if errorMsg := dataValidate(accountReq); errorMsg != nil {
		logs.Error("入参校验失败:", errorMsg)
		myCtx.JSON(API.RequestParseError, "")
		return
	}

	/*获取连接*/
	client, err1 := ethclient.Dial(accountReq.Url)
	if err1 != nil {
		logs.Error("获取以太坊节点失败:", err.Error())
		myCtx.JSON(API.DataError, err.Error())
		return
	}

	address := common.HexToAddress(accountReq.Address)

	at, _ := client.BalanceAt(context.Background(), address, nil)

	float2 := utils.ConvertToMathFloat2(at, 18, 2)

	myCtx.JSON(API.SUCCESS, float2)

}

func AddPool(c *gin.Context) {
	myCtx := &CustomContext{c}

	var accountReq req.AccountBalanceReq
	err := myCtx.ShouldBindJSON(&accountReq)

	if err != nil {
		logs.Error("Error binding JSON:", err.Error())
		myCtx.JSON(API.DataError, err.Error())
		return
	}

	if errorMsg := dataValidate(accountReq); errorMsg != nil {
		logs.Error("入参校验失败:", errorMsg)
		myCtx.JSON(API.RequestParseError, "")
		return
	}

	/*获取连接*/
	client, err1 := ethclient.Dial(accountReq.Url)
	if err1 != nil {
		logs.Error("获取以太坊节点失败:", err.Error())
		myCtx.JSON(API.DataError, err.Error())
		return
	}

	address := common.HexToAddress(accountReq.Address)

	at, _ := client.BalanceAt(context.Background(), address, nil)

	float2 := utils.ConvertToMathFloat2(at, 18, 2)

	myCtx.JSON(API.SUCCESS, float2)

}

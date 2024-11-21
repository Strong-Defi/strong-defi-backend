package service

import (
	"context"
	"encoding/json"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"math/big"
	API "strong-defi-backend/common"
	sc_token "strong-defi-backend/contract/scToken"
	"strong-defi-backend/model"
	"strong-defi-backend/pkg/log"
	"strong-defi-backend/request"
	"strong-defi-backend/utils"
)

// UserTransfer 用户本地代币转账
func UserTransfer(c *gin.Context) {

	myCtx := &CustomContext{c}

	var req request.UserTransferReq

	err := myCtx.ShouldBindJSON(&req)

	if err != nil {
		log.Error().Msgf("Error binding JSON:(%+v)", err.Error())
		myCtx.JSON(API.DATA_ERROR, err.Error())
		return
	}

	if errorMsg := dataValidate(req); errorMsg != nil {
		log.Error().Msgf("入参校验失败:(%+v)", errorMsg)
		myCtx.JSON(API.DATA_VALIDATE_ERROR, "")
		return
	}
	log.Info().Msgf("用户本地代币转账入参为：(%+v)", req)

	scUser := &model.ScUser{}
	//获取登陆人信息
	value, exists := myCtx.Get(API.TOKEN_KEY)
	if exists {
		_ = json.Unmarshal([]byte(value.(string)), scUser)
	} else {
		myCtx.JSON2(API.COMMOM_ERROR, "请登陆后在操作")
		return
	}
	if scUser.UserWalletAddress != req.SendAddress {
		log.Warn().Msgf("用户地址不匹配,%+v", req)
		myCtx.JSON2(API.COMMOM_ERROR, "用户地址不匹配")
		return
	}
	//建立链接
	client, err := ethclient.Dial(app.DeployAddress)

	if err != nil {
		log.Error().Msgf("以太坊连接失败。(%+v)", err)
		myCtx.JSON2(API.COMMOM_ERROR, "以太坊连接失败")
		return
	}
	token, err := sc_token.NewSCTToken(common.HexToAddress(req.SendAddress), client)

	balance, _ := token.BalanceOf(nil, common.HexToAddress(req.SendAddress))

	if balance.Int64() < req.Amount {
		log.Error().Msg("用户余额不足")
		myCtx.JSON2(API.COMMOM_ERROR, "用户余额不足")
		return
	}
	//金额授权，这里的opt是nil因为不需要传入对应的参数，如果想要控制gasPrice或者限制，可以设置
	//opts中value是用于发送eth的数量
	approveTx, err := token.Approve(nil, common.HexToAddress(req.ReceiveAddress), new(big.Int).SetInt64(req.Amount))
	if err != nil {
		log.Error().Msgf("授权失败:(%+v)", err)
		myCtx.JSON2(API.COMMOM_ERROR, "代币授权失败")
		return
	}

	receipt, _ := client.TransactionReceipt(context.Background(), approveTx.Hash())

	if receipt.Status == types.ReceiptStatusSuccessful {
		//保存交易信息
		saveTransLogs(receipt, approveTx, common.HexToAddress(req.SendAddress), scUser, API.UserApprove.String())
	} else {
		log.Error().Msgf("授权交易失败:(%+v)", receipt.Status)
		myCtx.JSON2(API.COMMOM_ERROR, "授权交易失败")
		return
	}
	//代币转账
	transferTx, err := token.Transfer(nil, common.HexToAddress(req.ReceiveAddress), new(big.Int).SetInt64(req.Amount))

	if err != nil {
		log.Error().Msgf("代币转账失败:(%+v)", err)
		myCtx.JSON2(API.COMMOM_ERROR, "代币转账失败")
		return
	}

	transactionReceipt, _ := client.TransactionReceipt(context.Background(), transferTx.Hash())
	if transactionReceipt.Status == types.ReceiptStatusSuccessful {
		//保存交易信息
		saveTransLogs(transactionReceipt, transferTx, common.HexToAddress(req.SendAddress), scUser, API.UserTransfer.String())
	} else {
		log.Error().Msgf("授权交易失败:(%+v)", receipt.Status)
		myCtx.JSON2(API.COMMOM_ERROR, "授权交易失败")
		return
	}

	myCtx.JSON(API.SUCCESS, "")
}

// TokenBalance 查询代币余额
func TokenBalance(c *gin.Context) {
	myCtx := &CustomContext{c}

	//获取用户信息
	scUser := &model.ScUser{}
	//获取登陆人信息
	value, exists := myCtx.Get(API.TOKEN_KEY)
	if exists {
		_ = json.Unmarshal([]byte(value.(string)), scUser)
	} else {
		myCtx.JSON2(API.COMMOM_ERROR, "请登陆后在操作")
		return
	}
	//建立链接
	client, err := ethclient.Dial(app.DeployAddress)

	if err != nil {
		log.Error().Msgf("以太坊连接失败。(%+v)", err)
		myCtx.JSON2(API.COMMOM_ERROR, "以太坊连接失败")
		return
	}
	token, err := sc_token.NewSCTToken(common.HexToAddress(scUser.UserWalletAddress), client)

	if err != nil {
		log.Error().Msgf("连接合约失败。(%+v)", err)
		myCtx.JSON2(API.COMMOM_ERROR, "连接合约失败")
		return
	}

	balance, _ := token.BalanceOf(nil, common.HexToAddress(scUser.UserWalletAddress))

	decimals, _ := token.Decimals(nil)

	//进行数据转换，例：将1000000000转成1.00
	float2 := utils.ConvertToMathFloat2(balance, int(decimals), 2)
	//获取合约数据
	log.Info().Msgf("用户余额为：(%+v)", float2)
	//这里要看代币自己的精度
	myCtx.JSON(API.SUCCESS, float2)

}

package service

import (
	"context"
	"encoding/json"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
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

// UserApprove 用户转账授权
func UserApprove(c *gin.Context) {

	myCtx := &CustomContext{c}

	var req request.UserApproveReq

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

	if req.Amount <= 0 {
		log.Info().Msgf("转账数量必须大于0")
		myCtx.JSON2(API.DATA_ERROR, "转账数量必须大于0")
		return
	}

	value, _ := myCtx.Get(API.TOKEN_KEY)

	var scUser model.ScUser
	err = json.Unmarshal([]byte(value.(string)), &scUser)
	if err != nil {
		log.Error().Msgf("解析用户信息失败。(%+v)", err)
		myCtx.JSON2(API.COMMOM_ERROR, "解析用户信息失败")
		return
	}
	//用户授权，优先查询用户本地代币数量
	scTokenAddress := common.HexToAddress(app.TokenAddress)
	client, err := ethclient.Dial(app.DeployAddress)
	if err != nil {
		log.Error().Msgf("连接合约失败。(%+v)", err)
		myCtx.JSON2(API.COMMOM_ERROR, "连接合约失败")
		return
	}

	at, _ := client.PendingNonceAt(context.Background(), scTokenAddress)

	scToken, err := sc_token.NewSCTToken(scTokenAddress, client)
	//获取chaiID
	chaiID, err := client.ChainID(context.Background())
	//获取交易签名函数
	ecdsa, _ := crypto.HexToECDSA(req.UserPrivateKey)

	//创建签名函数
	signer := utils.CreateSigner(ecdsa, chaiID)

	opts := &bind.TransactOpts{
		From: common.HexToAddress(scUser.UserWalletAddress),
		//设置nonce值进行防重处理
		Nonce: new(big.Int).SetUint64(at),
		//这个是给合约发送的值，不是转的值
		Value:  big.NewInt(0),
		Signer: signer,
	}
	approve, err := scToken.Approve(opts, common.HexToAddress(scUser.UserWalletAddress), big.NewInt(req.Amount))
	if err != nil {
		log.Error().Msgf("授权失败。(%+v)", err)
		myCtx.JSON2(API.COMMOM_ERROR, "授权失败")
		return
	}
	receipt, _ := client.TransactionReceipt(context.Background(), approve.Hash())

	//保存交易数据
	go func() {
		saveTransLogs(receipt, approve, scTokenAddress, &scUser, API.UserStake.String())
	}()
	if receipt.Status == types.ReceiptStatusSuccessful {
		myCtx.JSON(API.SUCCESS, "")
		return
	}
	log.Error().Msgf("授权状态失败。(%+v)", receipt.Logs)
	myCtx.JSON2(API.COMMOM_ERROR, "授权状态失败")
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

package service

import (
	"context"
	"github.com/ethereum/go-ethereum/crypto"

	"encoding/json"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"math/big"
	API "strong-defi-backend/common"
	contract "strong-defi-backend/contract/schStake"
	"strong-defi-backend/model"
	"strong-defi-backend/req"
	"strong-defi-backend/utils"
)

// GetAccountBalance 获取账号余额
func GetAccountBalance(c *gin.Context) {
	myCtx := &CustomContext{c}

	var accountReq req.AccountBalanceReq
	err := myCtx.ShouldBindJSON(&accountReq)

	if err != nil {
		logs.Error("Error binding JSON:", err.Error())
		myCtx.JSON(API.DATA_ERROR, err.Error())
		return
	}

	if errorMsg := dataValidate(accountReq); errorMsg != nil {
		logs.Error("入参校验失败:", errorMsg)
		myCtx.JSON(API.DATA_VALIDATE_ERROR, "")
		return
	}

	/*获取连接*/
	client, err1 := ethclient.Dial(accountReq.Url)
	if err1 != nil {
		logs.Error("获取以太坊节点失败:", err.Error())
		myCtx.JSON(API.DATA_ERROR, err.Error())
		return
	}

	address := common.HexToAddress(accountReq.Address)

	at, _ := client.BalanceAt(context.Background(), address, nil)

	float2 := utils.ConvertToMathFloat2(at, 18, 2)

	myCtx.JSON(API.SUCCESS, float2)

}

// AddPool 添加用户池
func AddPool(c *gin.Context) {
	myCtx := &CustomContext{c}

	//判断是否是管理员
	scUser := new(model.ScUser)
	value, exists := myCtx.Get(API.TOKEN_KEY)
	if exists {
		json.Unmarshal([]byte(value.(string)), scUser)
	} else {
		myCtx.JSON2(API.COMMOM_ERROR, "请登陆后在操作")
		return
	}

	//如果不是管理员，不允许添加质押池
	if scUser.UserWalletAddress != app.AdminAddress {
		myCtx.JSON2(API.COMMOM_ERROR, "非管理员不允许操作质押池")
		return
	}

	var addPoolReq req.AddPoolReq
	err := myCtx.ShouldBindJSON(&addPoolReq)

	if err != nil {
		logs.Error("Error binding JSON:", err.Error())
		myCtx.JSON(API.DATA_ERROR, err.Error())
		return
	}

	if errorMsg := dataValidate(addPoolReq); errorMsg != nil {
		logs.Error("入参校验失败:", errorMsg)
		myCtx.JSON(API.DATA_VALIDATE_ERROR, "")
		return
	}
	//建立链接
	client, err := ethclient.Dial(app.DeployAddress)

	if err != nil {
		logs.Error("以太坊连接失败。", err)
		myCtx.JSON2(API.COMMOM_ERROR, "以太坊连接失败")
		return
	}
	address := common.HexToAddress(scUser.UserWalletAddress)
	//获取nonce值
	nonce, err := client.PendingNonceAt(context.Background(), address)

	if err != nil {
		logs.Error("获取nonce失败。", err)
		myCtx.JSON2(API.COMMOM_ERROR, "获取nonce失败")
		return
	}
	newContract, _ := contract.NewContract(address, client)

	//获取chaiID
	chaiID, err := client.ChainID(context.Background())
	//获取交易签名函数
	ecdsa, _ := crypto.HexToECDSA(addPoolReq.UserPrivateKey)
	//创建签名函数
	signer := utils.CreateSigner(ecdsa, chaiID)
	//token地址
	tokenAddress := common.HexToAddress(addPoolReq.TokenAddress)

	opts := &bind.TransactOpts{
		GasLimit: utils.GetDefaultValueInt(addPoolReq.GasLimit == 0, addPoolReq.GasLimit, 21000),
		From:     address,
		//设置nonce值进行防重处理
		Nonce:  new(big.Int).SetInt64(int64(nonce)),
		Signer: signer,
	}
	tx, _ := newContract.AddPool(opts, new(big.Int).SetInt64(int64(addPoolReq.Wight)),
		tokenAddress, new(big.Int).SetInt64(int64(addPoolReq.MinStakeAmount)),
		addPoolReq.IsUpdatePool, new(big.Int).SetInt64(int64(addPoolReq.UnStakeLockBlockNumber)))

	myCtx.JSON(API.SUCCESS, "")
}

// GetPoolBalance 获取用户池余额
func GetPoolBalance(c *gin.Context) {

}

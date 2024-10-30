package service

import (
	"context"
	"encoding/binary"
	"encoding/json"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"math/big"
	"strings"
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
	newContract, _ := contract.NewSCHStake(address, client)

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
	tx, err := newContract.AddPool(opts, new(big.Int).SetInt64(int64(addPoolReq.Wight)),
		tokenAddress, new(big.Int).SetInt64(int64(addPoolReq.MinStakeAmount)),
		addPoolReq.IsUpdatePool, new(big.Int).SetInt64(int64(addPoolReq.UnStakeLockBlockNumber)))

	//获取返回内容
	receipt, _ := client.TransactionReceipt(context.Background(), tx.Hash())

	if len(receipt.Logs) > 0 {
		//异步存入质押池信息
		go func() {
			for _, vlog := range receipt.Logs {

				parseABI, _ := abi.JSON(strings.NewReader(contract.SCHStakeMetaData.ABI))
				//这是固定的，获取事件信息的，第一个是合约事件返回的结构，第二个是事件名称
				var events []req.Event
				err2 := parseABI.UnpackIntoInterface(&events, "AddPool", vlog.Data)

				if err2 != nil {
					logs.Error("获取nonce失败。", err)
					return
				}
				pool := model.StakePool{
					Weight:               int64(addPoolReq.Wight),
					TokenAddress:         addPoolReq.TokenAddress,
					MinStakeAmount:       uint64(addPoolReq.MinStakeAmount),
					LockStakeBlockNumber: int64(addPoolReq.UnStakeLockBlockNumber),
					//不知道events[0].Value[:]啥意思
					PoolID: int64(binary.BigEndian.Uint64(events[0].Value[:])),
					//创建人，保存的是地址
					Creator: scUser.UserAddress,
				}
				//保存质押池信息
				_ = dao.CreateStakePool(&pool)
			}
		}()
	}
	if err != nil {
		logs.Error("添加质押池报错。", err)
		myCtx.JSON2(API.COMMOM_ERROR, "添加质押池报错")
		return
	}

	//获取交易收据

	myCtx.JSON(API.SUCCESS, "")
}

// GetPoolBalance 获取用户池余额
func GetPoolBalance(c *gin.Context) {

}

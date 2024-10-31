package service

import (
	"context"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/strong-defi/strong-defi-backend/internal/dao"
	"github.com/strong-defi/strong-defi-backend/internal/model"
	"github.com/strong-defi/strong-defi-backend/internal/req"
	"github.com/strong-defi/strong-defi-backend/pkg/math"
	string2 "github.com/strong-defi/strong-defi-backend/pkg/string"

	"encoding/json"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	API "github.com/strong-defi/strong-defi-backend/common"
	contract "github.com/strong-defi/strong-defi-backend/contract/schStake"
	"math/big"
)

type StakeService struct {
	dao *dao.Dao
}

func NewStakeService(d *dao.Dao) *StakeService {
	return &StakeService{dao: d}
}

// GetAccountBalance 获取账号余额
func (svc *StakeService) GetAccountBalance(c *gin.Context) {
	myCtx := &Context{c}
	var accountReq req.AccountBalanceReq
	err := myCtx.ShouldBindJSON(&accountReq)
	if err != nil {
		log.Error("Error binding JSON:", err.Error())
		myCtx.JSON(API.DATA_ERROR, err.Error())
		return
	}

	if errorMsg := dataValidate(accountReq); errorMsg != nil {
		log.Error("入参校验失败:", errorMsg)
		myCtx.JSON(API.DATA_VALIDATE_ERROR, "")
		return
	}

	/*获取连接*/
	client, err1 := ethclient.Dial(accountReq.Url)
	if err1 != nil {
		log.Error("获取以太坊节点失败:", err.Error())
		myCtx.JSON(API.DATA_ERROR, err.Error())
		return
	}

	address := common.HexToAddress(accountReq.Address)

	at, _ := client.BalanceAt(context.Background(), address, nil)

	float2 := math.ConvertToMathFloat2(at, 18, 2)

	myCtx.JSON(API.SUCCESS, float2)

}

// AddPool 添加用户池
func (svc *StakeService) AddPool(c *gin.Context) {
	myCtx := &Context{c}

	//判断是否是管理员
	scUser := new(model.ScUser)
	value, exists := myCtx.Get(API.TOKEN_KEY)
	if exists {
		json.Unmarshal([]byte(value.(string)), scUser)
	} else {
		myCtx.JSON2(API.COMMOM_ERROR, "请登陆后在操作")
		return
	}

	//如果不是管理员，不允许添加质押池 TODO
	//if scUser.UserWalletAddress != app.AdminAddress {
	//	myCtx.JSON2(API.COMMOM_ERROR, "非管理员不允许操作质押池")
	//	return
	//}

	var addPoolReq req.AddPoolReq
	err := myCtx.ShouldBindJSON(&addPoolReq)

	if err != nil {
		log.Error("Error binding JSON:", err.Error())
		myCtx.JSON(API.DATA_ERROR, err.Error())
		return
	}

	if errorMsg := dataValidate(addPoolReq); errorMsg != nil {
		log.Error("入参校验失败:", errorMsg)
		myCtx.JSON(API.DATA_VALIDATE_ERROR, "")
		return
	}
	//建立链接 TODO
	client, err := ethclient.Dial("")

	if err != nil {
		log.Error("以太坊连接失败。", err)
		myCtx.JSON2(API.COMMOM_ERROR, "以太坊连接失败")
		return
	}
	address := common.HexToAddress(scUser.UserWalletAddress)
	//获取nonce值
	nonce, err := client.PendingNonceAt(context.Background(), address)

	if err != nil {
		log.Error("获取nonce失败。", err)
		myCtx.JSON2(API.COMMOM_ERROR, "获取nonce失败")
		return
	}
	newContract, _ := contract.NewContract(address, client)

	//获取chaiID
	chaiID, err := client.ChainID(context.Background())
	//获取交易签名函数
	ecdsa, _ := crypto.HexToECDSA(addPoolReq.UserPrivateKey)
	//创建签名函数
	signer := string2.CreateSigner(ecdsa, chaiID)
	//token地址
	tokenAddress := common.HexToAddress(addPoolReq.TokenAddress)

	opts := &bind.TransactOpts{
		GasLimit: math.GetDefaultValueInt(addPoolReq.GasLimit == 0, addPoolReq.GasLimit, 21000),
		From:     address,
		//设置nonce值进行防重处理
		Nonce:  new(big.Int).SetInt64(int64(nonce)),
		Signer: signer,
	}
	tx, _ := newContract.AddPool(opts, new(big.Int).SetInt64(int64(addPoolReq.Wight)),
		tokenAddress, new(big.Int).SetInt64(int64(addPoolReq.MinStakeAmount)),
		addPoolReq.IsUpdatePool, new(big.Int).SetInt64(int64(addPoolReq.UnStakeLockBlockNumber)))
	log.Info("tx:", tx.Hash().Hex())
	myCtx.JSON(API.SUCCESS, "")
}

// GetPoolBalance 获取用户池余额
func (svc *StakeService) GetPoolBalance(c *gin.Context) {

}

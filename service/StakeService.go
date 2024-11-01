package service

import (
	"context"
	"encoding/binary"
	"encoding/json"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"math/big"
	"strings"
	API "strong-defi-backend/common"
	contract "strong-defi-backend/contract/schStake"
	"strong-defi-backend/model"
	"strong-defi-backend/request"
	"strong-defi-backend/utils"
)

// GetAccountBalance 获取账号余额
func GetAccountBalance(c *gin.Context) {
	myCtx := &CustomContext{c}

	var accountReq request.AccountBalanceReq
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

	var addPoolReq request.AddPoolReq
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
			if saveTransLogs(receipt, tx, address, scUser) {
				return
			}

			for _, vlog := range receipt.Logs {
				parseABI, _ := abi.JSON(strings.NewReader(contract.SCHStakeMetaData.ABI))
				//这是固定的，获取事件信息的，第一个是合约事件返回的结构，第二个是事件名称
				var events []request.Event
				err2 := parseABI.UnpackIntoInterface(&events, "AddPool", vlog.Data)

				if err2 != nil {
					logs.Error("获取nonce失败。", err)
					return
				}
				poolId := int64(binary.BigEndian.Uint64(events[0].Value[:]))
				pool := model.StakePool{
					Weight:               int64(addPoolReq.Wight),
					TokenAddress:         addPoolReq.TokenAddress,
					MinStakeAmount:       uint64(addPoolReq.MinStakeAmount),
					LockStakeBlockNumber: int64(addPoolReq.UnStakeLockBlockNumber),
					//不知道events[0].Value[:]啥意思
					PoolID: poolId,
					//生成唯一标识，用于在系统中流转
					PoolCode: utils.Generate16CharInt(poolId),
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

// UserStake 用户质押
func UserStake(c *gin.Context) {
	myCtx := &CustomContext{c}

	var req request.UserStakeReq

	err := myCtx.ShouldBindJSON(&req)

	if err != nil {
		logs.Error("Error binding JSON:", err.Error())
		myCtx.JSON(API.DATA_ERROR, err.Error())
		return
	}

	if errorMsg := dataValidate(req); errorMsg != nil {
		logs.Error("入参校验失败:", errorMsg)
		myCtx.JSON(API.DATA_VALIDATE_ERROR, "")
		return
	}
}

func GetUserInfo(c *gin.Context) {

	myCtx := &CustomContext{c}

	var req request.QueryUserInfoReq

	err := myCtx.ShouldBindJSON(&req)

	if err != nil {
		logs.Error("Error binding JSON:", err.Error())
		myCtx.JSON(API.DATA_ERROR, err.Error())
		return
	}

	if errorMsg := dataValidate(req); errorMsg != nil {
		logs.Error("入参校验失败:", errorMsg)
		myCtx.JSON(API.DATA_VALIDATE_ERROR, "")
		return
	}

	//建立连接
	dial, _ := ethclient.Dial(app.DeployAddress)
	//获取合约地址
	address := common.HexToAddress(app.ContractAddress)

	/*获取合约实例*/
	instance, err := contract.NewSCHStake(address, dial)

	//获取登陆人信息
	scUser := new(model.ScUser)
	value, exists := myCtx.Get(API.TOKEN_KEY)
	if exists {
		json.Unmarshal([]byte(value.(string)), scUser)
	} else {
		myCtx.JSON2(API.COMMOM_ERROR, "请登陆后在操作")
		return
	}
	logs.Info("获取用户质押信息")

	//通过poolCode查询pool
	pool, err := dao.GetStakePoolByCode(req.PoolCode)

	if err != nil {
		logs.Error("根据poolCode查询pool异常", err)
		myCtx.JSON2(API.COMMOM_ERROR, "根据poolCode查询pool异常")
		return
	} else if pool == nil {
		logs.Warn("该质押池不存在", req)
		myCtx.JSON(API.POOL_NOT_EXIST, "")
		return
	}

	userAddress := common.HexToAddress(scUser.UserWalletAddress)
	info, err := instance.GetUserInfo(nil, new(big.Int).SetInt64(pool.PoolID), userAddress)

	if err != nil {
		logs.Error("获取用户质押信息失败", err)
		myCtx.JSON2(API.COMMOM_ERROR, "获取用户质押信息失败")
		return
	}
	myCtx.JSON(API.SUCCESS, info)
}

// GetAllPool 获取所有池子信息
func GetAllPool(c *gin.Context) {
	myCtx := &CustomContext{c}

	var req request.AllPoolReq

	err := myCtx.ShouldBindJSON(&req)

	if err != nil {
		logs.Error("Error binding JSON:", err.Error())
		myCtx.JSON(API.DATA_ERROR, err.Error())
		return
	}

	if errorMsg := dataValidate(req); errorMsg != nil {
		logs.Error("入参校验失败:", errorMsg)
		myCtx.JSON(API.DATA_VALIDATE_ERROR, "")
		return
	}

	logs.Info("获取所有池子信息入参为：", req)

	allPool, err := dao.SelectStakePoolPage(req.PageSize, req.PageNum, "", "")

	if err != nil {
		logs.Error("获取所有池子信息失败", err)
		myCtx.JSON2(API.COMMOM_ERROR, "获取所有池子信息失败")
		return
	}
	logs.Info("查询质押池结果为：", allPool)
	myCtx.JSON(API.SUCCESS, allPool)
}

// GetPoolDetail 获取质押池详情
func GetPoolDetail(c *gin.Context) {

	myCtx := &CustomContext{c}

	poolCode := myCtx.Query("poolCode")

	if len(poolCode) <= 0 {
		logs.Info("获取质押池详情，入参为空")
		myCtx.JSON2(API.COMMOM_ERROR, "获取质押池详情入参不能为空")
		return
	}

	pool, err := dao.GetStakePoolByCode(poolCode)
	if err != nil {
		logs.Error("根据poolCode查询pool异常", err)
		myCtx.JSON2(API.COMMOM_ERROR, "查询质押池异常")
		return
	}
	logs.Info("根据poolCode查询pool结果为：", pool)

	contractAddress := common.HexToAddress(app.ContractAddress)

	dial, err := ethclient.Dial(app.DeployAddress)
	if err != nil {
		logs.Error("获取以太坊节点失败:", err.Error())
		myCtx.JSON(API.DATA_ERROR, err.Error())
		return
	}

	schStake, err := contract.NewSCHStake(contractAddress, dial)
	if err != nil {
		logs.Error("获取合约实例失败:", err)
		myCtx.JSON2(API.COMMOM_ERROR, "获取合约实例失败")
		return
	}
	poolDetail, _ := schStake.GetPool(nil, new(big.Int).SetInt64(pool.PoolID))

	//todo 这里要进行质押池数据的组装,暂时先不组装
	myCtx.JSON(API.SUCCESS, poolDetail)

}

// 保存交易信息
func saveTransLogs(receipt *types.Receipt, tx *types.Transaction, address common.Address, scUser *model.ScUser) bool {
	//保存交易信息
	transLogs, _ := utils.ToJson(receipt.Logs)
	scTransLogs := model.ScTransactionLog{
		TransType:   API.AddPool.String(),
		TransHash:   tx.Hash().Hex(),
		Gas:         int64(tx.Gas()),
		GasPrice:    tx.GasPrice().Int64(),
		TransStatus: int8(receipt.Status),
		TransLogs:   transLogs,
		TransCost:   tx.Cost().Int64(),
		TransFrom:   address.Hex(),
		TransTo:     tx.To().Hex(),
		Creator:     scUser.UserAddress,
	}

	err2 := dao.CreateScTransLog(&scTransLogs)
	if err2 != nil {
		logs.Error("保存交易信息失败", err2)
		return true
	}
	return false
}

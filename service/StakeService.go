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
	"strong-defi-backend/contract/schStake"
	"strong-defi-backend/model"
	"strong-defi-backend/pkg/log"
	"strong-defi-backend/request"
	"strong-defi-backend/utils"
)

// GetAccountBalance 获取账号余额
func GetAccountBalance(c *gin.Context) {
	myCtx := &CustomContext{c}

	var accountReq request.AccountBalanceReq
	err := myCtx.ShouldBindJSON(&accountReq)

	if err != nil {
		log.Error().Msgf("Error binding JSON:(%+v)", err.Error())
		myCtx.JSON(API.DATA_ERROR, err.Error())
		return
	}

	if errorMsg := dataValidate(accountReq); errorMsg != nil {
		log.Error().Msgf("入参校验失败:(%+v)", errorMsg)
		myCtx.JSON(API.DATA_VALIDATE_ERROR, "")
		return
	}

	/*获取连接*/
	client, err1 := ethclient.Dial(accountReq.Url)
	if err1 != nil {
		log.Error().Msgf("获取以太坊节点失败:(%+v)", err.Error())
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
		_ = json.Unmarshal([]byte(value.(string)), scUser)
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
		log.Error().Msgf("Error binding JSON:(%+v)", err.Error())
		myCtx.JSON(API.DATA_ERROR, err.Error())
		return
	}

	if errorMsg := dataValidate(addPoolReq); errorMsg != nil {
		log.Error().Msgf("入参校验失败:(%+v)", errorMsg)
		myCtx.JSON(API.DATA_VALIDATE_ERROR, "")
		return
	}
	//建立链接
	client, err := ethclient.Dial(app.DeployAddress)

	if err != nil {
		log.Error().Msgf("以太坊连接失败。(%+v)", err)
		myCtx.JSON2(API.COMMOM_ERROR, "以太坊连接失败")
		return
	}
	address := common.HexToAddress(scUser.UserWalletAddress)
	//获取nonce值
	nonce, err := client.PendingNonceAt(context.Background(), address)

	if err != nil {
		log.Error().Msgf("获取nonce失败。(%+v)", err)
		myCtx.JSON2(API.COMMOM_ERROR, "获取nonce失败")
		return
	}
	newContract, _ := sch_stake.NewSCHStake(common.HexToAddress(app.ContractAddress), client)

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
			if saveTransLogs(receipt, tx, address, scUser, API.AddPool.String()) {
				return
			}

			for _, vlog := range receipt.Logs {
				parseABI, _ := abi.JSON(strings.NewReader(sch_stake.SCHStakeMetaData.ABI))
				//这是固定的，获取事件信息的，第一个是合约事件返回的结构，第二个是事件名称
				var events []request.Event
				err2 := parseABI.UnpackIntoInterface(&events, "AddPool", vlog.Data)

				if err2 != nil {
					log.Error().Msgf("获取nonce失败。(%+v)", err)
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
		log.Error().Msgf("添加质押池报错。(%+v)", err)
		myCtx.JSON2(API.COMMOM_ERROR, "添加质押池报错")
		return
	}
	log.Info().Msgf("添加质押池成功。")
	myCtx.JSON(API.SUCCESS, "")
}

// UserStake 用户质押
func UserStake(c *gin.Context) {
	myCtx := &CustomContext{c}

	var req request.UserStakeReq

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

	value, _ := myCtx.Get(API.TOKEN_KEY)

	var scUser model.ScUser
	err = json.Unmarshal([]byte(value.(string)), &scUser)
	if err != nil {
		log.Error().Msgf("解析用户信息失败。(%+v)", err)
		myCtx.JSON2(API.COMMOM_ERROR, "解析用户信息失败")
		return
	}

	dial, err := ethclient.Dial(app.DeployAddress)
	if err != nil {
		log.Error().Msgf("连接合约失败。(%+v)", err)
		myCtx.JSON2(API.COMMOM_ERROR, "连接合约失败")
		return
	}

	//查询质押池
	stakePool, err := dao.GetStakePoolByCode(req.PoolCode)
	if err != nil {
		log.Error().Msgf("查询质押池失败。(%+v)", err)
		myCtx.JSON2(API.COMMOM_ERROR, "查询质押池失败")
		return
	} else if stakePool == nil {
		log.Warn().Msgf("该质押池不存在,入参为：%+v", req)
		myCtx.JSON(API.POOL_NOT_EXIST, "")
		return
	}

	contractAddress := common.HexToAddress(app.ContractAddress)

	scStake, err := sch_stake.NewSCHStake(contractAddress, dial)
	if err != nil {
		log.Error().Msgf("创建合约实例失败。(%+v)", err)
		myCtx.JSON2(API.COMMOM_ERROR, "创建合约实例失败")
		return
	}

	userAddress := common.HexToAddress(scUser.UserWalletAddress)
	opts := &bind.TransactOpts{
		From: userAddress,
	}
	//用户质押
	tx, err := scStake.Stake(opts, new(big.Int).SetInt64(stakePool.PoolID), userAddress, new(big.Int).SetInt64(req.Amount))
	if err != nil {
		log.Error().Msgf("用户质押失败。%+v", err)
		myCtx.JSON2(API.COMMOM_ERROR, "用户质押失败")
		return
	}

	receipt, _ := dial.TransactionReceipt(context.Background(), tx.Hash())

	//如果交易成功，则保存交易日志
	if receipt.Status == types.ReceiptStatusSuccessful {

		scStakeRecord := model.ScStakeRecord{
			StakeType:         API.UserStake.String(),
			PoolID:            stakePool.PoolID,
			UserWalletAddress: scUser.UserWalletAddress,
			TransHash:         tx.Hash().String(),
			TransAmount:       req.Amount,
			Creator:           scUser.UserAddress,
		}
		//保存交易记录
		_ = dao.CreateScStakeRecord(scStakeRecord)
		//存入交易日志表
		if saveTransLogs(receipt, tx, userAddress, &scUser, API.UserStake.String()) {
			log.Error().Msgf("用户质押失败。%+v", err)
			return
		}
	}
}

func GetUserInfo(c *gin.Context) {

	myCtx := &CustomContext{c}

	var req request.QueryUserInfoReq

	err := myCtx.ShouldBindJSON(&req)

	if err != nil {
		log.Error().Msgf("Error binding JSON:%+v", err.Error())
		myCtx.JSON(API.DATA_ERROR, err.Error())
		return
	}

	if errorMsg := dataValidate(req); errorMsg != nil {
		log.Error().Msgf("入参校验失败:(%+v)", errorMsg)
		myCtx.JSON(API.DATA_VALIDATE_ERROR, "")
		return
	}

	//建立连接
	dial, _ := ethclient.Dial(app.DeployAddress)
	//获取合约地址
	address := common.HexToAddress(app.ContractAddress)

	/*获取合约实例*/
	instance, err := sch_stake.NewSCHStake(address, dial)

	//获取登陆人信息
	scUser := new(model.ScUser)
	value, exists := myCtx.Get(API.TOKEN_KEY)
	if exists {
		_ = json.Unmarshal([]byte(value.(string)), scUser)
	} else {
		myCtx.JSON2(API.COMMOM_ERROR, "请登陆后在操作")
		return
	}
	log.Info().Msgf("获取用户质押信息")

	//通过poolCode查询pool
	pool, err := dao.GetStakePoolByCode(req.PoolCode)

	if err != nil {
		log.Error().Msgf("根据poolCode查询pool异常:(%+v)", err)
		myCtx.JSON2(API.COMMOM_ERROR, "根据poolCode查询pool异常")
		return
	} else if pool == nil {
		log.Warn().Msgf("该质押池不存在", req)
		myCtx.JSON(API.POOL_NOT_EXIST, "")
		return
	}

	userAddress := common.HexToAddress(scUser.UserWalletAddress)
	//交易一定要有opts，不然会报错。获取不到交易发起者
	opts := &bind.CallOpts{
		//true是在待处理状态下进行交易，这样会有失败的风险，false合约调用将会在
		//最后已确认的状态 下执行。即调用会基于最新的已完成的区块来计算合约的结果，不考虑当前待处理交易的影响
		Pending: false,
		From:    userAddress,
	}
	info, err := instance.GetUserInfo(opts, new(big.Int).SetInt64(pool.PoolID), userAddress)

	if err != nil {
		log.Error().Msgf("获取用户质押信息失败:(%+v)", err)
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
		log.Error().Msgf("Error binding JSON:(%+v)", err.Error())
		myCtx.JSON(API.DATA_ERROR, err.Error())
		return
	}

	if errorMsg := dataValidate(req); errorMsg != nil {
		log.Error().Msgf("入参校验失败:(%+v)", errorMsg)
		myCtx.JSON(API.DATA_VALIDATE_ERROR, "")
		return
	}

	log.Info().Msgf("获取所有池子信息入参为：(%+v)", req)

	allPool, err := dao.SelectStakePoolPage(req.PageSize, req.PageNum, "", "")

	if err != nil {
		log.Error().Msgf("获取所有池子信息失败:(%+v)", err)
		myCtx.JSON2(API.COMMOM_ERROR, "获取所有池子信息失败")
		return
	}
	log.Info().Msgf("查询质押池结果为：(%+v)", allPool)
	myCtx.JSON(API.SUCCESS, allPool)
}

// GetPoolDetail 获取质押池详情
func GetPoolDetail(c *gin.Context) {

	myCtx := &CustomContext{c}

	poolCode := myCtx.Query("poolCode")

	if len(poolCode) <= 0 {
		log.Info().Msgf("获取质押池详情，入参为空")
		myCtx.JSON2(API.COMMOM_ERROR, "获取质押池详情入参不能为空")
		return
	}

	pool, err := dao.GetStakePoolByCode(poolCode)
	if err != nil {
		log.Error().Msgf("根据poolCode查询pool异常", err)
		myCtx.JSON2(API.COMMOM_ERROR, "查询质押池异常")
		return
	}
	log.Info().Msgf("根据poolCode查询pool结果为：(%+v)", pool)

	contractAddress := common.HexToAddress(app.ContractAddress)

	dial, err := ethclient.Dial(app.DeployAddress)
	if err != nil {
		log.Error().Msgf("获取以太坊节点失败:(%+v)", err.Error())
		myCtx.JSON(API.DATA_ERROR, err.Error())
		return
	}

	schStake, err := sch_stake.NewSCHStake(contractAddress, dial)
	if err != nil {
		log.Error().Msgf("获取合约实例失败:(%+v)", err)
		myCtx.JSON2(API.COMMOM_ERROR, "获取合约实例失败")
		return
	}
	poolDetail, _ := schStake.GetPool(nil, new(big.Int).SetInt64(pool.PoolID))

	//todo 这里要进行质押池数据的组装,暂时先不组装
	myCtx.JSON(API.SUCCESS, poolDetail)
}

// UserPauseStake 用户暂停质押
func UserPauseStake(c *gin.Context) {

	myCtx := &CustomContext{c}

	var req request.UserPauseStakeReq

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

	value, _ := myCtx.Get(API.TOKEN_KEY)

	var scUser model.ScUser
	err = json.Unmarshal([]byte(value.(string)), &scUser)
	if err != nil {
		log.Error().Msgf("解析用户信息失败。(%+v)", err)
		myCtx.JSON2(API.COMMOM_ERROR, "解析用户信息失败")
		return
	}
	client, err := ethclient.Dial(app.DeployAddress)
	if err != nil {
		log.Error().Msgf("连接合约失败。(%+v)", err)
		myCtx.JSON2(API.COMMOM_ERROR, "连接合约失败")
		return
	}
	//用户地址
	userAddress := common.HexToAddress(scUser.UserWalletAddress)

	stake, _ := sch_stake.NewSCHStake(common.HexToAddress(app.ContractAddress), client)

	//获取chaiID
	chaiID, err := client.ChainID(context.Background())

	//nonce值
	nonce, _ := client.NonceAt(context.Background(), userAddress, nil)
	//获取交易签名函数
	ecdsa, _ := crypto.HexToECDSA(req.UserPrivateKey)
	//创建签名函数
	signer := utils.CreateSigner(ecdsa, chaiID)

	//一定要有这个opts。因为要有From的值，不然获取不到msg.sender,这里的opts相当于通过什么账户链接合约的
	opts := &bind.TransactOpts{
		GasLimit: 21000,
		From:     userAddress,
		//设置nonce值进行防重处理
		Nonce:  new(big.Int).SetInt64(int64(nonce)),
		Signer: signer,
	}

	pool, _ := dao.GetStakePoolByCode(req.PoolCode)
	pauseStakeTx, err := stake.PauseStake(opts, new(big.Int).SetInt64(pool.PoolID), userAddress, new(big.Int).SetInt64(req.Amount))

	if err != nil {
		log.Error().Msgf("暂停质押失败:(%+v)", err)
		myCtx.JSON2(API.COMMOM_ERROR, "暂停质押失败")
		return
	}
	transactionReceipt, _ := client.TransactionReceipt(context.Background(), pauseStakeTx.Hash())
	log.Info().Msgf("暂停质押交易状态失败:(%+v)", transactionReceipt.Logs)
	if transactionReceipt.Status == types.ReceiptStatusSuccessful {
		//保存交易信息
		saveTransLogs(transactionReceipt, pauseStakeTx, userAddress, &scUser, API.UserPauseStake.String())
	} else {
		log.Error().Msgf("暂停质押交易状态失败:(%+v)", transactionReceipt.Status)
		myCtx.JSON2(API.COMMOM_ERROR, "暂停质押交易状态失败")
		return
	}
	myCtx.JSON(API.SUCCESS, "")
}

// 保存交易信息
func saveTransLogs(receipt *types.Receipt,
	tx *types.Transaction,
	address common.Address,
	scUser *model.ScUser,
	transType string) bool {
	//保存交易信息
	transLogs, _ := utils.ToJson(receipt.Logs)
	scTransLogs := model.ScTransactionLog{
		TransType:   transType,
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
		log.Error().Msgf("保存交易信息失败:(%+v)", err2)
		return true
	}
	return false
}

package request

type AccountBalanceReq struct {
	/*需要节点的地址*/
	Url string `json:"url" validate:"required"`

	/*账号地址*/
	Address string `json:"address" validate:"required"`
}

type AddPoolReq struct {
	//池权重
	Wight int `json:"wight" validate:"required"`
	//token地址
	TokenAddress string `json:"tokenAddress" validate:"required"`
	//最小质押金额
	MinStakeAmount int `json:"minStakeAmount" validate:"required"`
	//是否要更新质押池
	IsUpdatePool bool `json:"isUpdatePool"`
	//解除质押的的区块数
	UnStakeLockBlockNumber int `json:"UnStakeLockBlockNumber"`

	GasLimit uint64 `json:"gasLimit"`

	//用户私钥，暂时这样写，回头看能不能从哪里获取，或者进行加密处理
	UserPrivateKey string `json:"userPrivateKey"`
}

// UserApproveReq 用户授权请求
type UserApproveReq struct {
	//用户私钥，todo 这里还需要重新设计私钥的管理，不能让用户直接将私钥暴露在钱后端交互中
	UserPrivateKey string `json:"userPrivateKey" validate:"required"`

	//授权地址
	ApproveAddress string `json:"approveAddress" validate:"required"`

	//授权金额
	Amount int64 `json:"amount" validate:"required"`
}

// UserStakeReq 用户质押请求
type UserStakeReq struct {
	PoolCode string `json:"poolCode" validate:"required"`
	Amount   int64  `json:"amount" validate:"required"`
}

// 查询用户信息
type QueryUserInfoReq struct {

	//池id
	PoolCode string `json:"poolCode" validate:"required"`
}

type AllPoolReq struct {
	PageNum  int `json:"pageNum" validate:"required"`
	PageSize int `json:"pageSize" validate:"required"`
}

// UserTransferReq 用户转账请求
type UserTransferReq struct {

	//接收人地址
	ReceiveAddress string `json:"receiveAddress" validate:"required"`
	//发起人地址
	SendAddress string `json:"sendAddress" validate:"required"`
	//转账数量
	Amount int64 `json:"amount" validate:"required"`
}

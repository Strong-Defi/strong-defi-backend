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

// 用户质押请求
type UserStakeReq struct {
}

package req

type AccountBalanceReq struct {
	/*需要节点的地址*/
	Url string `json:"url" validate:"required"`

	/*账号地址*/
	Address string `json:"address" validate:"required"`
}

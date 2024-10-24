package req

type UserLoginReq struct {

	//钱包地址
	WalletAddress string `json:"walletAddress" validate:"required"`
}

package request

type UserLoginReq struct {

	//钱包地址
	WalletAddress string `json:"walletAddress" validate:"required"`
}

type UserRegister struct {
	UserWalletAddress string `json:"userWalletAddress" validate:"required"`

	UserTelephone string `json:"userTelephone" validate:"required"`
	UserName      string `json:"userName" validate:"required"`
	UserPassword  string `json:"userPassword" validate:"required"`
	UserEmail     string `json:"userEmail" validate:"required"`
	UserCountry   string `json:"userCountry" validate:"required"`
	UserAddress   string `json:"userAddress" validate:"required"`
	Remark        string `json:"remark" validate:"required"`
}

type ModifyUserReq struct {
	// 用户id
	UserId uint64 `json:"userId" validate:"required"`
	// 钱包地址
	UserWalletAddress string `json:"userWalletAddress"`
	// 手机号
	UserTelephone string `json:"userTelephone" validate:"required"`
	// 用户名
	UserName string `json:"userName" validate:"required"`
	// 密码
	UserPassword string `json:"userPassword" validate:"required"`
	// 邮箱
	UserEmail string `json:"userEmail" validate:"required"`
	// 国家
	UserCountry string `json:"userCountry" validate:"required"`
	// 用户住址
	UserAddress string `json:"userAddress" validate:"required"`
	// 备注
	Remark string `json:"remark" validate:"required"`
	// nil:修改所有（并且有值就改，没有值就不改），1:修改手机号，2：修改邮箱，3：修改钱包地址，4：需改密码
	ModifyType *int `json:"modifyType"`
}

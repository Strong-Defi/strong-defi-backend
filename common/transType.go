package API

type TransType int

const (
	//添加质押池
	AddPool TransType = iota
	//用户质押
	UserStake
	UserApprove
	UserTransfer
	UserPauseStake
)

// String 方法用于转换为字符串
func (s TransType) String() string {
	switch s {
	case AddPool:
		return "ADD_POOL"
	case UserStake:
		return "USER_STAKE"
	case UserApprove:
		return "USER_APPROVE"
	case UserTransfer:
		return "USER_TRANSFER"
	case UserPauseStake:
		return "USER_PAUSE_STAKE"
	default:
		return "UNKNOWN"
	}
}

func (s TransType) TransTypeParse(value string) TransType {
	switch value {
	case "ADD_POOL":
		return 0
	case "USER_STAKE":
		return 1
	case "USER_APPROVE":
		return 2
	case "USER_TRANSFER":
		return 3
	case "USER_PAUSE_STAKE":
		return 4
	default:
		return -1
	}
}

// IsValid 验证枚举值是否有效
func (s TransType) IsValid() bool {
	switch s {
	case AddPool:
		return true
	case UserStake:
		return true
	case UserApprove:
		return true
	case UserTransfer:
		return true
	case UserPauseStake:
		return true
	default:
		return false
	}
}

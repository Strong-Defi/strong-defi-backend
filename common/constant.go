package API

const TOKEN_KEY = "userInfo"

type TransType int

const (
	AddPool TransType = iota
)

// String 方法用于转换为字符串
func (s TransType) String() string {
	switch s {
	case AddPool:
		return "ADD_POOL"
	default:
		return "UNKNOWN"
	}
}

func (s TransType) TransTypeParse(value string) TransType {
	switch value {
	case "ADD_POOL":
		return AddPool
	default:
		return -1
	}
}

// IsValid 验证枚举值是否有效
func (s TransType) IsValid() bool {
	switch s {
	case AddPool:
		return true
	default:
		return false
	}
}

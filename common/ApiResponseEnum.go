package API

//这里是一个枚举，用于返回给前端的code和message
//同时也是返回内容的状态码
//要设置一个返回枚举如下：
//第一步：在const 中设置自己的状态码，例如：数据异常对对应的状态码为100001，如下的配置

type ApiResponseEnum string

const (
	SUCCESS             ApiResponseEnum = "SUCCESS"
	USER_EXIST          ApiResponseEnum = "USER_EXIST"
	DATA_VALIDATE_ERROR ApiResponseEnum = "DATA_VALIDATE_ERROR"
	DATA_ERROR          ApiResponseEnum = "DATA_ERROR"
	DATA_PARSE_ERROR    ApiResponseEnum = "DATA_PARSE_ERROR"
	COMMOM_ERROR        ApiResponseEnum = "COMMOM_ERROR"
	NO_AUTH_ERROR       ApiResponseEnum = "NO_AUTH_ERROR"
)

func GetMessage(enum ApiResponseEnum) string {
	codeMap := map[ApiResponseEnum]string{
		SUCCESS:             "success",
		DATA_ERROR:          "数据异常",
		DATA_VALIDATE_ERROR: "数据校验异常",
		DATA_PARSE_ERROR:    "数据解析异常",
		COMMOM_ERROR:        "通用错误",
		NO_AUTH_ERROR:       "请登录后操作",
		USER_EXIST:          "用户已存在",
	}
	if codeMap[enum] == "" {
		return ""
	}
	return codeMap[enum]
}

func GetCode(enum ApiResponseEnum) string {
	codeMap := map[ApiResponseEnum]string{
		SUCCESS:             "200",
		DATA_ERROR:          "100001",
		DATA_VALIDATE_ERROR: "100002",
		DATA_PARSE_ERROR:    "100003",
		COMMOM_ERROR:        "100004",
		NO_AUTH_ERROR:       "100005",
		USER_EXIST:          "100006",
	}
	if codeMap[enum] == "" {
		return ""
	}
	return codeMap[enum]
}

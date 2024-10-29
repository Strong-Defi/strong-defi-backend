package string

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"time"
)

func Generate16CharString(fixedString string) string {
	// 获取当前时间戳
	timestamp := time.Now().Unix()

	// 将时间戳和固定字符串拼接
	input := fmt.Sprintf("%d%s", timestamp, fixedString)

	// 计算 MD5 哈希
	hash := md5.Sum([]byte(input))

	// 返回前 16 位的十六进制字符串
	return hex.EncodeToString(hash[:])[:16]
}

func ToJson(obj any) (str string, err error) {
	strByte, err := json.Marshal(obj)
	return string(strByte), err
}

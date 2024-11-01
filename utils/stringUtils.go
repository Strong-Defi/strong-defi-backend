package utils

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"strconv"
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

// Generate16CharInt 通过int生成32位随机字符串
func Generate16CharInt(fixedString int64) string {
	// 获取当前时间戳
	timestamp := time.Now().Unix()

	//将固定字符串转换为字符串
	fixedStringStr := strconv.FormatInt(fixedString, 10)
	// 生成 UUID
	uniqueID := uuid.New().String()

	// 将时间戳和固定字符串拼接
	input := fmt.Sprintf("%d%s%s", timestamp, fixedStringStr, uniqueID)

	// 计算 MD5 哈希
	hash := md5.Sum([]byte(input))

	// 返回前 16 位的十六进制字符串
	return hex.EncodeToString(hash[:])
}

func ToJson(obj any) (str string, err error) {
	strByte, err := json.Marshal(obj)
	return string(strByte), err
}

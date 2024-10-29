package utils

import (
	"math"
	"math/big"
)

// ConvertToMathFloat2 /*  -
//   - @param var1		目标值
//   - @param decimal	精度
//   - @param point   	小数位数
func ConvertToMathFloat2(var1 *big.Int, decimal int, point int) string {
	float1 := new(big.Float)
	value, _ := float1.SetString(var1.String())
	// 除以精度，结果保留两位小数
	result := new(big.Float).SetPrec(64).SetMode(big.ToNearestEven).Quo(value, big.NewFloat(math.Pow10(decimal)))
	return result.Text('f', point)
}

// GetDefaultValueInt 设置三元运算数据
func GetDefaultValueInt(condition bool, targetValue uint64, defaultValue uint64) uint64 {
	if condition {
		return defaultValue
	} else {
		return targetValue
	}
}

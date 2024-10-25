package utils

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"time"
)

var secretKey = []byte("your_secret_key")

// CreateToken 创建 JWT
func CreateToken(userInfo string) (string, error) {
	claims := jwt.MapClaims{
		"userInfo": userInfo,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

// ValidateToken 验证 JWT
func ValidateToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// 确保签名方法是预期的
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})
}

func getUserInfo(token string) (userInfo string, err error) {
	parse, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		// 确保签名方法是预期的
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})
	claims := parse.Claims.(jwt.MapClaims)["userInfo"].(string)

	return claims, err
}

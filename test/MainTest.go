package main

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"strong-defi-backend/utils"
)

func main() {
	tokenString, _ := utils.CreateToken("123")

	fmt.Println(tokenString)

	token, _ := utils.ValidateToken(tokenString)
	claims := token.Claims.(jwt.MapClaims)

	fmt.Println(claims["userInfo"])
}

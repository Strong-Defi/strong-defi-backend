package main

import (
	"fmt"
	"strong-defi-backend/utils"
)

func main() {
	tokenString, _ := utils.CreateToken("123")

	fmt.Println(tokenString)

	err := utils.ValidateToken(tokenString)

	fmt.Println(err)
}

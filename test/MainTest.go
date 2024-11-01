package main

import (
	"fmt"
	"strong-defi-backend/utils"
)

func main() {
	charInt := utils.Generate16CharInt(12)

	fmt.Println(charInt)
}

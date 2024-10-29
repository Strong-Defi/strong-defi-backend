package test

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
)

func main() {
	address := common.HexToAddress("0x9cae7e457b6fe0f378e66eee21659e005487a976d2f866669bb1c165017a8bc3")
	fmt.Println(address)

	s := address.String()

	fmt.Println(s)
}

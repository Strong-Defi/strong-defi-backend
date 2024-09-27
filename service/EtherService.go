package service

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
)

func GetContractInfo(c *gin.Context) {
	logs.Info("Info to Ethereum")
	logs.Error("Error to Ethereum")
	logs.Debug("Debug to Ethereum")
	logs.Warn("Warn to Ethereum")
	client, _ := ethclient.Dial("https://cloudflare-eth.com")
	if client == nil {
		logs.Error("Connected to Ethereum")
	}
}

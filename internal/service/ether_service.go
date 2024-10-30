package service

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/gin-gonic/gin"
)

var (
	glog log.Logger
)

type EtherService struct {
}

func NewEtherService() *EtherService {

	return NewEtherService()
}

func GetContractInfo(c *gin.Context) {
	logs.Info()
	logs.Info("Info to Ethereum")
	logs.Error("Error to Ethereum")
	logs.Debug("Debug to Ethereum")
	logs.Warn("Warn to Ethereum")
	client, _ := ethclient.Dial("https://cloudflare-eth.com")
	if client == nil {
		logs.Error("Connected to Ethereum")
	}
}

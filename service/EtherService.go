package service

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"strong-defi-backend/pkg/log"
)

func GetContractInfo(c *gin.Context) {
	log.Info().Msgf("Info to Ethereum")
	client, _ := ethclient.Dial("https://cloudflare-eth.com")
	if client == nil {
		log.Error().Msgf("Connected to Ethereum")
	}
}

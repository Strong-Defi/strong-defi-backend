package ethereum

import (
	"context"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	testPrivateKey = "YOUR_PRIVATE_KEY" // 替换为你的测试私钥
	testEndpoint   = "https://sepolia.infura.io/v3/YOUR-API-KEY"
)

func setupTestClient(t *testing.T) *Client {
	client, err := NewClient([]string{testEndpoint},
		WithRetryTimeout(time.Second*2),
		WithMaxRetries(3),
		WithHealthCheckInterval(time.Second*30),
	)
	require.NoError(t, err)
	return client
}

func TestCreateAndSignTransaction(t *testing.T) {
	client := setupTestClient(t)
	defer client.Close()

	// 创建新账户
	privateKey, address, err := client.CreateAccount()
	require.NoError(t, err)
	assert.NotNil(t, privateKey)
	assert.NotEqual(t, common.Address{}, address)

	// 测试签名消息
	message := []byte("Hello, Ethereum!")
	signature, err := client.SignMessage(message, privateKey)
	require.NoError(t, err)

	// 验证签名
	isValid := client.VerifySignature(message, signature, address)
	assert.True(t, isValid)
}

func TestSendAndWaitTransaction(t *testing.T) {
	client := setupTestClient(t)
	defer client.Close()

	ctx := context.Background()
	privateKey, err := client.GetPrivateKeyFromHex(testPrivateKey)
	require.NoError(t, err)

	from := client.GetAddressFromPrivateKey(privateKey)
	to := common.HexToAddress("0x742d35Cc6634C0532925a3b844Bc454e4438f44e") // 测试接收地址
	value := big.NewInt(100000)                                             // 0.0001 ETH

	// 发送交易
	hash, err := client.SendTransaction(ctx, from, to, value, privateKey)
	require.NoError(t, err)
	assert.NotEqual(t, common.Hash{}, hash)

	// 等待交易确认
	receipt, err := client.WaitForTransaction(ctx, hash, 1)
	require.NoError(t, err)
	assert.Equal(t, hash, receipt.TxHash)
}

func TestGetTransactionInfo(t *testing.T) {
	client := setupTestClient(t)
	defer client.Close()

	ctx := context.Background()
	testAddress := common.HexToAddress("0x742d35Cc6634C0532925a3b844Bc454e4438f44e")

	// 测试获取余额
	balance, err := client.GetBalance(ctx, testAddress)
	require.NoError(t, err)
	assert.NotNil(t, balance)

	// 测试获取nonce
	nonce, err := client.GetTransactionCount(ctx, testAddress)
	require.NoError(t, err)
	assert.NotNil(t, nonce)

	// 测试获取最新区块
	blockNum, err := client.GetLatestBlockNumber(ctx)
	require.NoError(t, err)
	assert.NotEqual(t, uint64(0), blockNum)

	// 测试获取区块信息
	block, err := client.GetBlockByNumber(ctx, big.NewInt(int64(blockNum)))
	require.NoError(t, err)
	assert.NotNil(t, block)
}

func TestEstimateGas(t *testing.T) {
	client := setupTestClient(t)
	defer client.Close()

	ctx := context.Background()
	from := common.HexToAddress("0x742d35Cc6634C0532925a3b844Bc454e4438f44e")
	to := common.HexToAddress("0x742d35Cc6634C0532925a3b844Bc454e4438f44e")
	value := big.NewInt(100000)

	gas, err := client.EstimateGas(ctx, from, to, value, nil)
	require.NoError(t, err)
	assert.Equal(t, uint64(21000), gas) // 基础转账应该是21000 gas
}

func TestDeployContract(t *testing.T) {
	client := setupTestClient(t)
	defer client.Close()

	ctx := context.Background()
	privateKey, err := crypto.HexToECDSA(testPrivateKey)
	require.NoError(t, err)

	// 创建部署选项
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(11155111)) // Sepolia chainID
	require.NoError(t, err)

	// 简单的合约字节码（一个返回常量的合约）
	bytecode := common.FromHex("608060405234801561001057600080fd5b5060b28061001f6000396000f3fe6080604052348015600f57600080fd5b506004361060285760003560e01c8063c6888fa114602d575b600080fd5b605660048036036020811015604157600080fd5b8101908080359060200190929190505050606c565b6040518082815260200191505060405180910390f35b600060078202905091905056fea265627a7a72315820f1db31c03f9acf639442b4a21b841c321f5df9c8613d6460b677e04243229edb64736f6c63430005100032")

	// 部署合约
	contractAddress, tx, err := client.DeployContract(ctx, auth, bytecode)
	require.NoError(t, err)
	assert.NotEqual(t, common.Address{}, contractAddress)
	assert.NotNil(t, tx)

	// 等待部署完成
	receipt, err := client.WaitForTransaction(ctx, tx.Hash(), 1)
	require.NoError(t, err)
	assert.Equal(t, contractAddress, receipt.ContractAddress)
}

func TestPrivateKeyManagement(t *testing.T) {
	client := setupTestClient(t)
	defer client.Close()

	// 测试创建账户
	privateKey, address, err := client.CreateAccount()
	require.NoError(t, err)
	assert.NotNil(t, privateKey)
	assert.NotEqual(t, common.Address{}, address)

	// 测试从私钥获取地址
	derivedAddress := client.GetAddressFromPrivateKey(privateKey)
	assert.Equal(t, address, derivedAddress)

	// 测试从十六进制字符串获取私钥
	hexKey := testPrivateKey
	recoveredKey, err := client.GetPrivateKeyFromHex(hexKey)
	require.NoError(t, err)
	assert.NotNil(t, recoveredKey)
}

package ethereum

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

// SendTransaction 发送原生代币交易
func (c *Client) SendTransaction(ctx context.Context, from, to common.Address, value *big.Int, privateKey *ecdsa.PrivateKey) (common.Hash, error) {
	client := c.getNextClient().client

	nonce, err := client.PendingNonceAt(ctx, from)
	if err != nil {
		return common.Hash{}, fmt.Errorf("获取nonce失败: %w", err)
	}

	gasPrice, err := client.SuggestGasPrice(ctx)
	if err != nil {
		return common.Hash{}, fmt.Errorf("获取gas价格失败: %w", err)
	}

	chainID, err := client.ChainID(ctx)
	if err != nil {
		return common.Hash{}, fmt.Errorf("获取chainID失败: %w", err)
	}

	tx := types.NewTransaction(nonce, to, value, 21000, gasPrice, nil)

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		return common.Hash{}, fmt.Errorf("签名交易失败: %w", err)
	}

	err = client.SendTransaction(ctx, signedTx)
	if err != nil {
		return common.Hash{}, fmt.Errorf("发送交易失败: %w", err)
	}

	return signedTx.Hash(), nil
}

// WaitForTransaction 等待交易被确认
func (c *Client) WaitForTransaction(ctx context.Context, hash common.Hash, confirmations uint64) (*types.Receipt, error) {
	client := c.getNextClient().client

	// 先获取交易对象
	tx, isPending, err := client.TransactionByHash(ctx, hash)
	if err != nil {
		return nil, fmt.Errorf("获取交易失败: %w", err)
	}

	if isPending {
		// 使用bind.WaitMined等待交易被确认
		receipt, err := bind.WaitMined(ctx, client, tx)
		if err != nil {
			return nil, fmt.Errorf("等待交易确认失败: %w", err)
		}

		if confirmations > 0 {
			latestBlock, err := client.BlockNumber(ctx)
			if err != nil {
				return nil, fmt.Errorf("获取最新区块失败: %w", err)
			}

			if receipt.BlockNumber.Uint64()+confirmations > latestBlock {
				return nil, fmt.Errorf("等待确认数不足")
			}
		}

		return receipt, nil
	}

	// 如果交易已经不是pending状态，直接获取receipt
	return client.TransactionReceipt(ctx, hash)
}

// SignMessage 签名消息
func (c *Client) SignMessage(message []byte, privateKey *ecdsa.PrivateKey) ([]byte, error) {
	// 添加以太坊签名前缀
	prefixedMessage := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(message), message)
	hash := crypto.Keccak256Hash([]byte(prefixedMessage))

	signature, err := crypto.Sign(hash.Bytes(), privateKey)
	if err != nil {
		return nil, fmt.Errorf("签名失败: %w", err)
	}

	// 调整 v 值
	signature[64] += 27

	return signature, nil
}

// VerifySignature 验证签名
func (c *Client) VerifySignature(message []byte, signature []byte, address common.Address) bool {
	// 添加以太坊签名前缀
	prefixedMessage := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(message), message)
	hash := crypto.Keccak256Hash([]byte(prefixedMessage))

	// 调整 v 值
	signature[64] -= 27

	// 从签名中恢复公钥
	pubKey, err := crypto.SigToPub(hash.Bytes(), signature)
	if err != nil {
		return false
	}

	// 从公钥获取地址
	recoveredAddress := crypto.PubkeyToAddress(*pubKey)

	return address == recoveredAddress
}

// EstimateGas 估算交易gas
func (c *Client) EstimateGas(ctx context.Context, from, to common.Address, value *big.Int, data []byte) (uint64, error) {
	client := c.getNextClient().client

	msg := ethereum.CallMsg{
		From:  from,
		To:    &to,
		Value: value,
		Data:  data,
	}

	gas, err := client.EstimateGas(ctx, msg)
	if err != nil {
		return 0, fmt.Errorf("估算gas失败: %w", err)
	}

	return gas, nil
}

// DeployContract 部署合约
func (c *Client) DeployContract(ctx context.Context, auth *bind.TransactOpts, bytecode []byte, params ...interface{}) (common.Address, *types.Transaction, error) {
	client := c.getNextClient().client

	// 创建合约部署交易
	tx := types.NewContractCreation(
		auth.Nonce.Uint64(),
		auth.Value,
		auth.GasLimit,
		auth.GasPrice,
		bytecode,
	)

	// 签名交易
	signedTx, err := auth.Signer(auth.From, tx)
	if err != nil {
		return common.Address{}, nil, fmt.Errorf("签名交易失败: %w", err)
	}

	// 发送交易
	err = client.SendTransaction(ctx, signedTx)
	if err != nil {
		return common.Address{}, nil, fmt.Errorf("发送交易失败: %w", err)
	}

	// 等待交易被确认
	receipt, err := c.WaitForTransaction(ctx, signedTx.Hash(), 1)
	if err != nil {
		return common.Address{}, nil, fmt.Errorf("等待交易确认失败: %w", err)
	}

	return receipt.ContractAddress, signedTx, nil
}

// CreateAccount 创建新账户
func (c *Client) CreateAccount() (*ecdsa.PrivateKey, common.Address, error) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		return nil, common.Address{}, fmt.Errorf("生成私钥失败: %w", err)
	}

	address := crypto.PubkeyToAddress(privateKey.PublicKey)

	return privateKey, address, nil
}

// GetPrivateKeyFromHex 从十六进制字符串获取私钥
func (c *Client) GetPrivateKeyFromHex(hexKey string) (*ecdsa.PrivateKey, error) {
	privateKey, err := crypto.HexToECDSA(hexKey)
	if err != nil {
		return nil, fmt.Errorf("解析私钥失败: %w", err)
	}
	return privateKey, nil
}

// GetAddressFromPrivateKey 从私钥获取地址
func (c *Client) GetAddressFromPrivateKey(privateKey *ecdsa.PrivateKey) common.Address {
	return crypto.PubkeyToAddress(privateKey.PublicKey)
}

// GetBalance 获取账户余额
func (c *Client) GetBalance(ctx context.Context, address common.Address) (*big.Int, error) {
	client := c.getNextClient().client
	return client.BalanceAt(ctx, address, nil)
}

// GetTransactionCount 获取账户nonce
func (c *Client) GetTransactionCount(ctx context.Context, address common.Address) (uint64, error) {
	client := c.getNextClient().client
	return client.NonceAt(ctx, address, nil)
}

// GetTransaction 获取交易详情
func (c *Client) GetTransaction(ctx context.Context, hash common.Hash) (*types.Transaction, bool, error) {
	client := c.getNextClient().client
	return client.TransactionByHash(ctx, hash)
}

// GetTransactionReceipt 获取交易收据
func (c *Client) GetTransactionReceipt(ctx context.Context, hash common.Hash) (*types.Receipt, error) {
	client := c.getNextClient().client
	return client.TransactionReceipt(ctx, hash)
}

// GetBlockByNumber 获取区块信息
func (c *Client) GetBlockByNumber(ctx context.Context, number *big.Int) (*types.Block, error) {
	client := c.getNextClient().client
	return client.BlockByNumber(ctx, number)
}

// GetLatestBlockNumber 获取最新区块号
func (c *Client) GetLatestBlockNumber(ctx context.Context) (uint64, error) {
	client := c.getNextClient().client
	return client.BlockNumber(ctx)
}

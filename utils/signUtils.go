package utils

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
)

// CreateSigner 返回一个 SignerFn 实现 交易的签名
func CreateSigner(privateKey *ecdsa.PrivateKey, chainID *big.Int) bind.SignerFn {
	return func(addr common.Address, tx *types.Transaction) (*types.Transaction, error) {
		// 使用 SignTx 对交易进行签名
		signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
		if err != nil {
			return nil, err
		}
		return signedTx, nil
	}
}

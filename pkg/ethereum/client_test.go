package ethereum

import (
	"context"
	"github.com/ethereum/go-ethereum/ethclient"
	"testing"
)

func TestClient(t *testing.T) {
	// 使用测试网络的endpoints
	endpoints := []string{
		"https://eth-goerli.g.alchemy.com/v2/YOUR-API-KEY",
		"https://goerli.infura.io/v3/YOUR-API-KEY",
	}

	client, err := NewClient(endpoints)
	if err != nil {
		t.Fatalf("创建客户端失败: %v", err)
	}
	defer client.Close()

	// 测试基本调用
	err = client.Call(context.Background(), func(c *ethclient.Client) error {
		blockNum, err := c.BlockNumber(context.Background())
		if err != nil {
			return err
		}
		t.Logf("当前区块高度: %d", blockNum)
		return nil
	})

	if err != nil {
		t.Errorf("调用失败: %v", err)
	}

	// 测试故障转移
	// 模拟第一个endpoint失败的情况
	client.clients[0].client.Close()

	err = client.Call(context.Background(), func(c *ethclient.Client) error {
		blockNum, err := c.BlockNumber(context.Background())
		if err != nil {
			return err
		}
		t.Logf("故障转移后获取区块高度: %d", blockNum)
		return nil
	})

	if err != nil {
		t.Errorf("故障转移失败: %v", err)
	}
}

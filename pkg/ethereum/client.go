package ethereum

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
)

type Client struct {
	endpoints           []string
	clients             []*endpointClient
	current             int
	mu                  sync.RWMutex
	retryTimeout        time.Duration
	maxRetries          int
	healthCheckInterval time.Duration
}

type endpointClient struct {
	url        string
	client     *ethclient.Client
	lastActive time.Time
	mu         sync.RWMutex
}

type ClientOption func(*Client)

// WithRetryTimeout 设置重试超时
func WithRetryTimeout(timeout time.Duration) ClientOption {
	return func(c *Client) {
		c.retryTimeout = timeout
	}
}

// WithMaxRetries 设置最大重试次数
func WithMaxRetries(maxRetries int) ClientOption {
	return func(c *Client) {
		c.maxRetries = maxRetries
	}
}

// WithHealthCheckInterval 设置健康检查间隔
func WithHealthCheckInterval(interval time.Duration) ClientOption {
	return func(c *Client) {
		c.healthCheckInterval = interval
	}
}

// NewClient 创建新的客户端
func NewClient(endpoints []string, opts ...ClientOption) (*Client, error) {
	if len(endpoints) == 0 {
		return nil, fmt.Errorf("至少需要一个endpoint")
	}

	c := &Client{
		endpoints: endpoints,
		clients:   make([]*endpointClient, len(endpoints)),
	}

	// 初始化所有endpoint客户端
	for i, url := range endpoints {
		ec := &endpointClient{url: url}
		if err := ec.connect(); err != nil {
			return nil, fmt.Errorf("连接endpoint %s 失败: %w", url, err)
		}
		c.clients[i] = ec
	}

	// 启动健康检查
	go c.healthCheck()

	return c, nil
}

func (ec *endpointClient) connect() error {
	ec.mu.Lock()
	defer ec.mu.Unlock()

	client, err := ethclient.Dial(ec.url)
	if err != nil {
		return err
	}

	ec.client = client
	ec.lastActive = time.Now()
	return nil
}

// Call 执行RPC调用并自动处理重试
func (c *Client) Call(ctx context.Context, f func(*ethclient.Client) error) error {
	for retry := 0; retry < c.maxRetries; retry++ {
		client := c.getNextClient()
		if client == nil {
			return fmt.Errorf("没有可用的endpoint")
		}

		if err := f(client.client); err != nil {
			// 如果调用失败，标记该endpoint需要重连
			go client.connect()
			continue
		}

		client.lastActive = time.Now()
		return nil
	}

	return fmt.Errorf("达到最大重试次数")
}

func (c *Client) getNextClient() *endpointClient {
	c.mu.Lock()
	defer c.mu.Unlock()

	if len(c.clients) == 0 {
		return nil
	}

	client := c.clients[c.current]
	c.current = (c.current + 1) % len(c.clients)
	return client
}

func (c *Client) healthCheck() {
	ticker := time.NewTicker(c.healthCheckInterval)
	defer ticker.Stop()

	for range ticker.C {
		for _, client := range c.clients {
			go func(ec *endpointClient) {
				if _, err := ec.client.BlockNumber(context.Background()); err != nil {
					// 如果健康检查失败，触发重连
					ec.connect()
				}
			}(client)
		}
	}
}

// Close 关闭所有客户端连接
func (c *Client) Close() {
	c.mu.Lock()
	defer c.mu.Unlock()

	for _, client := range c.clients {
		if client.client != nil {
			client.client.Close()
		}
	}
}

package ethereum

import (
	"fmt"
	"time"
)

type ChainConfig struct {
	ChainID     int64    `yaml:"chain_id"`
	Name        string   `yaml:"name"`
	Endpoints   []string `yaml:"endpoints"`
	APIKeys     []string `yaml:"api_keys,omitempty"`
	RetryConfig `yaml:"retry"`
}

type RetryConfig struct {
	MaxRetries  int           `yaml:"max_retries"`
	RetryDelay  time.Duration `yaml:"retry_delay"`
	HealthCheck time.Duration `yaml:"health_check_interval"`
}

type Config struct {
	Chains map[string]*ChainConfig `yaml:"chains"`
}

// NewClientFromConfig 根据配置创建特定链的客户端
func NewClientFromConfig(chainName string, cfg *Config) (*Client, error) {
	chainCfg, ok := cfg.Chains[chainName]
	if !ok {
		return nil, fmt.Errorf("chain %s not found in config", chainName)
	}

	// 处理endpoint和apikey的组合
	endpoints := make([]string, len(chainCfg.Endpoints))
	for i, endpoint := range chainCfg.Endpoints {
		if i < len(chainCfg.APIKeys) {
			endpoints[i] = fmt.Sprintf("%s/%s", endpoint, chainCfg.APIKeys[i])
		} else {
			endpoints[i] = endpoint
		}
	}

	opts := []ClientOption{
		WithRetryTimeout(chainCfg.RetryDelay),
		WithMaxRetries(chainCfg.MaxRetries),
		WithHealthCheckInterval(chainCfg.HealthCheck),
	}

	return NewClient(endpoints, opts...)
}

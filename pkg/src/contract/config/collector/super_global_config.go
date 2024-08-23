package collector

import (
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/config"
)

type SuperGlobalConfig struct {
	*config.CollectorConfig
	collectorType config.CollectorType
	config        []string
}

func NewSuperGlobalConfig(config []string) *SuperGlobalConfig {
	return &SuperGlobalConfig{
		CollectorConfig: &config.CollectorConfig{},
		collectorType:   config.TypeSuperGlobal,
		config:          config,
	}
}

func CreateSuperGlobalConfig(config ...string) *SuperGlobalConfig {
	return NewSuperGlobalConfig(config)
}

func (c *SuperGlobalConfig) ToArray() map[string]interface{} {
	parent := c.CollectorConfig.ToArray()

	parent["type"] = string(c.collectorType)
	parent["value"] = c.config

	return parent
}

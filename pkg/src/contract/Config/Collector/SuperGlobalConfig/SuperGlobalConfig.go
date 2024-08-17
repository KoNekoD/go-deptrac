package SuperGlobalConfig

import (
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Config/CollectorConfig"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Config/CollectorType"
)

type SuperGlobalConfig struct {
	*CollectorConfig.CollectorConfig
	collectorType CollectorType.CollectorType
	config        []string
}

func NewSuperGlobalConfig(config []string) *SuperGlobalConfig {
	return &SuperGlobalConfig{
		CollectorConfig: &CollectorConfig.CollectorConfig{},
		collectorType:   CollectorType.TypeSuperGlobal,
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

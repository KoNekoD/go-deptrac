package configs

import (
	"github.com/KoNekoD/go-deptrac/pkg/collectors"
	"github.com/KoNekoD/go-deptrac/pkg/domain/enums"
)

type SuperGlobalConfig struct {
	*collectors.CollectorConfig
	collectorType enums.CollectorType
	config        []string
}

func NewSuperGlobalConfig(config []string) *SuperGlobalConfig {
	return &SuperGlobalConfig{
		CollectorConfig: &collectors.CollectorConfig{},
		collectorType:   enums.CollectorTypeTypeSuperGlobal,
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

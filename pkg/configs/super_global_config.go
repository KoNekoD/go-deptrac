package configs

import "github.com/KoNekoD/go-deptrac/pkg/collectors"

type SuperGlobalConfig struct {
	*collectors.CollectorConfig
	collectorType collectors.CollectorType
	config        []string
}

func NewSuperGlobalConfig(config []string) *SuperGlobalConfig {
	return &SuperGlobalConfig{
		CollectorConfig: &collectors.CollectorConfig{},
		collectorType:   collectors.CollectorTypeTypeSuperGlobal,
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

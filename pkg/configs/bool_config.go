package configs

import "github.com/KoNekoD/go-deptrac/pkg/collectors"

type BoolConfig struct {
	*collectors.CollectorConfig
	collectorType collectors.CollectorType
	mustNot       []collectors.CollectorConfig
	must          []collectors.CollectorConfig
}

func NewBoolConfig() *BoolConfig {
	return &BoolConfig{
		CollectorConfig: &collectors.CollectorConfig{},
		collectorType:   collectors.CollectorTypeTypeBool,
		mustNot:         make([]collectors.CollectorConfig, 0),
		must:            make([]collectors.CollectorConfig, 0),
	}
}

func (c *BoolConfig) ToArray() map[string]interface{} {
	parent := c.CollectorConfig.ToArray()

	parent["type"] = string(c.collectorType)
	parent["mustNot"] = c.mustNot
	parent["must"] = c.must

	return parent
}

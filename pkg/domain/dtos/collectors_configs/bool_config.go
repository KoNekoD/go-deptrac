package collectors_configs

import (
	"github.com/KoNekoD/go-deptrac/pkg/domain/enums"
)

type BoolConfig struct {
	*CollectorConfig
	collectorType enums.CollectorType
	mustNot       []CollectorConfig
	must          []CollectorConfig
}

func NewBoolConfig() *BoolConfig {
	return &BoolConfig{
		CollectorConfig: &CollectorConfig{},
		collectorType:   enums.CollectorTypeTypeBool,
		mustNot:         make([]CollectorConfig, 0),
		must:            make([]CollectorConfig, 0),
	}
}

func (c *BoolConfig) ToArray() map[string]interface{} {
	parent := c.CollectorConfig.ToArray()

	parent["type"] = string(c.collectorType)
	parent["mustNot"] = c.mustNot
	parent["must"] = c.must

	return parent
}

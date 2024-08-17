package BoolConfig

import (
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Config/CollectorConfig"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Config/CollectorType"
)

type BoolConfig struct {
	*CollectorConfig.CollectorConfig
	collectorType CollectorType.CollectorType
	mustNot       []CollectorConfig.CollectorConfig
	must          []CollectorConfig.CollectorConfig
}

func NewBoolConfig() *BoolConfig {
	return &BoolConfig{
		CollectorConfig: &CollectorConfig.CollectorConfig{},
		collectorType:   CollectorType.TypeBool,
		mustNot:         make([]CollectorConfig.CollectorConfig, 0),
		must:            make([]CollectorConfig.CollectorConfig, 0),
	}
}

func (c *BoolConfig) ToArray() map[string]interface{} {
	parent := c.CollectorConfig.ToArray()

	parent["type"] = string(c.collectorType)
	parent["mustNot"] = c.mustNot
	parent["must"] = c.must

	return parent
}

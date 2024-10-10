package configs

import (
	"github.com/KoNekoD/go-deptrac/pkg/collectors"
	"github.com/KoNekoD/go-deptrac/pkg/domain/enums"
)

type BoolConfig struct {
	*collectors.CollectorConfig
	collectorType enums.CollectorType
	mustNot       []collectors.CollectorConfig
	must          []collectors.CollectorConfig
}

func NewBoolConfig() *BoolConfig {
	return &BoolConfig{
		CollectorConfig: &collectors.CollectorConfig{},
		collectorType:   enums.CollectorTypeTypeBool,
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

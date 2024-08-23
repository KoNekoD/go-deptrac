package collector

import (
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/config"
)

type BoolConfig struct {
	*config.CollectorConfig
	collectorType config.CollectorType
	mustNot       []config.CollectorConfig
	must          []config.CollectorConfig
}

func NewBoolConfig() *BoolConfig {
	return &BoolConfig{
		CollectorConfig: &config.CollectorConfig{},
		collectorType:   config.TypeBool,
		mustNot:         make([]config.CollectorConfig, 0),
		must:            make([]config.CollectorConfig, 0),
	}
}

func (c *BoolConfig) ToArray() map[string]interface{} {
	parent := c.CollectorConfig.ToArray()

	parent["type"] = string(c.collectorType)
	parent["mustNot"] = c.mustNot
	parent["must"] = c.must

	return parent
}

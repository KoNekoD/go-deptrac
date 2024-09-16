package collector

import (
	config_contract2 "github.com/KoNekoD/go-deptrac/pkg/config_contract"
)

type BoolConfig struct {
	*config_contract2.CollectorConfig
	collectorType config_contract2.CollectorType
	mustNot       []config_contract2.CollectorConfig
	must          []config_contract2.CollectorConfig
}

func NewBoolConfig() *BoolConfig {
	return &BoolConfig{
		CollectorConfig: &config_contract2.CollectorConfig{},
		collectorType:   config_contract2.TypeBool,
		mustNot:         make([]config_contract2.CollectorConfig, 0),
		must:            make([]config_contract2.CollectorConfig, 0),
	}
}

func (c *BoolConfig) ToArray() map[string]interface{} {
	parent := c.CollectorConfig.ToArray()

	parent["type"] = string(c.collectorType)
	parent["mustNot"] = c.mustNot
	parent["must"] = c.must

	return parent
}

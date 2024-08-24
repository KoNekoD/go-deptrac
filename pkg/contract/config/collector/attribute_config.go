package collector

import (
	"github.com/KoNekoD/go-deptrac/pkg/contract/config"
)

type AttributeConfig struct {
	*config.ConfigurableCollectorConfig
	collectorType config.CollectorType
}

func NewAttributeConfig(config string) *AttributeConfig {
	return &AttributeConfig{
		ConfigurableCollectorConfig: config.CreateConfigurableCollectorConfig(config),
		collectorType:               config.TypeAttribute,
	}
}

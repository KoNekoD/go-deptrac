package collector

import (
	"github.com/KoNekoD/go-deptrac/pkg/contract/config"
)

type InheritsConfig struct {
	*config.ConfigurableCollectorConfig
	collectorType config.CollectorType
}

func NewInheritsConfig(config string) *InheritsConfig {
	return &InheritsConfig{
		ConfigurableCollectorConfig: config.CreateConfigurableCollectorConfig(config),
		collectorType:               config.TypeInherits,
	}
}

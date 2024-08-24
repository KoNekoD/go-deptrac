package collector

import (
	"github.com/KoNekoD/go-deptrac/pkg/contract/config"
)

type ClassConfig struct {
	*config.ConfigurableCollectorConfig
	collectorType config.CollectorType
}

func NewClassConfig(config string) *ClassConfig {
	return &ClassConfig{
		ConfigurableCollectorConfig: config.CreateConfigurableCollectorConfig(config),
		collectorType:               config.TypeClass,
	}
}

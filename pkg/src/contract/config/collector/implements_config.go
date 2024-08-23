package collector

import (
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/config"
)

type ImplementsConfig struct {
	*config.ConfigurableCollectorConfig
	collectorType config.CollectorType
}

func NewImplementsConfig(config string) *ImplementsConfig {
	return &ImplementsConfig{
		ConfigurableCollectorConfig: config.CreateConfigurableCollectorConfig(config),
		collectorType:               config.TypeImplements,
	}
}

package collector

import (
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/config"
)

type UsesConfig struct {
	*config.ConfigurableCollectorConfig
	collectorType config.CollectorType
}

func NewUsesConfig(config string) *UsesConfig {
	return &UsesConfig{
		ConfigurableCollectorConfig: config.CreateConfigurableCollectorConfig(config),
		collectorType:               config.TypeUses,
	}
}

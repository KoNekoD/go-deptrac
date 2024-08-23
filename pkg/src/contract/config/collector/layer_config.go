package collector

import (
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/config"
)

type LayerConfig struct {
	*config.ConfigurableCollectorConfig
	collectorType config.CollectorType
}

func NewLayerConfig(config string) *LayerConfig {
	return &LayerConfig{
		ConfigurableCollectorConfig: config.CreateConfigurableCollectorConfig(config),
		collectorType:               config.TypeLayer,
	}
}

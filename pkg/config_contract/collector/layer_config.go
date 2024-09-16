package collector

import (
	config_contract2 "github.com/KoNekoD/go-deptrac/pkg/config_contract"
)

type LayerConfig struct {
	*config_contract2.ConfigurableCollectorConfig
	collectorType config_contract2.CollectorType
}

func NewLayerConfig(config string) *LayerConfig {
	return &LayerConfig{
		ConfigurableCollectorConfig: config.CreateConfigurableCollectorConfig(config),
		collectorType:               config.TypeLayer,
	}
}

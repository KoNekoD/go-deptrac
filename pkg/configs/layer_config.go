package configs

import "github.com/KoNekoD/go-deptrac/pkg/collectors"

type LayerConfig struct {
	*collectors.ConfigurableCollectorConfig
	collectorType collectors.CollectorType
}

func NewLayerConfig(config string) *LayerConfig {
	return &LayerConfig{
		ConfigurableCollectorConfig: collectors.CreateConfigurableCollectorConfig(config),
		collectorType:               collectors.CollectorTypeTypeLayer,
	}
}

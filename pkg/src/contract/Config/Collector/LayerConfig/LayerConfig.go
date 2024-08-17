package LayerConfig

import (
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Config/CollectorType"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Config/ConfigurableCollectorConfig"
)

type LayerConfig struct {
	*ConfigurableCollectorConfig.ConfigurableCollectorConfig
	collectorType CollectorType.CollectorType
}

func NewLayerConfig(config string) *LayerConfig {
	return &LayerConfig{
		ConfigurableCollectorConfig: ConfigurableCollectorConfig.CreateConfigurableCollectorConfig(config),
		collectorType:               CollectorType.TypeLayer,
	}
}

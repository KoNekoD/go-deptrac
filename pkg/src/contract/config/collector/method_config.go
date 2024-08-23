package collector

import (
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/config"
)

type MethodConfig struct {
	*config.ConfigurableCollectorConfig
	collectorType config.CollectorType
}

func NewMethodConfig(config string) *MethodConfig {
	return &MethodConfig{
		ConfigurableCollectorConfig: config.CreateConfigurableCollectorConfig(config),
		collectorType:               config.TypeMethod,
	}
}

package collector

import (
	"github.com/KoNekoD/go-deptrac/pkg/contract/config"
)

type InterfaceConfig struct {
	*config.ConfigurableCollectorConfig
	collectorType config.CollectorType
}

func NewInterfaceConfig(config string) *InterfaceConfig {
	return &InterfaceConfig{
		ConfigurableCollectorConfig: config.CreateConfigurableCollectorConfig(config),
		collectorType:               config.TypeInterface,
	}
}

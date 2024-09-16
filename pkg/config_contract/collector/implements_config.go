package collector

import (
	config_contract2 "github.com/KoNekoD/go-deptrac/pkg/config_contract"
)

type ImplementsConfig struct {
	*config_contract2.ConfigurableCollectorConfig
	collectorType config_contract2.CollectorType
}

func NewImplementsConfig(config string) *ImplementsConfig {
	return &ImplementsConfig{
		ConfigurableCollectorConfig: config.CreateConfigurableCollectorConfig(config),
		collectorType:               config.TypeImplements,
	}
}

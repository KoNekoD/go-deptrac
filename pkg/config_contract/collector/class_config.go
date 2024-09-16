package collector

import (
	config_contract2 "github.com/KoNekoD/go-deptrac/pkg/config_contract"
)

type ClassConfig struct {
	*config_contract2.ConfigurableCollectorConfig
	collectorType config_contract2.CollectorType
}

func NewClassConfig(config string) *ClassConfig {
	return &ClassConfig{
		ConfigurableCollectorConfig: config.CreateConfigurableCollectorConfig(config),
		collectorType:               config.TypeClass,
	}
}

package collector

import (
	config_contract2 "github.com/KoNekoD/go-deptrac/pkg/config_contract"
)

type ClassNameRegexConfig struct {
	*config_contract2.ConfigurableCollectorConfig
	collectorType config_contract2.CollectorType
}

func NewClassNameRegexConfig(config string) *ClassNameRegexConfig {
	return &ClassNameRegexConfig{
		ConfigurableCollectorConfig: config.CreateConfigurableCollectorConfig(config),
		collectorType:               config.TypeClassNameRegex,
	}
}

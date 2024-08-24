package collector

import (
	"github.com/KoNekoD/go-deptrac/pkg/contract/config"
)

type ClassNameRegexConfig struct {
	*config.ConfigurableCollectorConfig
	collectorType config.CollectorType
}

func NewClassNameRegexConfig(config string) *ClassNameRegexConfig {
	return &ClassNameRegexConfig{
		ConfigurableCollectorConfig: config.CreateConfigurableCollectorConfig(config),
		collectorType:               config.TypeClassNameRegex,
	}
}

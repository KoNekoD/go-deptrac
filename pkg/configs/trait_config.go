package configs

import "github.com/KoNekoD/go-deptrac/pkg/collectors"

type TraitConfig struct {
	*collectors.ConfigurableCollectorConfig
	collectorType collectors.CollectorType
}

func NewTraitConfig(config string) *TraitConfig {
	return &TraitConfig{
		ConfigurableCollectorConfig: collectors.CreateConfigurableCollectorConfig(config),
		collectorType:               collectors.CollectorTypeTypeTrait,
	}
}

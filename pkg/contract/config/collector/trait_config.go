package collector

import (
	"github.com/KoNekoD/go-deptrac/pkg/contract/config"
)

type TraitConfig struct {
	*config.ConfigurableCollectorConfig
	collectorType config.CollectorType
}

func NewTraitConfig(config string) *TraitConfig {
	return &TraitConfig{
		ConfigurableCollectorConfig: config.CreateConfigurableCollectorConfig(config),
		collectorType:               config.TypeTrait,
	}
}

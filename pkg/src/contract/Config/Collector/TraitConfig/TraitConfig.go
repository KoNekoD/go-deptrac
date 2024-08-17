package TraitConfig

import (
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Config/CollectorType"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Config/ConfigurableCollectorConfig"
)

type TraitConfig struct {
	*ConfigurableCollectorConfig.ConfigurableCollectorConfig
	collectorType CollectorType.CollectorType
}

func NewTraitConfig(config string) *TraitConfig {
	return &TraitConfig{
		ConfigurableCollectorConfig: ConfigurableCollectorConfig.CreateConfigurableCollectorConfig(config),
		collectorType:               CollectorType.TypeTrait,
	}
}

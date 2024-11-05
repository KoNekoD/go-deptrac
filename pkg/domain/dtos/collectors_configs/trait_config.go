package collectors_configs

import (
	"github.com/KoNekoD/go-deptrac/pkg/domain/enums"
)

type TraitConfig struct {
	*ConfigurableCollectorConfig
	collectorType enums.CollectorType
}

func NewTraitConfig(config string) *TraitConfig {
	return &TraitConfig{
		ConfigurableCollectorConfig: CreateConfigurableCollectorConfig(config),
		collectorType:               enums.CollectorTypeTypeTrait,
	}
}

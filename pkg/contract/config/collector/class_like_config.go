package collector

import (
	"github.com/KoNekoD/go-deptrac/pkg/contract/config"
)

type ClassLikeConfig struct {
	*config.ConfigurableCollectorConfig
	collectorType config.CollectorType
}

func NewClassLikeConfig(config string) *ClassLikeConfig {
	return &ClassLikeConfig{
		ConfigurableCollectorConfig: config.CreateConfigurableCollectorConfig(config),
		collectorType:               config.TypeClasslike,
	}
}

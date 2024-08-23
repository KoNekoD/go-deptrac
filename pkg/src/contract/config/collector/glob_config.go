package collector

import (
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/config"
)

type GlobConfig struct {
	*config.ConfigurableCollectorConfig
	collectorType config.CollectorType
}

func NewGlobConfig(config string) *GlobConfig {
	return &GlobConfig{
		ConfigurableCollectorConfig: config.CreateConfigurableCollectorConfig(config),
		collectorType:               config.TypeGlob,
	}
}

package collector

import (
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/config"
)

type ExtendsConfig struct {
	*config.ConfigurableCollectorConfig
	collectorType config.CollectorType
}

func NewExtendsConfig(config string) *ExtendsConfig {
	return &ExtendsConfig{
		ConfigurableCollectorConfig: config.CreateConfigurableCollectorConfig(config),
		collectorType:               config.TypeExtends,
	}
}

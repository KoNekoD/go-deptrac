package collector

import (
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/config"
)

type PhpInteralConfig struct {
	*config.ConfigurableCollectorConfig
	collectorType config.CollectorType
}

func NewPhpInteralConfig(config string) *PhpInteralConfig {
	return &PhpInteralConfig{
		ConfigurableCollectorConfig: config.CreateConfigurableCollectorConfig(config),
		collectorType:               config.TypePhpInternal,
	}
}

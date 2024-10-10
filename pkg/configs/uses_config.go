package configs

import "github.com/KoNekoD/go-deptrac/pkg/collectors"

type UsesConfig struct {
	*collectors.ConfigurableCollectorConfig
	collectorType collectors.CollectorType
}

func NewUsesConfig(config string) *UsesConfig {
	return &UsesConfig{
		ConfigurableCollectorConfig: collectors.CreateConfigurableCollectorConfig(config),
		collectorType:               collectors.CollectorTypeTypeUses,
	}
}

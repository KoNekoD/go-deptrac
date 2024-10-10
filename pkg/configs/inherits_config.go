package configs

import "github.com/KoNekoD/go-deptrac/pkg/collectors"

type InheritsConfig struct {
	*collectors.ConfigurableCollectorConfig
	collectorType collectors.CollectorType
}

func NewInheritsConfig(config string) *InheritsConfig {
	return &InheritsConfig{
		ConfigurableCollectorConfig: collectors.CreateConfigurableCollectorConfig(config),
		collectorType:               collectors.CollectorTypeTypeInherits,
	}
}

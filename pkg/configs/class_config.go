package configs

import "github.com/KoNekoD/go-deptrac/pkg/collectors"

type ClassConfig struct {
	*collectors.ConfigurableCollectorConfig
	collectorType collectors.CollectorType
}

func NewClassConfig(config string) *ClassConfig {
	return &ClassConfig{
		ConfigurableCollectorConfig: collectors.CreateConfigurableCollectorConfig(config),
		collectorType:               collectors.CollectorTypeTypeClass,
	}
}

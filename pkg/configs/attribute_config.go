package configs

import "github.com/KoNekoD/go-deptrac/pkg/collectors"

type AttributeConfig struct {
	*collectors.ConfigurableCollectorConfig
	collectorType collectors.CollectorType
}

func NewAttributeConfig(config string) *AttributeConfig {
	return &AttributeConfig{
		ConfigurableCollectorConfig: collectors.CreateConfigurableCollectorConfig(config),
		collectorType:               collectors.CollectorTypeTypeAttribute,
	}
}

package configs

import "github.com/KoNekoD/go-deptrac/pkg/collectors"

type MethodConfig struct {
	*collectors.ConfigurableCollectorConfig
	collectorType collectors.CollectorType
}

func NewMethodConfig(config string) *MethodConfig {
	return &MethodConfig{
		ConfigurableCollectorConfig: collectors.CreateConfigurableCollectorConfig(config),
		collectorType:               collectors.CollectorTypeTypeMethod,
	}
}

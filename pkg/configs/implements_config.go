package configs

import "github.com/KoNekoD/go-deptrac/pkg/collectors"

type ImplementsConfig struct {
	*collectors.ConfigurableCollectorConfig
	collectorType collectors.CollectorType
}

func NewImplementsConfig(config string) *ImplementsConfig {
	return &ImplementsConfig{
		ConfigurableCollectorConfig: collectors.CreateConfigurableCollectorConfig(config),
		collectorType:               collectors.CollectorTypeTypeImplements,
	}
}

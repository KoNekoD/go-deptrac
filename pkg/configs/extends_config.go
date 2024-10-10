package configs

import "github.com/KoNekoD/go-deptrac/pkg/collectors"

type ExtendsConfig struct {
	*collectors.ConfigurableCollectorConfig
	collectorType collectors.CollectorType
}

func NewExtendsConfig(config string) *ExtendsConfig {
	return &ExtendsConfig{
		ConfigurableCollectorConfig: collectors.CreateConfigurableCollectorConfig(config),
		collectorType:               collectors.CollectorTypeTypeExtends,
	}
}

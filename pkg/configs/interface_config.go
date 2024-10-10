package configs

import "github.com/KoNekoD/go-deptrac/pkg/collectors"

type InterfaceConfig struct {
	*collectors.ConfigurableCollectorConfig
	collectorType collectors.CollectorType
}

func NewInterfaceConfig(config string) *InterfaceConfig {
	return &InterfaceConfig{
		ConfigurableCollectorConfig: collectors.CreateConfigurableCollectorConfig(config),
		collectorType:               collectors.CollectorTypeTypeInterface,
	}
}

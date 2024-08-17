package InterfaceConfig

import (
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Config/CollectorType"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Config/ConfigurableCollectorConfig"
)

type InterfaceConfig struct {
	*ConfigurableCollectorConfig.ConfigurableCollectorConfig
	collectorType CollectorType.CollectorType
}

func NewInterfaceConfig(config string) *InterfaceConfig {
	return &InterfaceConfig{
		ConfigurableCollectorConfig: ConfigurableCollectorConfig.CreateConfigurableCollectorConfig(config),
		collectorType:               CollectorType.TypeInterface,
	}
}

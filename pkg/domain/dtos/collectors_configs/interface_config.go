package collectors_configs

import (
	"github.com/KoNekoD/go-deptrac/pkg/domain/enums"
)

type InterfaceConfig struct {
	*ConfigurableCollectorConfig
	collectorType enums.CollectorType
}

func NewInterfaceConfig(config string) *InterfaceConfig {
	return &InterfaceConfig{
		ConfigurableCollectorConfig: CreateConfigurableCollectorConfig(config),
		collectorType:               enums.CollectorTypeTypeInterface,
	}
}

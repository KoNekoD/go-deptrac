package configs

import (
	"github.com/KoNekoD/go-deptrac/pkg/collectors"
	"github.com/KoNekoD/go-deptrac/pkg/domain/enums"
)

type InterfaceConfig struct {
	*collectors.ConfigurableCollectorConfig
	collectorType enums.CollectorType
}

func NewInterfaceConfig(config string) *InterfaceConfig {
	return &InterfaceConfig{
		ConfigurableCollectorConfig: collectors.CreateConfigurableCollectorConfig(config),
		collectorType:               enums.CollectorTypeTypeInterface,
	}
}

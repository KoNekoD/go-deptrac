package collectors_configs

import (
	"github.com/KoNekoD/go-deptrac/pkg/domain/enums"
)

type MethodConfig struct {
	*ConfigurableCollectorConfig
	collectorType enums.CollectorType
}

func NewMethodConfig(config string) *MethodConfig {
	return &MethodConfig{
		ConfigurableCollectorConfig: CreateConfigurableCollectorConfig(config),
		collectorType:               enums.CollectorTypeTypeMethod,
	}
}

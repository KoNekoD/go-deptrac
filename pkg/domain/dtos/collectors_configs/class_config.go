package collectors_configs

import (
	"github.com/KoNekoD/go-deptrac/pkg/domain/enums"
)

type ClassConfig struct {
	*ConfigurableCollectorConfig
	collectorType enums.CollectorType
}

func NewClassConfig(config string) *ClassConfig {
	return &ClassConfig{
		ConfigurableCollectorConfig: CreateConfigurableCollectorConfig(config),
		collectorType:               enums.CollectorTypeTypeClass,
	}
}

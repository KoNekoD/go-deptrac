package collectors_configs

import (
	"github.com/KoNekoD/go-deptrac/pkg/domain/enums"
)

type InheritsConfig struct {
	*ConfigurableCollectorConfig
	collectorType enums.CollectorType
}

func NewInheritsConfig(config string) *InheritsConfig {
	return &InheritsConfig{
		ConfigurableCollectorConfig: CreateConfigurableCollectorConfig(config),
		collectorType:               enums.CollectorTypeTypeInherits,
	}
}

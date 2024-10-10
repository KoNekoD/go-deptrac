package collectors_configs

import (
	"github.com/KoNekoD/go-deptrac/pkg/domain/enums"
)

type AttributeConfig struct {
	*ConfigurableCollectorConfig
	collectorType enums.CollectorType
}

func NewAttributeConfig(config string) *AttributeConfig {
	return &AttributeConfig{
		ConfigurableCollectorConfig: CreateConfigurableCollectorConfig(config),
		collectorType:               enums.CollectorTypeTypeAttribute,
	}
}

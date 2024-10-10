package collectors_configs

import (
	"github.com/KoNekoD/go-deptrac/pkg/domain/enums"
)

type UsesConfig struct {
	*ConfigurableCollectorConfig
	collectorType enums.CollectorType
}

func NewUsesConfig(config string) *UsesConfig {
	return &UsesConfig{
		ConfigurableCollectorConfig: CreateConfigurableCollectorConfig(config),
		collectorType:               enums.CollectorTypeTypeUses,
	}
}

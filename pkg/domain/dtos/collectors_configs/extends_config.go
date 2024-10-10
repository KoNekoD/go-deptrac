package collectors_configs

import (
	"github.com/KoNekoD/go-deptrac/pkg/domain/enums"
)

type ExtendsConfig struct {
	*ConfigurableCollectorConfig
	collectorType enums.CollectorType
}

func NewExtendsConfig(config string) *ExtendsConfig {
	return &ExtendsConfig{
		ConfigurableCollectorConfig: CreateConfigurableCollectorConfig(config),
		collectorType:               enums.CollectorTypeTypeExtends,
	}
}

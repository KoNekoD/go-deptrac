package collectors_configs

import (
	"github.com/KoNekoD/go-deptrac/pkg/domain/enums"
)

type ClassLikeConfig struct {
	*ConfigurableCollectorConfig
	collectorType enums.CollectorType
}

func NewClassLikeConfig(config string) *ClassLikeConfig {
	return &ClassLikeConfig{
		ConfigurableCollectorConfig: CreateConfigurableCollectorConfig(config),
		collectorType:               enums.CollectorTypeTypeClasslike,
	}
}

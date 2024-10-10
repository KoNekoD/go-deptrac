package configs

import (
	"github.com/KoNekoD/go-deptrac/pkg/collectors"
	"github.com/KoNekoD/go-deptrac/pkg/domain/enums"
)

type ClassLikeConfig struct {
	*collectors.ConfigurableCollectorConfig
	collectorType enums.CollectorType
}

func NewClassLikeConfig(config string) *ClassLikeConfig {
	return &ClassLikeConfig{
		ConfigurableCollectorConfig: collectors.CreateConfigurableCollectorConfig(config),
		collectorType:               enums.CollectorTypeTypeClasslike,
	}
}

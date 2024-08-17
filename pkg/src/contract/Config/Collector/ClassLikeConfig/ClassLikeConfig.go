package ClassLikeConfig

import (
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Config/CollectorType"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Config/ConfigurableCollectorConfig"
)

type ClassLikeConfig struct {
	*ConfigurableCollectorConfig.ConfigurableCollectorConfig
	collectorType CollectorType.CollectorType
}

func NewClassLikeConfig(config string) *ClassLikeConfig {
	return &ClassLikeConfig{
		ConfigurableCollectorConfig: ConfigurableCollectorConfig.CreateConfigurableCollectorConfig(config),
		collectorType:               CollectorType.TypeClasslike,
	}
}

package configs

import "github.com/KoNekoD/go-deptrac/pkg/collectors"

type ClassLikeConfig struct {
	*collectors.ConfigurableCollectorConfig
	collectorType collectors.CollectorType
}

func NewClassLikeConfig(config string) *ClassLikeConfig {
	return &ClassLikeConfig{
		ConfigurableCollectorConfig: collectors.CreateConfigurableCollectorConfig(config),
		collectorType:               collectors.CollectorTypeTypeClasslike,
	}
}

package configs

import "github.com/KoNekoD/go-deptrac/pkg/collectors"

type GlobConfig struct {
	*collectors.ConfigurableCollectorConfig
	collectorType collectors.CollectorType
}

func NewGlobConfig(config string) *GlobConfig {
	return &GlobConfig{
		ConfigurableCollectorConfig: collectors.CreateConfigurableCollectorConfig(config),
		collectorType:               collectors.CollectorTypeTypeGlob,
	}
}

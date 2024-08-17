package GlobConfig

import (
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Config/CollectorType"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Config/ConfigurableCollectorConfig"
)

type GlobConfig struct {
	*ConfigurableCollectorConfig.ConfigurableCollectorConfig
	collectorType CollectorType.CollectorType
}

func NewGlobConfig(config string) *GlobConfig {
	return &GlobConfig{
		ConfigurableCollectorConfig: ConfigurableCollectorConfig.CreateConfigurableCollectorConfig(config),
		collectorType:               CollectorType.TypeGlob,
	}
}

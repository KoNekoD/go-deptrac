package collectors_configs

import (
	"github.com/KoNekoD/go-deptrac/pkg/domain/enums"
)

type GlobConfig struct {
	*ConfigurableCollectorConfig
	collectorType enums.CollectorType
}

func NewGlobConfig(config string) *GlobConfig {
	return &GlobConfig{
		ConfigurableCollectorConfig: CreateConfigurableCollectorConfig(config),
		collectorType:               enums.CollectorTypeTypeGlob,
	}
}

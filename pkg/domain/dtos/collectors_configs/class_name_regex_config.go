package collectors_configs

import (
	"github.com/KoNekoD/go-deptrac/pkg/domain/enums"
)

type ClassNameRegexConfig struct {
	*ConfigurableCollectorConfig
	collectorType enums.CollectorType
}

func NewClassNameRegexConfig(config string) *ClassNameRegexConfig {
	return &ClassNameRegexConfig{
		ConfigurableCollectorConfig: CreateConfigurableCollectorConfig(config),
		collectorType:               enums.CollectorTypeTypeClassNameRegex,
	}
}

package configs

import (
	"github.com/KoNekoD/go-deptrac/pkg/collectors"
	"github.com/KoNekoD/go-deptrac/pkg/domain/enums"
)

type ClassNameRegexConfig struct {
	*collectors.ConfigurableCollectorConfig
	collectorType enums.CollectorType
}

func NewClassNameRegexConfig(config string) *ClassNameRegexConfig {
	return &ClassNameRegexConfig{
		ConfigurableCollectorConfig: collectors.CreateConfigurableCollectorConfig(config),
		collectorType:               enums.CollectorTypeTypeClassNameRegex,
	}
}

package configs

import "github.com/KoNekoD/go-deptrac/pkg/collectors"

type ClassNameRegexConfig struct {
	*collectors.ConfigurableCollectorConfig
	collectorType collectors.CollectorType
}

func NewClassNameRegexConfig(config string) *ClassNameRegexConfig {
	return &ClassNameRegexConfig{
		ConfigurableCollectorConfig: collectors.CreateConfigurableCollectorConfig(config),
		collectorType:               collectors.CollectorTypeTypeClassNameRegex,
	}
}

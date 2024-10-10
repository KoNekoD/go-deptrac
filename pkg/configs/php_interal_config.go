package configs

import "github.com/KoNekoD/go-deptrac/pkg/collectors"

type PhpInteralConfig struct {
	*collectors.ConfigurableCollectorConfig
	collectorType collectors.CollectorType
}

func NewPhpInteralConfig(config string) *PhpInteralConfig {
	return &PhpInteralConfig{
		ConfigurableCollectorConfig: collectors.CreateConfigurableCollectorConfig(config),
		collectorType:               collectors.CollectorTypeTypePhpInternal,
	}
}

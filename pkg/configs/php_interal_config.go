package configs

import (
	"github.com/KoNekoD/go-deptrac/pkg/collectors"
	"github.com/KoNekoD/go-deptrac/pkg/domain/enums"
)

type PhpInteralConfig struct {
	*collectors.ConfigurableCollectorConfig
	collectorType enums.CollectorType
}

func NewPhpInteralConfig(config string) *PhpInteralConfig {
	return &PhpInteralConfig{
		ConfigurableCollectorConfig: collectors.CreateConfigurableCollectorConfig(config),
		collectorType:               enums.CollectorTypeTypePhpInternal,
	}
}

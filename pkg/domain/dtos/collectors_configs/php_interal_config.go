package collectors_configs

import (
	"github.com/KoNekoD/go-deptrac/pkg/domain/enums"
)

type PhpInteralConfig struct {
	*ConfigurableCollectorConfig
	collectorType enums.CollectorType
}

func NewPhpInteralConfig(config string) *PhpInteralConfig {
	return &PhpInteralConfig{
		ConfigurableCollectorConfig: CreateConfigurableCollectorConfig(config),
		collectorType:               enums.CollectorTypeTypePhpInternal,
	}
}

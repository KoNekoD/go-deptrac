package collectors_configs

import (
	"github.com/KoNekoD/go-deptrac/pkg/domain/enums"
)

type FunctionNameConfig struct {
	*ConfigurableCollectorConfig
	collectorType enums.CollectorType
}

func NewFunctionNameConfig(config string) *FunctionNameConfig {
	return &FunctionNameConfig{
		ConfigurableCollectorConfig: CreateConfigurableCollectorConfig(config),
		collectorType:               enums.CollectorTypeTypeFunctionName,
	}
}

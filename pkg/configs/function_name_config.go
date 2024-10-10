package configs

import "github.com/KoNekoD/go-deptrac/pkg/collectors"

type FunctionNameConfig struct {
	*collectors.ConfigurableCollectorConfig
	collectorType collectors.CollectorType
}

func NewFunctionNameConfig(config string) *FunctionNameConfig {
	return &FunctionNameConfig{
		ConfigurableCollectorConfig: collectors.CreateConfigurableCollectorConfig(config),
		collectorType:               collectors.CollectorTypeTypeFunctionName,
	}
}

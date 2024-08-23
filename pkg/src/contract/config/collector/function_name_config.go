package collector

import (
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/config"
)

type FunctionNameConfig struct {
	*config.ConfigurableCollectorConfig
	collectorType config.CollectorType
}

func NewFunctionNameConfig(config string) *FunctionNameConfig {
	return &FunctionNameConfig{
		ConfigurableCollectorConfig: config.CreateConfigurableCollectorConfig(config),
		collectorType:               config.TypeFunctionName,
	}
}

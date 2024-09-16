package collector

import (
	config_contract2 "github.com/KoNekoD/go-deptrac/pkg/config_contract"
)

type FunctionNameConfig struct {
	*config_contract2.ConfigurableCollectorConfig
	collectorType config_contract2.CollectorType
}

func NewFunctionNameConfig(config string) *FunctionNameConfig {
	return &FunctionNameConfig{
		ConfigurableCollectorConfig: config.CreateConfigurableCollectorConfig(config),
		collectorType:               config.TypeFunctionName,
	}
}

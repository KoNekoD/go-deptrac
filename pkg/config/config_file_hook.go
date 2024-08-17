package config

import "flag"

type configFileHook struct{}
type ConfigFileHook interface {
	GetConfigFile() string
}

func NewConfigFileHook() ConfigFileHook {
	return &configFileHook{}
}

const ArgNameConfigFileHook = "config-file"
const DefaultConfigFileHook = "deptrac.yaml"
const UsageConfigFileHook = "config file path"

func (h *configFileHook) GetConfigFile() string {
	configFile := flag.String(ArgNameConfigFileHook, DefaultConfigFileHook, UsageConfigFileHook)

	return *configFile
}

package configs

import "flag"

type configFileHook struct{}
type ConfigFileHook interface {
	GetConfigFile() string
}

func NewConfigFileHook() ConfigFileHook {
	return &configFileHook{}
}

const ArgNameConfigFileHook = "config"
const DefaultConfigFileHook = "deptrac.yaml"
const UsageConfigFileHook = "config_contract file_supportive path"

func (h *configFileHook) GetConfigFile() string {
	configFile := flag.String(ArgNameConfigFileHook, DefaultConfigFileHook, UsageConfigFileHook)

	return *configFile
}

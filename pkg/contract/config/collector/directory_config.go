package collector

import (
	"github.com/KoNekoD/go-deptrac/pkg/contract/config"
)

type DirectoryConfig struct {
	*config.ConfigurableCollectorConfig
	collectorType config.CollectorType
}

func NewDirectoryConfig(config string) *DirectoryConfig {
	return &DirectoryConfig{
		ConfigurableCollectorConfig: config.CreateConfigurableCollectorConfig(config),
		collectorType:               config.TypeDirectory,
	}
}

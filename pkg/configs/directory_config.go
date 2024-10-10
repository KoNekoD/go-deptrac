package configs

import "github.com/KoNekoD/go-deptrac/pkg/collectors"

type DirectoryConfig struct {
	*collectors.ConfigurableCollectorConfig
	collectorType collectors.CollectorType
}

func NewDirectoryConfig(config string) *DirectoryConfig {
	return &DirectoryConfig{
		ConfigurableCollectorConfig: collectors.CreateConfigurableCollectorConfig(config),
		collectorType:               collectors.CollectorTypeTypeDirectory,
	}
}

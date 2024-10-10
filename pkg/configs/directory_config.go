package configs

import (
	"github.com/KoNekoD/go-deptrac/pkg/collectors"
	"github.com/KoNekoD/go-deptrac/pkg/domain/enums"
)

type DirectoryConfig struct {
	*collectors.ConfigurableCollectorConfig
	collectorType enums.CollectorType
}

func NewDirectoryConfig(config string) *DirectoryConfig {
	return &DirectoryConfig{
		ConfigurableCollectorConfig: collectors.CreateConfigurableCollectorConfig(config),
		collectorType:               enums.CollectorTypeTypeDirectory,
	}
}

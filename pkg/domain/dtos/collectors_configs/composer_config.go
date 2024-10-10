package collectors_configs

import (
	"github.com/KoNekoD/go-deptrac/pkg/domain/enums"
)

type ComposerConfig struct {
	*CollectorConfig
	collectorType    enums.CollectorType
	packages         []string
	composerPath     string
	composerLockPath string
}

func NewComposerConfig(packages []string, composerPath *string, composerLockPath *string) *ComposerConfig {
	composerPathDefault := "composer.json"
	composerLockPathDefault := "composer.lock"

	if composerPath == nil {
		composerPath = &composerPathDefault
	}
	if composerLockPath == nil {
		composerLockPath = &composerLockPathDefault

	}

	return &ComposerConfig{
		CollectorConfig:  &CollectorConfig{},
		collectorType:    enums.CollectorTypeTypeComposer,
		packages:         packages,
		composerPath:     *composerPath,
		composerLockPath: *composerLockPath,
	}
}

func (c *ComposerConfig) ToArray() map[string]interface{} {
	parent := c.CollectorConfig.ToArray()

	parent["composerPath"] = c.composerPath
	parent["composerLockPath"] = c.composerLockPath
	parent["packages"] = c.packages

	return parent
}

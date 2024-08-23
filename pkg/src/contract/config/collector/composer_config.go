package collector

import (
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/config"
)

type ComposerConfig struct {
	*config.CollectorConfig
	collectorType    config.CollectorType
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
		CollectorConfig:  &config.CollectorConfig{},
		collectorType:    config.TypeComposer,
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

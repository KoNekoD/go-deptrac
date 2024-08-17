package ComposerConfig

import (
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Config/CollectorConfig"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Config/CollectorType"
)

type ComposerConfig struct {
	*CollectorConfig.CollectorConfig
	collectorType    CollectorType.CollectorType
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
		CollectorConfig:  &CollectorConfig.CollectorConfig{},
		collectorType:    CollectorType.TypeComposer,
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

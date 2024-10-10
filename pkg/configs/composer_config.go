package configs

import "github.com/KoNekoD/go-deptrac/pkg/collectors"

type ComposerConfig struct {
	*collectors.CollectorConfig
	collectorType    collectors.CollectorType
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
		CollectorConfig:  &collectors.CollectorConfig{},
		collectorType:    collectors.CollectorTypeTypeComposer,
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

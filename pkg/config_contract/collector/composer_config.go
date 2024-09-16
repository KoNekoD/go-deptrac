package collector

import (
	config_contract2 "github.com/KoNekoD/go-deptrac/pkg/config_contract"
)

type ComposerConfig struct {
	*config_contract2.CollectorConfig
	collectorType    config_contract2.CollectorType
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
		CollectorConfig:  &config_contract2.CollectorConfig{},
		collectorType:    config_contract2.TypeComposer,
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

package collector

import (
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/config"
)

type TagValueRegexConfig struct {
	*config.CollectorConfig
	collectorType config.CollectorType
	tag           string
	value         *string
}

func newTagValueRegexConfig(tag string, value *string) *TagValueRegexConfig {
	return &TagValueRegexConfig{
		CollectorConfig: &config.CollectorConfig{},
		collectorType:   config.TypeTagValueRegex,
		tag:             tag,
		value:           value,
	}
}

func CreateTagValueRegexConfig(tag string, regexpr *string) *TagValueRegexConfig {
	return newTagValueRegexConfig(tag, regexpr)
}

func (c *TagValueRegexConfig) Match(regexpr string) *TagValueRegexConfig {
	c.value = &regexpr
	return c
}

func (c *TagValueRegexConfig) ToArray() map[string]interface{} {
	parent := c.CollectorConfig.ToArray()

	parent["type"] = string(c.collectorType)
	parent["tag"] = c.tag
	parent["value"] = c.value

	return parent
}

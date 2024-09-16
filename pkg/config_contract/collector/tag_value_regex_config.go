package collector

import (
	config_contract2 "github.com/KoNekoD/go-deptrac/pkg/config_contract"
)

type TagValueRegexConfig struct {
	*config_contract2.CollectorConfig
	collectorType config_contract2.CollectorType
	tag           string
	value         *string
}

func newTagValueRegexConfig(tag string, value *string) *TagValueRegexConfig {
	return &TagValueRegexConfig{
		CollectorConfig: &config_contract2.CollectorConfig{},
		collectorType:   config_contract2.TypeTagValueRegex,
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

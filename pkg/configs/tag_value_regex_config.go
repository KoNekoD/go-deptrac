package configs

import "github.com/KoNekoD/go-deptrac/pkg/collectors"

type TagValueRegexConfig struct {
	*collectors.CollectorConfig
	collectorType collectors.CollectorType
	tag           string
	value         *string
}

func newTagValueRegexConfig(tag string, value *string) *TagValueRegexConfig {
	return &TagValueRegexConfig{
		CollectorConfig: &collectors.CollectorConfig{},
		collectorType:   collectors.CollectorTypeTypeTagValueRegex,
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

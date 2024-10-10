package configs

import (
	"github.com/KoNekoD/go-deptrac/pkg/collectors"
	"github.com/KoNekoD/go-deptrac/pkg/domain/enums"
)

type TagValueRegexConfig struct {
	*collectors.CollectorConfig
	collectorType enums.CollectorType
	tag           string
	value         *string
}

func newTagValueRegexConfig(tag string, value *string) *TagValueRegexConfig {
	return &TagValueRegexConfig{
		CollectorConfig: &collectors.CollectorConfig{},
		collectorType:   enums.CollectorTypeTypeTagValueRegex,
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

package dtos

import (
	"github.com/KoNekoD/go-deptrac/pkg/domain/enums"
)

type AnalyserConfig struct {
	Types       map[string]enums.EmitterType
	InternalTag *string
}

func newAnalyserConfig() *AnalyserConfig {
	return &AnalyserConfig{Types: make(map[string]enums.EmitterType)}
}

func Create(types []enums.EmitterType, internalTag *string) *AnalyserConfig {
	analyser := newAnalyserConfig()

	if types == nil {
		types = []enums.EmitterType{enums.EmitterTypeClassToken, enums.EmitterTypeFunctionToken}
	}

	analyser.setTypes(types...)
	analyser.setInternalTag(internalTag)

	return analyser
}

func (c *AnalyserConfig) setTypes(types ...enums.EmitterType) *AnalyserConfig {
	c.Types = make(map[string]enums.EmitterType)

	for _, emitterType := range types {
		c.Types[string(emitterType)] = emitterType
	}

	return c
}

func (c *AnalyserConfig) setInternalTag(internalTag *string) *AnalyserConfig {
	c.InternalTag = internalTag

	return c
}

func (c *AnalyserConfig) ToArray() map[string]interface{} {
	types := make([]string, len(c.Types))

	for _, value := range c.Types {
		types = append(types, string(value))
	}

	return map[string]interface{}{
		"types":       types,
		"internalTag": c.InternalTag,
	}
}

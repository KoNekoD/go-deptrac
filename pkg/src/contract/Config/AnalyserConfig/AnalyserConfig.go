package AnalyserConfig

import (
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Config/EmitterType"
)

type AnalyserConfig struct {
	Types       map[string]EmitterType.EmitterType
	InternalTag *string
}

func newAnalyserConfig() *AnalyserConfig {
	return &AnalyserConfig{Types: make(map[string]EmitterType.EmitterType)}
}

func Create(types []EmitterType.EmitterType, internalTag *string) *AnalyserConfig {
	analyser := newAnalyserConfig()

	if types == nil {
		types = []EmitterType.EmitterType{EmitterType.ClassToken, EmitterType.FunctionToken}
	}

	analyser.setTypes(types...)
	analyser.setInternalTag(internalTag)

	return analyser
}

func (c *AnalyserConfig) setTypes(types ...EmitterType.EmitterType) *AnalyserConfig {
	c.Types = make(map[string]EmitterType.EmitterType)

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

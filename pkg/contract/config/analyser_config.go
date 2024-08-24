package config

type AnalyserConfig struct {
	Types       map[string]EmitterType
	InternalTag *string
}

func newAnalyserConfig() *AnalyserConfig {
	return &AnalyserConfig{Types: make(map[string]EmitterType)}
}

func Create(types []EmitterType, internalTag *string) *AnalyserConfig {
	analyser := newAnalyserConfig()

	if types == nil {
		types = []EmitterType{ClassToken, FunctionToken}
	}

	analyser.setTypes(types...)
	analyser.setInternalTag(internalTag)

	return analyser
}

func (c *AnalyserConfig) setTypes(types ...EmitterType) *AnalyserConfig {
	c.Types = make(map[string]EmitterType)

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

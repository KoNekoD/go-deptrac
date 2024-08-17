package MermaidJsConfig

import (
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Config/Formatter/FormatterConfigInterface/FormatterType"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Config/Layer"
)

type MermaidJsConfig struct {
	name      string
	direction string
	groups    map[string][]*Layer.Layer
}

func CreateMermaidJsConfig() *MermaidJsConfig {
	return &MermaidJsConfig{
		name:      "mermaidjs",
		direction: "TD",
		groups:    make(map[string][]*Layer.Layer),
	}
}

func (m *MermaidJsConfig) GetName() FormatterType.FormatterType {
	return FormatterType.FormatterTypeMermaidJsConfig
}

func (m *MermaidJsConfig) Direction(direction string) *MermaidJsConfig {
	m.direction = direction
	return m
}

func (m *MermaidJsConfig) Groups(name string, layerConfigs ...*Layer.Layer) *MermaidJsConfig {
	for _, config := range layerConfigs {
		m.groups[name] = append(m.groups[name], config)
	}
	return m
}

func (m *MermaidJsConfig) ToArray() map[string]interface{} {
	output := make(map[string]interface{})
	if len(m.groups) > 0 {
		groups := make(map[string][]string)
		for key, configs := range m.groups {
			layerNames := make([]string, len(configs))
			i := 0
			for _, layer := range configs {
				layerNames[i] = layer.Name
				i++
			}
			groups[key] = layerNames
		}
		output["groups"] = groups
	}
	output["direction"] = m.direction
	return output
}

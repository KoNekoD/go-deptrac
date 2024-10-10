package configs

import (
	"github.com/KoNekoD/go-deptrac/pkg/domain/enums"
	"github.com/KoNekoD/go-deptrac/pkg/layers"
)

type MermaidJsConfig struct {
	name      string
	Direction string
	Groups    map[string][]*layers.Layer
}

func CreateMermaidJsConfig() *MermaidJsConfig {
	return &MermaidJsConfig{
		name:      "mermaidjs",
		Direction: "TD",
		Groups:    make(map[string][]*layers.Layer),
	}
}

func (m *MermaidJsConfig) GetName() enums.FormatterType {
	return enums.FormatterTypeMermaidJsConfig
}

func (m *MermaidJsConfig) SetDirection(direction string) *MermaidJsConfig {
	m.Direction = direction
	return m
}

func (m *MermaidJsConfig) SetGroups(name string, layerConfigs ...*layers.Layer) *MermaidJsConfig {
	for _, config := range layerConfigs {
		m.Groups[name] = append(m.Groups[name], config)
	}
	return m
}

func (m *MermaidJsConfig) ToArray() map[string]interface{} {
	output := make(map[string]interface{})
	if len(m.Groups) > 0 {
		groups := make(map[string][]string)
		for key, configs := range m.Groups {
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
	output["direction"] = m.Direction
	return output
}

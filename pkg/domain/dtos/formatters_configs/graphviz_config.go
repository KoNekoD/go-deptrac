package formatters_configs

import (
	"github.com/KoNekoD/go-deptrac/pkg/domain/dtos"
	"github.com/KoNekoD/go-deptrac/pkg/domain/enums"
)

type GraphvizConfig struct {
	name          string
	PointToGroups bool
	HiddenLayers  []*dtos.LayerConfig
	Groups        map[string][]*dtos.LayerConfig
}

func (g *GraphvizConfig) HasHiddenLayer(name string) bool {
	for _, layer := range g.HiddenLayers {
		if layer.Name == name {
			return true
		}
	}
	return false
}

func (g *GraphvizConfig) HiddenLayersNames() []string {
	names := make([]string, 0)

	for _, layer := range g.HiddenLayers {
		names = append(names, layer.Name)
	}

	return names
}

func newGraphvizConfig() *GraphvizConfig {
	return &GraphvizConfig{
		name:          "graphviz",
		PointToGroups: false,
		HiddenLayers:  make([]*dtos.LayerConfig, 0),
		Groups:        make(map[string][]*dtos.LayerConfig),
	}
}

func CreateGraphvizConfig() *GraphvizConfig {
	return newGraphvizConfig()
}

func (g *GraphvizConfig) SetPointToGroups(pointToGroups *bool) *GraphvizConfig {
	if pointToGroups == nil {
		pointsToGroupTmp := true
		pointToGroups = &pointsToGroupTmp
	}
	g.PointToGroups = *pointToGroups
	return g
}

func (g *GraphvizConfig) SetHiddenLayers(layerConfigs ...*dtos.LayerConfig) *GraphvizConfig {
	g.HiddenLayers = append(g.HiddenLayers, layerConfigs...)
	return g
}

func (g *GraphvizConfig) SetGroups(name string, layerConfigs ...*dtos.LayerConfig) *GraphvizConfig {
	g.Groups[name] = append(g.Groups[name], layerConfigs...)
	return g
}

func (g *GraphvizConfig) ToArray() map[string]interface{} {
	output := make(map[string]interface{})
	if len(g.HiddenLayers) > 0 {
		hiddenLayers := make([]string, len(g.HiddenLayers))
		i := 0
		for _, config := range g.HiddenLayers {
			hiddenLayers[i] = config.Name
			i++
		}
		output["hidden_layers"] = hiddenLayers
	}
	if len(g.Groups) > 0 {
		groups := make(map[string][]string)
		for key, configs := range g.Groups {
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
	output["point_to_groups"] = g.PointToGroups
	return output
}

func (g *GraphvizConfig) GetName() enums.FormatterType {
	return enums.FormatterTypeGraphvizConfig
}

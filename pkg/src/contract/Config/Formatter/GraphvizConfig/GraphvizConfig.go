package GraphvizConfig

import (
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Config/Formatter/FormatterConfigInterface/FormatterType"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Config/Layer"
)

type GraphvizConfig struct {
	name          string
	pointsToGroup bool
	hiddenLayers  []*Layer.Layer
	groups        map[string][]*Layer.Layer
}

func newGraphvizConfig() *GraphvizConfig {
	return &GraphvizConfig{
		name:          "graphviz",
		pointsToGroup: false,
		hiddenLayers:  make([]*Layer.Layer, 0),
		groups:        make(map[string][]*Layer.Layer),
	}
}

func CreateGraphvizConfig() *GraphvizConfig {
	return newGraphvizConfig()
}

func (g *GraphvizConfig) PointsToGroup(pointsToGroup *bool) *GraphvizConfig {
	if pointsToGroup == nil {
		pointsToGroupTmp := true
		pointsToGroup = &pointsToGroupTmp
	}
	g.pointsToGroup = *pointsToGroup
	return g
}

func (g *GraphvizConfig) HiddenLayers(layerConfigs ...*Layer.Layer) *GraphvizConfig {
	g.hiddenLayers = append(g.hiddenLayers, layerConfigs...)
	return g
}

func (g *GraphvizConfig) Groups(name string, layerConfigs ...*Layer.Layer) *GraphvizConfig {
	g.groups[name] = append(g.groups[name], layerConfigs...)
	return g
}

func (g *GraphvizConfig) ToArray() map[string]interface{} {
	output := make(map[string]interface{})
	if len(g.hiddenLayers) > 0 {
		hiddenLayers := make([]string, len(g.hiddenLayers))
		i := 0
		for _, config := range g.hiddenLayers {
			hiddenLayers[i] = config.Name
			i++
		}
		output["hidden_layers"] = hiddenLayers
	}
	if len(g.groups) > 0 {
		groups := make(map[string][]string)
		for key, configs := range g.groups {
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
	output["point_to_groups"] = g.pointsToGroup
	return output
}

func (g *GraphvizConfig) GetName() FormatterType.FormatterType {
	return FormatterType.FormatterTypeGraphvizConfig
}

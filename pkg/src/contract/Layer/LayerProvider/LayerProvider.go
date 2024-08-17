package LayerProvider

import (
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Config/Ruleset"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Layer/CircularReferenceException"
	"github.com/KoNekoD/go-deptrac/pkg/util"
	"strings"
)

type LayerProvider struct {
	AllowedLayers map[string]*Ruleset.Ruleset
}

func NewLayerProvider(allowedLayers map[string]*Ruleset.Ruleset) *LayerProvider {
	return &LayerProvider{AllowedLayers: allowedLayers}
}

func (l *LayerProvider) GetAllowedLayers(layerName string) ([]string, error) {
	return l.getTransitiveDependencies(layerName, []string{})
}

func (l *LayerProvider) getTransitiveDependencies(layerName string, previousLayers []string) ([]string, error) {
	if util.InArray(layerName, previousLayers) {
		return nil, CircularReferenceException.NewCircularReferenceExceptionFromCircularLayerDependency(layerName, previousLayers)
	}

	dependencies := make([]string, 0)

	ruleset, ok := l.AllowedLayers[layerName]
	if !ok {
		return dependencies, nil
	}

	allowedLayers := ruleset.AccessableLayers

	for _, layer := range allowedLayers {
		if strings.HasPrefix(layer.Name, "+") {
			dep, err := l.getTransitiveDependencies(layer.Name[1:], append(previousLayers, layerName))
			if err != nil {
				return nil, err
			}
			dependencies = append(dependencies, dep...)
		}

		dependencies = append(dependencies, layer.Name)
	}
	return dependencies, nil
}

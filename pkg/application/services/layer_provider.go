package services

import (
	"github.com/KoNekoD/go-deptrac/pkg/domain/apperrors"
	"github.com/KoNekoD/go-deptrac/pkg/domain/dtos/rules"
	"github.com/KoNekoD/go-deptrac/pkg/domain/utils"
	"strings"
)

type LayerProvider struct {
	AllowedLayers map[string]*rules.Ruleset
}

func NewLayerProvider(allowedLayers map[string]*rules.Ruleset) *LayerProvider {
	return &LayerProvider{AllowedLayers: allowedLayers}
}

func (l *LayerProvider) GetAllowedLayers(layerName string) ([]string, error) {
	return l.getTransitiveDependencies(layerName, []string{})
}

func (l *LayerProvider) getTransitiveDependencies(layerName string, previousLayers []string) ([]string, error) {
	if utils.InArray(layerName, previousLayers) {
		return nil, apperrors.NewCircularReferenceExceptionFromCircularLayerDependency(layerName, previousLayers)
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

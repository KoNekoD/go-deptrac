package layer

import (
	"errors"
	"github.com/KoNekoD/go-deptrac/pkg/contract/ast"
	"github.com/KoNekoD/go-deptrac/pkg/contract/config"
	"github.com/KoNekoD/go-deptrac/pkg/core/layer/collector"
	"github.com/KoNekoD/go-deptrac/pkg/core/layer/layer_resolver_interface"
	"reflect"
	"sync"
)

// LayerResolver - LayerResolverInterface defines the structure for a layer resolver
type LayerResolver struct {
	collectorResolver collector.CollectorResolverInterface
	layersConfig      []*config.Layer
	layers            map[string][]*collector.Collectable
	initialized       bool
	resolved          map[string]map[string]bool
	mu                sync.Mutex
}

// NewLayerResolver creates a new LayerResolverInterface
func NewLayerResolver(collectorResolver collector.CollectorResolverInterface, layersConfig []*config.Layer) layer_resolver_interface.LayerResolverInterface {
	return &LayerResolver{
		collectorResolver: collectorResolver,
		layersConfig:      layersConfig,
		layers:            make(map[string][]*collector.Collectable),
		resolved:          make(map[string]map[string]bool),
	}
}

// GetLayersForReference retrieves layers for a given reference
func (r *LayerResolver) GetLayersForReference(reference ast.TokenReferenceInterface) (map[string]bool, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if !r.initialized {
		err := r.initializeLayers()
		if err != nil {
			return nil, err
		}
	}

	// TODO: We need to correctly handle cases ( go/ast external packages and other )
	if reference == nil || reflect.ValueOf(reference).IsNil() {
		return make(map[string]bool), nil
	}

	tokenName := reference.GetToken().ToString()
	if resolvedLayers, ok := r.resolved[tokenName]; ok {
		return resolvedLayers, nil
	}

	r.resolved[tokenName] = make(map[string]bool)
	for layer, collectables := range r.layers {
		for _, collectable := range collectables {
			attributes := collectable.Attributes
			satisfied, err := collectable.Collector.Satisfy(attributes, reference)
			if err != nil {
				return nil, err
			}
			if satisfied {
				if _, exists := r.resolved[tokenName][layer]; exists && r.resolved[tokenName][layer] {
					continue
				}
				if private, ok := attributes["private"].(bool); ok && private {
					r.resolved[tokenName][layer] = false
				} else {
					r.resolved[tokenName][layer] = true
				}
			}
		}
	}
	return r.resolved[tokenName], nil
}

// IsReferenceInLayer checks if a reference is in a given layer
func (r *LayerResolver) IsReferenceInLayer(layer string, reference ast.TokenReferenceInterface) (bool, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if !r.initialized {
		err := r.initializeLayers()
		if err != nil {
			return false, err
		}
	}

	tokenName := reference.GetToken().ToString()
	if resolvedLayers, ok := r.resolved[tokenName]; ok && len(resolvedLayers) > 0 {
		_, exists := resolvedLayers[layer]
		return exists, nil
	}

	collectables, exists := r.layers[layer]
	if !exists {
		return false, nil
	}

	for _, collectable := range collectables {
		satisfied, err := collectable.Collector.Satisfy(collectable.Attributes, reference)
		if err != nil {
			return false, err
		}
		if satisfied {
			return true, nil
		}
	}

	return false, nil
}

// Has checks if a layer exists
func (r *LayerResolver) Has(layer string) (bool, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if !r.initialized {
		err := r.initializeLayers()
		if err != nil {
			return false, err
		}
	}
	_, exists := r.layers[layer]
	return exists, nil
}

// initializeLayers initializes the layers from the configuration
func (r *LayerResolver) initializeLayers() error {
	r.layers = make(map[string][]*collector.Collectable)
	for _, layer := range r.layersConfig {
		layerName := layer.Name
		if _, exists := r.layers[layerName]; exists {
			return errors.New("invalid layer definition: duplicate name " + layerName)
		}

		r.layers[layerName] = []*collector.Collectable{}
		for _, config := range layer.Collectors {
			resolvedCollector, err := r.collectorResolver.Resolve(config.ToArray())

			if err != nil {
				return err
			}

			r.layers[layerName] = append(r.layers[layerName], resolvedCollector)
		}

		if len(r.layers[layerName]) == 0 {
			return errors.New("invalid layer definition: collector required for " + layerName)
		}
	}

	if len(r.layers) == 0 {
		return errors.New("invalid layer definition: at least one layer is required")
	}

	r.initialized = true
	return nil
}
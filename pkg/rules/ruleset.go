package rules

import (
	"github.com/KoNekoD/go-deptrac/pkg/domain/dtos"
)

type Ruleset struct {
	LayerConfig      *dtos.LayerConfig
	AccessableLayers []*dtos.LayerConfig
}

func NewRuleset(layerConfig *dtos.LayerConfig, layerConfigs []*dtos.LayerConfig) *Ruleset {
	r := &Ruleset{LayerConfig: layerConfig}

	r.Accesses(layerConfigs...)

	return r
}

func NewForLayer(layerConfig *dtos.LayerConfig) *Ruleset {
	return &Ruleset{LayerConfig: layerConfig, AccessableLayers: make([]*dtos.LayerConfig, 0)}
}

func (r *Ruleset) Accesses(layerConfigs ...*dtos.LayerConfig) *Ruleset {
	for _, config := range layerConfigs {
		r.AccessableLayers = append(r.AccessableLayers, config)
	}

	return r
}

func (r *Ruleset) ToArray() map[string]interface{} {
	data := make([]map[string]interface{}, len(r.AccessableLayers))
	for i, layer := range r.AccessableLayers {
		data[i] = layer.ToArray()
	}

	return map[string]interface{}{
		"name":     r.LayerConfig.Name,
		"accesses": data,
	}
}

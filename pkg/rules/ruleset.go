package rules

import "github.com/KoNekoD/go-deptrac/pkg/layers"

type Ruleset struct {
	LayerConfig      *layers.Layer
	AccessableLayers []*layers.Layer
}

func NewRuleset(layerConfig *layers.Layer, layerConfigs []*layers.Layer) *Ruleset {
	r := &Ruleset{LayerConfig: layerConfig}

	r.Accesses(layerConfigs...)

	return r
}

func NewForLayer(layerConfig *layers.Layer) *Ruleset {
	return &Ruleset{LayerConfig: layerConfig, AccessableLayers: make([]*layers.Layer, 0)}
}

func (r *Ruleset) Accesses(layerConfigs ...*layers.Layer) *Ruleset {
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

package config_contract

type Ruleset struct {
	LayerConfig      *Layer
	AccessableLayers []*Layer
}

func NewRuleset(layerConfig *Layer, layerConfigs []*Layer) *Ruleset {
	r := &Ruleset{LayerConfig: layerConfig}

	r.Accesses(layerConfigs...)

	return r
}

func NewForLayer(layerConfig *Layer) *Ruleset {
	return &Ruleset{LayerConfig: layerConfig, AccessableLayers: make([]*Layer, 0)}
}

func (r *Ruleset) Accesses(layerConfigs ...*Layer) *Ruleset {
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

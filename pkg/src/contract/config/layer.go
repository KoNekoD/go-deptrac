package config

type Layer struct {
	Collectors []*CollectorConfig
	Name       string
}

func NewLayer(name string, collectorConfigs []*CollectorConfig) *Layer {
	l := &Layer{Name: name}

	l.setCollectors(collectorConfigs...)

	return l
}

func NewLayerWithName(name string) *Layer {
	return &Layer{Name: name}
}

func (l *Layer) setCollectors(collectorConfigs ...*CollectorConfig) *Layer {
	for _, collectorConfig := range collectorConfigs {
		l.Collectors = append(l.Collectors, collectorConfig)
	}

	return l
}

func (l *Layer) ToArray() map[string]interface{} {
	collectors := make([]interface{}, len(l.Collectors))
	for i, collector := range l.Collectors {
		collectors[i] = collector.ToArray()
	}

	return map[string]interface{}{
		"name":       l.Name,
		"collectors": collectors,
	}
}

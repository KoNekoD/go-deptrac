package dtos

import "github.com/KoNekoD/go-deptrac/pkg/domain/dtos/collectors_configs"

type LayerConfig struct {
	Collectors []*collectors_configs.CollectorConfig
	Name       string
}

func NewLayer(name string, collectorConfigs []*collectors_configs.CollectorConfig) *LayerConfig {
	l := &LayerConfig{Name: name}

	l.setCollectors(collectorConfigs...)

	return l
}

func NewLayerWithName(name string) *LayerConfig {
	return &LayerConfig{Name: name}
}

func (l *LayerConfig) setCollectors(collectorConfigs ...*collectors_configs.CollectorConfig) *LayerConfig {
	for _, collectorConfig := range collectorConfigs {
		l.Collectors = append(l.Collectors, collectorConfig)
	}

	return l
}

func (l *LayerConfig) ToArray() map[string]interface{} {
	collectors := make([]interface{}, len(l.Collectors))
	for i, collector := range l.Collectors {
		collectors[i] = collector.ToArray()
	}

	return map[string]interface{}{"name": l.Name, "collectors": collectors}
}

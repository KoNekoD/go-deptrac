package configuration

type ConfigurationGraphViz struct {
	hiddenLayers   []string
	groupsLayerMap map[string][]string
	pointToGroups  bool
}

func NewConfigurationGraphVizFromArray(array map[string]interface{}) *ConfigurationGraphViz {
	return newConfigurationGraphViz(
		array["hidden_layers"].([]string),
		array["groups"].(map[string][]string),
		array["point_to_groups"].(bool),
	)
}

func newConfigurationGraphViz(hiddenLayers []string, groupsLayerMap map[string][]string, pointToGroups bool) *ConfigurationGraphViz {
	return &ConfigurationGraphViz{hiddenLayers: hiddenLayers, groupsLayerMap: groupsLayerMap, pointToGroups: pointToGroups}
}

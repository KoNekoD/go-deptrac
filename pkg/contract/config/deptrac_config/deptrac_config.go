package deptrac_config

import (
	"errors"
	"github.com/KoNekoD/go-deptrac/pkg/contract/config"
	"github.com/KoNekoD/go-deptrac/pkg/contract/config/formatter"
	Layer2 "github.com/KoNekoD/go-deptrac/pkg/contract/layer"
	"github.com/KoNekoD/go-deptrac/pkg/util"
)

type DeptracConfig struct {
	Paths                          []string
	Analyser                       *config.AnalyserConfig
	Formatters                     map[formatter.FormatterType]formatter.FormatterConfigInterface
	Layers                         []*config.Layer
	Rulesets                       map[string]*config.Ruleset
	IgnoreUncoveredInternalStructs bool
	SkipViolations                 map[string][]string
	ExcludeFiles                   []string
	CacheFile                      *string
}

func NewDeptracConfig(parsed map[string]interface{}) (*DeptracConfig, error) {
	parsedDeptrac := parsed["deptrac"].(map[string]interface{})

	formatters := make(map[formatter.FormatterType]formatter.FormatterConfigInterface)
	layers := make([]*config.Layer, 0)

	for _, layerRawRaw := range parsedDeptrac["layers"].([]interface{}) {
		layerRaw := layerRawRaw.(map[string]interface{})
		collectorConfigs := make([]*config.CollectorConfig, 0)

		for _, collectorRawRaw := range layerRaw["collectors"].([]interface{}) {
			collectorRaw := collectorRawRaw.(map[string]interface{})

			if !util.MapKeyExists(collectorRaw, "type") || !util.MapKeyIsString(collectorRaw, "type") {
				return nil, Layer2.NewInvalidCollectorDefinitionExceptionMissingType()
			}

			collectorType, err := config.NewCollectorTypeFromString(collectorRaw["type"].(string))

			if err != nil {
				return nil, err
			}

			privateValue, ok := collectorRaw["private"].(bool)
			private := false
			if ok {
				private = privateValue
			}
			payload := collectorRaw

			// Delete private and type
			delete(payload, "private")
			delete(payload, "type")

			collectorConfig := config.NewCollectorConfig(
				collectorType,
				payload,
				private,
			)

			collectorConfigs = append(collectorConfigs, collectorConfig)
		}

		layerName, ok := layerRaw["name"]
		if !ok {
			return nil, errors.New("invalid layer definition: missing name")
		}
		layerNameStr, ok := layerName.(string)
		if !ok {
			return nil, errors.New("invalid layer definition: name must be a string")
		}

		layer := config.NewLayer(
			layerNameStr,
			collectorConfigs,
		)

		layers = append(layers, layer)
	}

	if parsedDeptracFormatters, ok := parsedDeptrac["formatters"]; ok {
		for formatterKey, formatterRawRaw := range parsedDeptracFormatters.(map[string]interface{}) {
			formatterRaw := formatterRawRaw.(map[string]interface{})
			switch formatterKey {
			case string(formatter.FormatterTypeCodeclimateConfig):
				formatters[formatter.FormatterTypeCodeclimateConfig] = formatter.CreateCodeclimateConfig(
					formatterRaw["failure"].(*config.CodeclimateLevelEnum),
					formatterRaw["skipped"].(*config.CodeclimateLevelEnum),
					formatterRaw["uncovered"].(*config.CodeclimateLevelEnum),
				)
			case string(formatter.FormatterTypeGraphvizConfig):
				hiddenLayers := make([]*config.Layer, 0)

				for _, hiddenLayer := range formatterRaw["hiddenLayers"].([]string) {
					for _, layer := range layers {
						if layer.Name == hiddenLayer {
							hiddenLayers = append(hiddenLayers, layer)
							break
						}
					}
				}

				formatterGraphvizConfig := formatter.CreateGraphvizConfig().
					PointsToGroup(formatterRaw["pointsToGroup"].(*bool)).
					HiddenLayers(hiddenLayers...)

				formatters[formatter.FormatterTypeGraphvizConfig] = formatterGraphvizConfig

				for groupLayerName, groupRaw := range formatterRaw["groups"].(map[string][]string) {
					groupLayer := make([]*config.Layer, 0)

					for _, layerName := range groupRaw {
						for _, layer := range layers {
							if layer.Name == layerName {
								groupLayer = append(groupLayer, layer)
								break
							}
						}
					}

					formatterGraphvizConfig.Groups(groupLayerName, groupLayer...)
				}
			case string(formatter.FormatterTypeMermaidJsConfig):
				formatterMermaidJsConfig := formatter.CreateMermaidJsConfig().
					Direction(formatterRaw["direction"].(string))

				formatters[formatter.FormatterTypeMermaidJsConfig] = formatterMermaidJsConfig

				for groupLayerName, groupRaw := range formatterRaw["groups"].(map[string][]string) {
					groupLayer := make([]*config.Layer, 0)

					for _, layerName := range groupRaw {
						for _, layer := range layers {
							if layer.Name == layerName {
								groupLayer = append(groupLayer, layer)
								break
							}
						}
					}

					formatterMermaidJsConfig.Groups(groupLayerName, groupLayer...)
				}
			}
		}
	}

	rulesets := make(map[string]*config.Ruleset)

	for rulesetLayerName, rulesetLayersNames := range parsedDeptrac["ruleset"].(map[string]interface{}) {
		var rulesetOwningLayer *config.Layer

		for _, layer := range layers {
			if layer.Name == rulesetLayerName {
				rulesetOwningLayer = layer
				break
			}
		}

		rulesetLayers := make([]*config.Layer, 0)

		if rulesetLayersNames != nil { // If not ~
			for _, layerNameRaw := range rulesetLayersNames.([]interface{}) {
				layerName := layerNameRaw.(string)
				for _, layer := range layers {
					if layer.Name == layerName {
						rulesetLayers = append(rulesetLayers, layer)
						break
					}
				}
			}
		}

		ruleset := config.NewRuleset(rulesetOwningLayer, rulesetLayers)

		rulesets[rulesetLayerName] = ruleset
	}

	analyzerTypesDefault := []config.EmitterType{config.ClassToken, config.FunctionToken}
	analyzerTypes := make([]config.EmitterType, 0)
	internalTag := "@internal"
	if parsedDeptracAnalyzer, ok := parsedDeptrac["analyzer"]; ok {
		analyzerRaw := parsedDeptracAnalyzer.(map[string]interface{})
		for _, typeRaw := range analyzerRaw["types"].([]interface{}) {
			analyzerType, err := config.NewEmitterTypeFromString(typeRaw.(string))

			if err != nil {
				return nil, err
			}

			analyzerTypes = append(analyzerTypes, analyzerType)
		}
		internalTag = analyzerRaw["internal_tag"].(string)
	} else {
		analyzerTypes = analyzerTypesDefault
	}

	analyser := config.Create(analyzerTypes, &internalTag)

	paths := make([]string, 0)
	for _, path := range parsedDeptrac["paths"].([]interface{}) {
		paths = append(paths, path.(string))
	}

	ignoreUncoveredInternalStructs := true
	if v, ok := parsedDeptrac["ignore_uncovered_internal_structs"]; ok {
		ignoreUncoveredInternalStructs = v.(bool)
	}

	skipViolations := make(map[string][]string)
	if v, ok := parsedDeptrac["skip_violations"]; ok {
		skipViolations = v.(map[string][]string)
	}

	excludeFiles := make([]string, 0)
	if v, ok := parsedDeptrac["exclude_files"]; ok {
		excludeFiles = v.([]string)
	}

	var cacheFile *string
	if v, ok := parsedDeptrac["cache_file"]; ok {
		vStr := v.(string)
		cacheFile = &vStr
	}

	return &DeptracConfig{
		Paths:                          paths,
		Analyser:                       analyser,
		Formatters:                     formatters,
		Layers:                         layers,
		Rulesets:                       rulesets,
		IgnoreUncoveredInternalStructs: ignoreUncoveredInternalStructs,
		SkipViolations:                 skipViolations,
		ExcludeFiles:                   excludeFiles,
		CacheFile:                      cacheFile,
	}, nil
}

func (c *DeptracConfig) SetRulesets(rulesets ...*config.Ruleset) *DeptracConfig {
	for _, ruleset := range rulesets {
		c.Rulesets[ruleset.LayerConfig.Name] = ruleset
	}
	return c
}

func (c *DeptracConfig) ToArray() map[string]interface{} {
	config := make(map[string]interface{})

	if len(c.Paths) > 0 {
		config["paths"] = c.Paths
	}
	if c.Analyser != nil {
		config["analyser"] = c.Analyser.ToArray()
	}
	if len(c.Formatters) > 0 {
		formatters := make([]map[string]interface{}, len(c.Formatters))
		i := 0
		for _, formatter := range c.Formatters {
			formatters[i] = formatter.ToArray()
			i++
		}
		config["formatters"] = formatters
	}
	if len(c.ExcludeFiles) > 0 {
		config["exclude_files"] = c.ExcludeFiles
	}
	if len(c.Layers) > 0 {
		layers := make([]map[string]interface{}, len(c.Layers))
		i := 0
		for _, layer := range c.Layers {
			layers[i] = layer.ToArray()
			i++
		}
		config["layers"] = layers
	}
	if len(c.Rulesets) > 0 {
		rulesets := make([]map[string]interface{}, len(c.Rulesets))
		i := 0
		for _, ruleset := range c.Rulesets {
			rulesets[i] = ruleset.ToArray()
			i++
		}
		config["ruleset"] = rulesets
	}
	if len(c.SkipViolations) > 0 {
		config["skip_violations"] = c.SkipViolations
	}
	config["ignore_uncovered_internal_structs"] = c.IgnoreUncoveredInternalStructs
	config["cache_file"] = c.CacheFile
	return config
}

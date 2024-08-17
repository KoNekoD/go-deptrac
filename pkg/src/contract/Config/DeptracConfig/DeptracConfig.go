package DeptracConfig

import (
	"errors"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Config/AnalyserConfig"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Config/CodeclimateLevelEnum"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Config/CollectorConfig"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Config/CollectorType"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Config/EmitterType"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Config/Formatter/CodeclimateConfig"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Config/Formatter/FormatterConfigInterface"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Config/Formatter/FormatterConfigInterface/FormatterType"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Config/Formatter/GraphvizConfig"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Config/Formatter/MermaidJsConfig"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Config/Layer"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Config/Ruleset"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Layer/InvalidCollectorDefinitionException"
	"github.com/KoNekoD/go-deptrac/pkg/util"
)

type DeptracConfig struct {
	Paths                          []string
	Analyser                       *AnalyserConfig.AnalyserConfig
	Formatters                     map[FormatterType.FormatterType]FormatterConfigInterface.FormatterConfigInterface
	Layers                         []*Layer.Layer
	Rulesets                       map[string]*Ruleset.Ruleset
	IgnoreUncoveredInternalStructs bool
	SkipViolations                 map[string][]string
	ExcludeFiles                   []string
	CacheFile                      *string
}

func NewDeptracConfig(parsed map[string]interface{}) (*DeptracConfig, error) {
	parsedDeptrac := parsed["deptrac"].(map[string]interface{})

	formatters := make(map[FormatterType.FormatterType]FormatterConfigInterface.FormatterConfigInterface)
	layers := make([]*Layer.Layer, 0)

	for _, layerRawRaw := range parsedDeptrac["layers"].([]interface{}) {
		layerRaw := layerRawRaw.(map[string]interface{})
		collectorConfigs := make([]*CollectorConfig.CollectorConfig, 0)

		for _, collectorRawRaw := range layerRaw["collectors"].([]interface{}) {
			collectorRaw := collectorRawRaw.(map[string]interface{})

			if !util.MapKeyExists(collectorRaw, "type") || !util.MapKeyIsString(collectorRaw, "type") {
				return nil, InvalidCollectorDefinitionException.NewInvalidCollectorDefinitionExceptionMissingType()
			}

			collectorType, err := CollectorType.NewCollectorTypeFromString(collectorRaw["type"].(string))

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

			collectorConfig := CollectorConfig.NewCollectorConfig(
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

		layer := Layer.NewLayer(
			layerNameStr,
			collectorConfigs,
		)

		layers = append(layers, layer)
	}

	for formatterKey, formatterRawRaw := range parsedDeptrac["formatters"].(map[string]interface{}) {
		formatterRaw := formatterRawRaw.(map[string]interface{})
		switch formatterKey {
		case string(FormatterType.FormatterTypeCodeclimateConfig):
			formatters[FormatterType.FormatterTypeCodeclimateConfig] = CodeclimateConfig.CreateCodeclimateConfig(
				formatterRaw["failure"].(*CodeclimateLevelEnum.CodeclimateLevelEnum),
				formatterRaw["skipped"].(*CodeclimateLevelEnum.CodeclimateLevelEnum),
				formatterRaw["uncovered"].(*CodeclimateLevelEnum.CodeclimateLevelEnum),
			)
		case string(FormatterType.FormatterTypeGraphvizConfig):
			hiddenLayers := make([]*Layer.Layer, 0)

			for _, hiddenLayer := range formatterRaw["hiddenLayers"].([]string) {
				for _, layer := range layers {
					if layer.Name == hiddenLayer {
						hiddenLayers = append(hiddenLayers, layer)
						break
					}
				}
			}

			formatter := GraphvizConfig.CreateGraphvizConfig().
				PointsToGroup(formatterRaw["pointsToGroup"].(*bool)).
				HiddenLayers(hiddenLayers...)

			formatters[FormatterType.FormatterTypeGraphvizConfig] = formatter

			for groupLayerName, groupRaw := range formatterRaw["groups"].(map[string][]string) {
				groupLayer := make([]*Layer.Layer, 0)

				for _, layerName := range groupRaw {
					for _, layer := range layers {
						if layer.Name == layerName {
							groupLayer = append(groupLayer, layer)
							break
						}
					}
				}

				formatter.Groups(groupLayerName, groupLayer...)
			}
		case string(FormatterType.FormatterTypeMermaidJsConfig):
			formatter := MermaidJsConfig.CreateMermaidJsConfig().
				Direction(formatterRaw["direction"].(string))

			formatters[FormatterType.FormatterTypeMermaidJsConfig] = formatter

			for groupLayerName, groupRaw := range formatterRaw["groups"].(map[string][]string) {
				groupLayer := make([]*Layer.Layer, 0)

				for _, layerName := range groupRaw {
					for _, layer := range layers {
						if layer.Name == layerName {
							groupLayer = append(groupLayer, layer)
							break
						}
					}
				}

				formatter.Groups(groupLayerName, groupLayer...)
			}
		}
	}

	rulesets := make(map[string]*Ruleset.Ruleset)

	for rulesetLayerName, rulesetLayersNames := range parsedDeptrac["ruleset"].(map[string]interface{}) {
		var rulesetOwningLayer *Layer.Layer

		for _, layer := range layers {
			if layer.Name == rulesetLayerName {
				rulesetOwningLayer = layer
				break
			}
		}

		rulesetLayers := make([]*Layer.Layer, 0)

		for _, layerNameRaw := range rulesetLayersNames.([]interface{}) {
			layerName := layerNameRaw.(string)
			for _, layer := range layers {
				if layer.Name == layerName {
					rulesetLayers = append(rulesetLayers, layer)
					break
				}
			}
		}

		ruleset := Ruleset.NewRuleset(rulesetOwningLayer, rulesetLayers)

		rulesets[rulesetLayerName] = ruleset
	}

	analyzerRaw := parsedDeptrac["analyser"].(map[string]interface{})
	analyzerTypes := make([]EmitterType.EmitterType, 0)

	for _, typeRaw := range analyzerRaw["types"].([]interface{}) {
		analyzerType, err := EmitterType.NewEmitterTypeFromString(typeRaw.(string))

		if err != nil {
			return nil, err
		}

		analyzerTypes = append(analyzerTypes, analyzerType)
	}

	internalTag := analyzerRaw["internal_tag"].(string)

	analyser := AnalyserConfig.Create(
		analyzerTypes,
		&internalTag,
	)

	paths := make([]string, 0)
	for _, path := range parsedDeptrac["paths"].([]interface{}) {
		paths = append(paths, path.(string))
	}

	ignoreUncoveredInternalStructs := false
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

func (c *DeptracConfig) SetRulesets(rulesets ...*Ruleset.Ruleset) *DeptracConfig {
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

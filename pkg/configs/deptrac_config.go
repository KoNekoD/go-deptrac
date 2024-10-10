package configs

import (
	"github.com/KoNekoD/go-deptrac/pkg/collectors"
	"github.com/KoNekoD/go-deptrac/pkg/domain/apperrors"
	enums2 "github.com/KoNekoD/go-deptrac/pkg/domain/enums"
	"github.com/KoNekoD/go-deptrac/pkg/domain/utils"
	"github.com/KoNekoD/go-deptrac/pkg/formatters"
	"github.com/KoNekoD/go-deptrac/pkg/layers"
	"github.com/KoNekoD/go-deptrac/pkg/rules"
	"github.com/pkg/errors"
)

type DeptracConfig struct {
	Paths                          []string
	Analyser                       *AnalyserConfig
	Formatters                     map[enums2.FormatterType]formatters.FormatterConfigInterface
	Layers                         []*layers.Layer
	Rulesets                       map[string]*rules.Ruleset
	IgnoreUncoveredInternalStructs bool
	SkipViolations                 map[string][]string
	ExcludeFiles                   []string
	CacheFile                      *string
}

func NewDeptracConfig(parsed map[string]interface{}) (*DeptracConfig, error) {
	deptracConfig := &DeptracConfig{}
	if parsedDeptrac, ok := parsed["deptrac"]; ok {
		if err := deptracConfig.SetupDeptracData(parsedDeptrac); err != nil {
			return nil, errors.WithStack(err)
		}
	}

	return deptracConfig, nil
}

func (c *DeptracConfig) ToArray() map[string]interface{} {
	config := make(map[string]interface{})

	if len(c.Paths) > 0 {
		config["paths"] = c.Paths
	}
	if c.Analyser != nil {
		config["analyser_contract"] = c.Analyser.ToArray()
	}
	if len(c.Formatters) > 0 {
		formattersList := make([]map[string]interface{}, len(c.Formatters))
		i := 0
		for _, formatter := range c.Formatters {
			formattersList[i] = formatter.ToArray()
			i++
		}
		config["formatters"] = formattersList
	}
	if len(c.ExcludeFiles) > 0 {
		config["exclude_files"] = c.ExcludeFiles
	}
	if len(c.Layers) > 0 {
		layersList := make([]map[string]interface{}, len(c.Layers))
		i := 0
		for _, layer := range c.Layers {
			layersList[i] = layer.ToArray()
			i++
		}
		config["layers"] = layersList
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

func (c *DeptracConfig) SetupDeptracData(parsedDeptrac interface{}) error {
	if deptracMapData, ok := parsedDeptrac.(map[string]interface{}); ok {
		if err := c.SetupDeptracMapData(deptracMapData); err != nil {
			return errors.WithStack(err)
		}
	}

	return nil
}

func (c *DeptracConfig) SetupDeptracMapData(data map[string]interface{}) error {
	if layersList, ok := data["layers"]; ok {
		if err := c.SetupLayersData(layersList); err != nil {
			return errors.WithStack(err)
		}
	}
	layersList := c.Layers

	formattersList := make(map[enums2.FormatterType]formatters.FormatterConfigInterface)
	if parsedDeptracFormatters, ok := data["formatters"]; ok {
		for formatterKey, formatterRawRaw := range parsedDeptracFormatters.(map[string]interface{}) {
			formatterRaw := formatterRawRaw.(map[string]interface{})
			switch formatterKey {
			case string(enums2.FormatterTypeCodeclimateConfig):
				formattersList[enums2.FormatterTypeCodeclimateConfig] = CreateCodeclimateConfig(
					formatterRaw["failure"].(*enums2.CodeclimateLevelEnum),
					formatterRaw["skipped"].(*enums2.CodeclimateLevelEnum),
					formatterRaw["uncovered"].(*enums2.CodeclimateLevelEnum),
				)
			case string(enums2.FormatterTypeGraphvizConfig):
				hiddenLayers := make([]*layers.Layer, 0)

				for _, hiddenLayer := range formatterRaw["hiddenLayers"].([]string) {
					for _, layer := range layersList {
						if layer.Name == hiddenLayer {
							hiddenLayers = append(hiddenLayers, layer)
							break
						}
					}
				}

				formatterGraphvizConfig := CreateGraphvizConfig().
					SetPointToGroups(formatterRaw["point_to_groups"].(*bool)).
					SetHiddenLayers(hiddenLayers...)

				formattersList[enums2.FormatterTypeGraphvizConfig] = formatterGraphvizConfig

				for groupLayerName, groupRaw := range formatterRaw["groups"].(map[string][]string) {
					groupLayer := make([]*layers.Layer, 0)

					for _, layerName := range groupRaw {
						for _, layer := range layersList {
							if layer.Name == layerName {
								groupLayer = append(groupLayer, layer)
								break
							}
						}
					}

					formatterGraphvizConfig.SetGroups(groupLayerName, groupLayer...)
				}
			case string(enums2.FormatterTypeMermaidJsConfig):
				formatterMermaidJsConfig := CreateMermaidJsConfig().
					SetDirection(formatterRaw["direction"].(string))

				formattersList[enums2.FormatterTypeMermaidJsConfig] = formatterMermaidJsConfig

				for groupLayerName, groupRaw := range formatterRaw["groups"].(map[string][]string) {
					groupLayer := make([]*layers.Layer, 0)

					for _, layerName := range groupRaw {
						for _, layer := range layersList {
							if layer.Name == layerName {
								groupLayer = append(groupLayer, layer)
								break
							}
						}
					}

					formatterMermaidJsConfig.SetGroups(groupLayerName, groupLayer...)
				}
			}
		}
	}

	rulesets := make(map[string]*rules.Ruleset)

	if rulesetsData, ok := data["ruleset"]; ok {
		for rulesetLayerName, rulesetLayersNames := range rulesetsData.(map[string]interface{}) {
			var rulesetOwningLayer *layers.Layer

			for _, layer := range layersList {
				if layer.Name == rulesetLayerName {
					rulesetOwningLayer = layer
					break
				}
			}

			rulesetLayers := make([]*layers.Layer, 0)

			if rulesetLayersNames != nil { // If not ~
				for _, layerNameRaw := range rulesetLayersNames.([]interface{}) {
					layerName := layerNameRaw.(string)
					for _, layer := range layersList {
						if layer.Name == layerName {
							rulesetLayers = append(rulesetLayers, layer)
							break
						}
					}
				}
			}

			ruleset := rules.NewRuleset(rulesetOwningLayer, rulesetLayers)

			rulesets[rulesetLayerName] = ruleset
		}
	}

	analyzerTypesDefault := []enums2.EmitterType{enums2.EmitterTypeClassToken, enums2.EmitterTypeFunctionToken}
	analyzerTypes := make([]enums2.EmitterType, 0)
	internalTag := "@internal"
	if parsedDeptracAnalyzer, ok := data["analyzer"]; ok {
		analyzerRaw := parsedDeptracAnalyzer.(map[string]interface{})
		for _, typeRaw := range analyzerRaw["types"].([]interface{}) {
			analyzerType, err := enums2.NewEmitterTypeFromString(typeRaw.(string))
			if err != nil {
				return errors.WithStack(err)
			}

			analyzerTypes = append(analyzerTypes, analyzerType)
		}
		internalTag = analyzerRaw["internal_tag"].(string)
	} else {
		analyzerTypes = analyzerTypesDefault
	}

	analyser := Create(analyzerTypes, &internalTag)

	paths := make([]string, 0)
	if dataPaths, ok := data["paths"]; ok {
		for _, path := range dataPaths.([]interface{}) {
			paths = append(paths, path.(string))
		}
	}

	ignoreUncoveredInternalStructs := true
	if v, ok := data["ignore_uncovered_internal_structs"]; ok {
		ignoreUncoveredInternalStructs = v.(bool)
	}

	skipViolations := make(map[string][]string)
	if v, ok := data["skip_violations"]; ok {
		skipViolations = v.(map[string][]string)
	}

	excludeFiles := make([]string, 0)
	if v, ok := data["exclude_files"]; ok {
		excludeFiles = v.([]string)
	}

	var cacheFile *string
	if v, ok := data["cache_file"]; ok {
		vStr := v.(string)
		cacheFile = &vStr
	}

	c.Paths = paths
	c.Analyser = analyser
	c.Formatters = formattersList
	c.Rulesets = rulesets
	c.IgnoreUncoveredInternalStructs = ignoreUncoveredInternalStructs
	c.SkipViolations = skipViolations
	c.ExcludeFiles = excludeFiles
	c.CacheFile = cacheFile

	return nil
}

func (c *DeptracConfig) SetupLayersData(layers interface{}) error {
	if layersList, ok := layers.([]interface{}); ok {
		if err := c.SetupLayersListData(layersList); err != nil {
			return errors.WithStack(err)
		}
	}

	return nil
}

func (c *DeptracConfig) SetupLayersListData(list []interface{}) error {
	layersList := make([]*layers.Layer, 0)

	for _, layerRawRaw := range list {
		layerRaw := layerRawRaw.(map[string]interface{})
		collectorConfigs := make([]*collectors.CollectorConfig, 0)

		for _, collectorRawRaw := range layerRaw["collectors"].([]interface{}) {
			collectorRaw := collectorRawRaw.(map[string]interface{})

			if !utils.MapKeyExists(collectorRaw, "type") || !utils.MapKeyIsString(collectorRaw, "type") {
				return apperrors.NewInvalidCollectorDefinitionMissingType()
			}

			collectorType, err := enums2.NewCollectorTypeFromString(collectorRaw["type"].(string))
			if err != nil {
				return errors.WithStack(err)
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

			collectorConfig := collectors.NewCollectorConfig(collectorType, payload, private)

			collectorConfigs = append(collectorConfigs, collectorConfig)
		}

		layerName, ok := layerRaw["name"]
		if !ok {
			return errors.New("invalid layer definition: missing name")
		}
		layerNameStr, ok := layerName.(string)
		if !ok {
			return errors.New("invalid layer definition: name must be a string")
		}

		layer := layers.NewLayer(layerNameStr, collectorConfigs)

		layersList = append(layersList, layer)
	}
	c.Layers = layersList

	return nil
}

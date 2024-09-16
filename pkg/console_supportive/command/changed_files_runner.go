package command

import (
	analyser_core2 "github.com/KoNekoD/go-deptrac/pkg/analyser_core"
	output_formatter_contract2 "github.com/KoNekoD/go-deptrac/pkg/output_formatter_contract"
	result_contract2 "github.com/KoNekoD/go-deptrac/pkg/result_contract"
	"github.com/KoNekoD/go-deptrac/pkg/result_contract/output_result"
	"golang.org/x/exp/maps"
	"strings"
)

// ChangedFilesRunner - Should only be used by ChangedFilesCommand
type ChangedFilesRunner struct {
	layerForTokenAnalyser    *analyser_core2.LayerForTokenAnalyser
	dependencyLayersAnalyser *analyser_core2.DependencyLayersAnalyser
}

func NewChangedFilesRunner(layerForTokenAnalyser *analyser_core2.LayerForTokenAnalyser, dependencyLayersAnalyser *analyser_core2.DependencyLayersAnalyser) *ChangedFilesRunner {
	return &ChangedFilesRunner{
		layerForTokenAnalyser:    layerForTokenAnalyser,
		dependencyLayersAnalyser: dependencyLayersAnalyser,
	}
}

func (r *ChangedFilesRunner) Run(files []string, withDependencies bool, output output_formatter_contract2.OutputInterface) error {
	layers := make(map[string]string)
	for _, file := range files {
		matches, err := r.layerForTokenAnalyser.FindLayerForToken(file, analyser_core2.TokenTypeFile)
		if err != nil {
			return NewCommandRunExceptionAnalyserException(err)
		}
		for _, match := range matches {
			for _, layer := range match {
				layers[layer] = layer
			}
		}
	}
	output.WriteLineFormatted(output_formatter_contract2.StringOrArrayOfStrings{String: strings.Join(maps.Keys(layers), ";")})
	if withDependencies {
		analyseResult, err := r.dependencyLayersAnalyser.Analyse()
		if err != nil {
			return NewCommandRunExceptionAnalyserException(err)
		}
		analysisResult := output_result.NewOutputResultFromAnalysisResult(analyseResult)
		layersDependOnLayers := r.calculateLayerDependencies(analysisResult.AllRules())
		layerDependencies := make(map[string]string)
		for _, layer := range layers {
			for key, value := range layersDependOnLayers[layer] {
				layerDependencies[key] = value
			}
		}
		size := 0
		for size != len(layerDependencies) {
			size = len(layerDependencies)
			layerDependenciesCopy := layerDependencies
			for _, layerDependency := range layerDependenciesCopy {
				layerDependencies[layerDependency] = layersDependOnLayers[layerDependency][layerDependency]
			}
		}
		output.WriteLineFormatted(output_formatter_contract2.StringOrArrayOfStrings{String: strings.Join(maps.Keys(layerDependencies), ";")})
	}
	return nil
}

func (r *ChangedFilesRunner) calculateLayerDependencies(rules []result_contract2.RuleInterface) map[string]map[string]string {
	layersDependOnLayers := make(map[string]map[string]string)
	for _, rule := range rules {
		if _, ok := rule.(result_contract2.CoveredRuleInterface); !ok {
			continue
		}
		rule := rule.(result_contract2.CoveredRuleInterface)
		layerA := rule.GetDependerLayer()
		layerB := rule.GetDependentLayer()
		if _, ok := layersDependOnLayers[layerB]; !ok {
			layersDependOnLayers[layerB] = make(map[string]string)
		}
		if _, ok := layersDependOnLayers[layerB][layerA]; !ok {
			layersDependOnLayers[layerB][layerA] = layerA
		}
	}
	return layersDependOnLayers
}

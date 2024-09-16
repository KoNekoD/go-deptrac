package command

import (
	"github.com/KoNekoD/go-deptrac/pkg/contract/output_formatter"
	"github.com/KoNekoD/go-deptrac/pkg/contract/result"
	"github.com/KoNekoD/go-deptrac/pkg/contract/result/output_result"
	"github.com/KoNekoD/go-deptrac/pkg/core/analyser"
	"golang.org/x/exp/maps"
	"strings"
)

// ChangedFilesRunner - Should only be used by ChangedFilesCommand
type ChangedFilesRunner struct {
	layerForTokenAnalyser    *analyser.LayerForTokenAnalyser
	dependencyLayersAnalyser *analyser.DependencyLayersAnalyser
}

func NewChangedFilesRunner(layerForTokenAnalyser *analyser.LayerForTokenAnalyser, dependencyLayersAnalyser *analyser.DependencyLayersAnalyser) *ChangedFilesRunner {
	return &ChangedFilesRunner{
		layerForTokenAnalyser:    layerForTokenAnalyser,
		dependencyLayersAnalyser: dependencyLayersAnalyser,
	}
}

func (r *ChangedFilesRunner) Run(files []string, withDependencies bool, output output_formatter.OutputInterface) error {
	layers := make(map[string]string)
	for _, file := range files {
		matches, err := r.layerForTokenAnalyser.FindLayerForToken(file, analyser.TokenTypeFile)
		if err != nil {
			return NewCommandRunExceptionAnalyserException(err)
		}
		for _, match := range matches {
			for _, layer := range match {
				layers[layer] = layer
			}
		}
	}
	output.WriteLineFormatted(output_formatter.StringOrArrayOfStrings{String: strings.Join(maps.Keys(layers), ";")})
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
		output.WriteLineFormatted(output_formatter.StringOrArrayOfStrings{String: strings.Join(maps.Keys(layerDependencies), ";")})
	}
	return nil
}

func (r *ChangedFilesRunner) calculateLayerDependencies(rules []result.RuleInterface) map[string]map[string]string {
	layersDependOnLayers := make(map[string]map[string]string)
	for _, rule := range rules {
		if _, ok := rule.(result.CoveredRuleInterface); !ok {
			continue
		}
		rule := rule.(result.CoveredRuleInterface)
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

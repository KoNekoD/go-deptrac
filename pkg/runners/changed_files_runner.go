package runners

import (
	"github.com/KoNekoD/go-deptrac/pkg/dependencies"
	"github.com/KoNekoD/go-deptrac/pkg/domain/apperrors"
	"github.com/KoNekoD/go-deptrac/pkg/domain/enums"
	"github.com/KoNekoD/go-deptrac/pkg/results"
	"github.com/KoNekoD/go-deptrac/pkg/rules"
	"github.com/KoNekoD/go-deptrac/pkg/tokens"
	"golang.org/x/exp/maps"
	"strings"
)

// ChangedFilesRunner - Should only be used by ChangedFilesCommand
type ChangedFilesRunner struct {
	layerForTokenAnalyser    *tokens.LayerForTokenAnalyser
	dependencyLayersAnalyser *dependencies.DependencyLayersAnalyser
}

func NewChangedFilesRunner(layerForTokenAnalyser *tokens.LayerForTokenAnalyser, dependencyLayersAnalyser *dependencies.DependencyLayersAnalyser) *ChangedFilesRunner {
	return &ChangedFilesRunner{
		layerForTokenAnalyser:    layerForTokenAnalyser,
		dependencyLayersAnalyser: dependencyLayersAnalyser,
	}
}

func (r *ChangedFilesRunner) Run(files []string, withDependencies bool, output results.OutputInterface) error {
	layers := make(map[string]string)
	for _, file := range files {
		matches, err := r.layerForTokenAnalyser.FindLayerForToken(file, enums.TokenTypeFile)
		if err != nil {
			return apperrors.NewCommandRunExceptionAnalyserException(err)
		}
		for _, match := range matches {
			for _, layer := range match {
				layers[layer] = layer
			}
		}
	}
	output.WriteLineFormatted(results.StringOrArrayOfStrings{String: strings.Join(maps.Keys(layers), ";")})
	if withDependencies {
		analyseResult, err := r.dependencyLayersAnalyser.Analyse()
		if err != nil {
			return apperrors.NewCommandRunExceptionAnalyserException(err)
		}
		analysisResult := results.NewOutputResultFromAnalysisResult(analyseResult)
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
		output.WriteLineFormatted(results.StringOrArrayOfStrings{String: strings.Join(maps.Keys(layerDependencies), ";")})
	}
	return nil
}

func (r *ChangedFilesRunner) calculateLayerDependencies(rulesList []rules.RuleInterface) map[string]map[string]string {
	layersDependOnLayers := make(map[string]map[string]string)
	for _, rule := range rulesList {
		if _, ok := rule.(rules.CoveredRuleInterface); !ok {
			continue
		}
		rule := rule.(rules.CoveredRuleInterface)
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

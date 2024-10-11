package runners

import (
	services2 "github.com/KoNekoD/go-deptrac/pkg/application/services"
	"github.com/KoNekoD/go-deptrac/pkg/application/services/analysers"
	"github.com/KoNekoD/go-deptrac/pkg/domain/apperrors"
	"github.com/KoNekoD/go-deptrac/pkg/domain/dtos/results"
	"github.com/KoNekoD/go-deptrac/pkg/domain/dtos/results/violations_rules"
	"github.com/KoNekoD/go-deptrac/pkg/domain/enums"
	"golang.org/x/exp/maps"
	"strings"
)

// ChangedFilesRunner - Should only be used by ChangedFilesCommand
type ChangedFilesRunner struct {
	layerForTokenAnalyser    *analysers.LayerForTokenAnalyser
	dependencyLayersAnalyser *analysers.DependencyLayersAnalyser
}

func NewChangedFilesRunner(layerForTokenAnalyser *analysers.LayerForTokenAnalyser, dependencyLayersAnalyser *analysers.DependencyLayersAnalyser) *ChangedFilesRunner {
	return &ChangedFilesRunner{
		layerForTokenAnalyser:    layerForTokenAnalyser,
		dependencyLayersAnalyser: dependencyLayersAnalyser,
	}
}

func (r *ChangedFilesRunner) Run(files []string, withDependencies bool, output services2.OutputInterface) error {
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
	output.WriteLineFormatted(services2.StringOrArrayOfStrings{String: strings.Join(maps.Keys(layers), ";")})
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
		output.WriteLineFormatted(services2.StringOrArrayOfStrings{String: strings.Join(maps.Keys(layerDependencies), ";")})
	}
	return nil
}

func (r *ChangedFilesRunner) calculateLayerDependencies(rulesList []violations_rules.RuleInterface) map[string]map[string]string {
	layersDependOnLayers := make(map[string]map[string]string)
	for _, rule := range rulesList {
		if _, ok := rule.(violations_rules.CoveredRuleInterface); !ok {
			continue
		}
		rule := rule.(violations_rules.CoveredRuleInterface)
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

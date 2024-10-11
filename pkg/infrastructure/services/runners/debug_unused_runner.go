package runners

import (
	"fmt"
	"github.com/KoNekoD/go-deptrac/pkg/application/services/analysers"
	"github.com/KoNekoD/go-deptrac/pkg/domain/apperrors"
	"github.com/KoNekoD/go-deptrac/pkg/infrastructure/services"
)

type DebugUnusedRunner struct {
	analyser *analysers.RulesetUsageAnalyser
}

func NewDebugUnusedRunner(analyser *analysers.RulesetUsageAnalyser) *DebugUnusedRunner {
	return &DebugUnusedRunner{analyser: analyser}
}

func (r *DebugUnusedRunner) Run(output services.OutputInterface, limit int) error {
	rulesetUsages, err := r.analyser.Analyse()
	if err != nil {
		return apperrors.NewCommandRunExceptionAnalyserException(err)
	}

	outputTable := r.prepareOutputTable(rulesetUsages, limit)
	output.GetStyle().Table([]string{"Unused"}, outputTable)
	return nil
}

func (r *DebugUnusedRunner) prepareOutputTable(layerNames map[string]map[string]int, limit int) [][]string {
	rows := make([][]string, 0)
	for dependerLayerName, dependentLayerNames := range layerNames {
		for dependentLayerName, numberOfDependencies := range dependentLayerNames {
			if numberOfDependencies <= limit {
				if numberOfDependencies == 0 {
					rows = append(rows, []string{fmt.Sprintf("<info>%s</> layer is not dependent on <info>%s</>", dependerLayerName, dependentLayerName)})
				} else {
					rows = append(rows, []string{fmt.Sprintf("<info>%s</> layer is dependent <info>%s</> layer %d times", dependerLayerName, dependentLayerName, numberOfDependencies)})
				}
			}
		}
	}
	return rows
}

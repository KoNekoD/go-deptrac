package runners

import (
	"fmt"
	"github.com/KoNekoD/go-deptrac/pkg/application/services/analysers"
	"github.com/KoNekoD/go-deptrac/pkg/domain/apperrors"
	"github.com/KoNekoD/go-deptrac/pkg/domain/dtos/results/violations_rules"
	"github.com/KoNekoD/go-deptrac/pkg/domain/utils"
	"github.com/KoNekoD/go-deptrac/pkg/infrastructure/services"
	"github.com/gookit/color"
)

type DebugDependenciesRunner struct {
	analyser *analysers.LayerDependenciesAnalyser
}

func NewDebugDependenciesRunner(analyser *analysers.LayerDependenciesAnalyser) *DebugDependenciesRunner {
	return &DebugDependenciesRunner{analyser: analyser}
}

func (d *DebugDependenciesRunner) Run(output services.OutputStyleInterface, layer string, target *string) error {
	uncoveredMap, err := d.analyser.GetDependencies(layer, target)
	if err != nil {
		return apperrors.NewCommandRunExceptionAnalyserException(err)
	}

	for targetLayer, violations := range uncoveredMap {
		output.Table([]string{targetLayer}, utils.Map(violations, d.FormatRow))
	}

	return nil
}

func (d *DebugDependenciesRunner) FormatRow(rule *violations_rules.Uncovered) []string {
	dependency := rule.GetDependency()
	message := color.Sprintf("<info>%s</> depends on <info>%s</> (%s)", dependency.GetDepender().ToString(), dependency.GetDependent().ToString(), rule.Layer)
	fileOccurrence := dependency.GetContext().FileOccurrence
	message += fmt.Sprintf("\n%s:%d", fileOccurrence.FilePath, fileOccurrence.Line)
	return []string{message}
}

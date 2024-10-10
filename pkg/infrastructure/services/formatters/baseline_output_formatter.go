package formatters

import (
	"github.com/KoNekoD/go-deptrac/pkg/domain/dtos/results"
	"github.com/KoNekoD/go-deptrac/pkg/domain/enums"
	"github.com/KoNekoD/go-deptrac/pkg/infrastructure/services"
	"github.com/gookit/color"
	"golang.org/x/exp/maps"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
	"slices"
)

const BaselineOutputFormatterDefaultPath = "./deptrac.baseline.yaml"

type BaselineOutputFormatter struct{}

func NewBaselineOutputFormatter() *BaselineOutputFormatter {
	return &BaselineOutputFormatter{}
}

func (b *BaselineOutputFormatter) Finish(outputResult *results.OutputResult, output services.OutputInterface, outputFormatterInput *OutputFormatterInput) error {
	groupedViolations := b.collectViolations(outputResult)

	for _, violations := range groupedViolations {
		slices.Sort(violations)
	}

	baselineFile := BaselineOutputFormatterDefaultPath
	if outputFormatterInput.OutputPath != nil {
		baselineFile = *outputFormatterInput.OutputPath
	}

	dirname := filepath.Dir(baselineFile)
	if stat, _ := os.Stat(dirname); stat == nil || !stat.IsDir() {
		if err := os.MkdirAll(dirname, 0777); err != nil {
			if stat2, _ := os.Stat(dirname); stat2 == nil || !stat2.IsDir() {
				output.WriteLineFormatted(services.StringOrArrayOfStrings{String: color.Sprintf("<error>Unable to create %s</>", dirname)})
				return err
			}
		}
	}

	marshalled, err := yaml.Marshal(map[string]interface{}{"deptrac": map[string]interface{}{"skip_violations": groupedViolations}})
	if err != nil {
		return err
	}

	err = os.WriteFile(baselineFile, marshalled, 0666)
	if err != nil {
		return err
	}

	output.WriteLineFormatted(services.StringOrArrayOfStrings{String: color.Sprintf("<info>Baseline dumped to %s</>", baselineFile)})

	return nil
}

func (b *BaselineOutputFormatter) collectViolations(outputResult *results.OutputResult) map[string][]string {
	violations := make(map[string]map[string]string)
	for _, rule := range append(outputResult.AllOf(enums.TypeViolation), outputResult.AllOf(enums.TypeSkippedViolation)...) {
		dependency := rule.GetDependency()
		dependerClass := dependency.GetDepender().ToString()
		dependentClass := dependency.GetDependent().ToString()
		if _, ok := violations[dependerClass]; !ok {
			violations[dependerClass] = make(map[string]string)
		}
		violations[dependerClass][dependentClass] = dependentClass
	}

	mapped := make(map[string][]string)
	for dependerClass, dependentClasses := range violations {
		mapped[dependerClass] = maps.Values(dependentClasses)
	}

	return mapped
}

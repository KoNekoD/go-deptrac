package event_helper

import (
	"github.com/KoNekoD/go-deptrac/pkg/analyser_contract/analysis_result"
	"github.com/KoNekoD/go-deptrac/pkg/analyser_contract/process_event"
	"github.com/KoNekoD/go-deptrac/pkg/analyser_contract/violation_creating_interface"
	"github.com/KoNekoD/go-deptrac/pkg/layer_contract"
	result_contract2 "github.com/KoNekoD/go-deptrac/pkg/result_contract"
	"github.com/KoNekoD/go-deptrac/pkg/util"
)

// EventHelper - Utility class for managing adding violations that could be skipped.
type EventHelper struct {
	UnmatchedSkippedViolation map[string][]string
	SkippedViolations         map[string][]string
	LayerProvider             *layer_contract.LayerProvider
}

func NewEventHelper(skippedViolations map[string][]string, layerProvider *layer_contract.LayerProvider) *EventHelper {
	return &EventHelper{
		UnmatchedSkippedViolation: skippedViolations,
		SkippedViolations:         skippedViolations,
		LayerProvider:             layerProvider,
	}
}

func (e *EventHelper) shouldViolationBeSkipped(depender string, dependent string) bool {
	skippedViolation, ok := e.SkippedViolations[depender]
	if !ok {
		return false
	}

	matched := len(skippedViolation) > 0 && util.InArray(dependent, skippedViolation)
	if !matched {
		return false
	}

	// remove unmatched if exists
	unmatchedSkippedViolationDeonder, ok := e.UnmatchedSkippedViolation[depender]
	if ok && util.InArray(dependent, unmatchedSkippedViolationDeonder) {
		UnmatchedSkippedViolationNew := make([]string, 0)
		for _, s := range e.UnmatchedSkippedViolation[depender] {
			if dependent != s {
				UnmatchedSkippedViolationNew = append(UnmatchedSkippedViolationNew, s)
			}
		}

		e.UnmatchedSkippedViolation[depender] = UnmatchedSkippedViolationNew
	}

	return true
}

func (e *EventHelper) UnmatchedSkippedViolations() map[string][]string {
	return e.UnmatchedSkippedViolation
}

func (e *EventHelper) AddSkippableViolation(event *process_event.ProcessEvent, analysisResult *analysis_result.AnalysisResult, dependentLayer string, violationCreatingRule violation_creating_interface.ViolationCreatingInterface) {
	if e.shouldViolationBeSkipped(event.Dependency.GetDepender().ToString(), event.Dependency.GetDependent().ToString()) {
		analysisResult.AddRule(result_contract2.NewSkippedViolation(event.Dependency, event.DependerLayer, dependentLayer))
	} else {
		analysisResult.AddRule(result_contract2.NewViolation(event.Dependency, event.DependerLayer, dependentLayer, violationCreatingRule))
	}
}

package event_helper

import (
	"github.com/KoNekoD/go-deptrac/pkg/contract/analyser/analysis_result"
	"github.com/KoNekoD/go-deptrac/pkg/contract/analyser/process_event"
	"github.com/KoNekoD/go-deptrac/pkg/contract/analyser/violation_creating_interface"
	"github.com/KoNekoD/go-deptrac/pkg/contract/layer"
	"github.com/KoNekoD/go-deptrac/pkg/contract/result"
	"github.com/KoNekoD/go-deptrac/pkg/util"
)

// EventHelper - Utility class for managing adding violations that could be skipped.
type EventHelper struct {
	UnmatchedSkippedViolation map[string][]string
	SkippedViolations         map[string][]string
	LayerProvider             *layer.LayerProvider
}

func NewEventHelper(skippedViolations map[string][]string, layerProvider *layer.LayerProvider) *EventHelper {
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
		analysisResult.AddRule(result.NewSkippedViolation(event.Dependency, event.DependerLayer, dependentLayer))
	} else {
		analysisResult.AddRule(result.NewViolation(event.Dependency, event.DependerLayer, dependentLayer, violationCreatingRule))
	}
}

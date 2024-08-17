package EventHelper

import (
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Analyser/AnalysisResult"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Analyser/ProcessEvent"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Analyser/ViolationCreatingInterface"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Layer/LayerProvider"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Result/SkippedViolation"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Result/Violation"
	"github.com/KoNekoD/go-deptrac/pkg/util"
)

// EventHelper - Utility class for managing adding violations that could be skipped.
type EventHelper struct {
	UnmatchedSkippedViolation map[string][]string
	SkippedViolations         map[string][]string
	LayerProvider             *LayerProvider.LayerProvider
}

func NewEventHelper(skippedViolations map[string][]string, layerProvider *LayerProvider.LayerProvider) *EventHelper {
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

func (e *EventHelper) AddSkippableViolation(event *ProcessEvent.ProcessEvent, analysisResult *AnalysisResult.AnalysisResult, dependentLayer string, violationCreatingRule ViolationCreatingInterface.ViolationCreatingInterface) {
	if e.shouldViolationBeSkipped(event.Dependency.GetDepender().ToString(), event.Dependency.GetDependent().ToString()) {
		analysisResult.AddRule(SkippedViolation.NewSkippedViolation(event.Dependency, event.DependerLayer, dependentLayer))
	} else {
		analysisResult.AddRule(Violation.NewViolation(event.Dependency, event.DependerLayer, dependentLayer, violationCreatingRule))
	}
}

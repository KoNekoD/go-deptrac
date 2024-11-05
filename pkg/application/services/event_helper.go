package services

import (
	"github.com/KoNekoD/go-deptrac/pkg/domain/dtos/results"
	"github.com/KoNekoD/go-deptrac/pkg/domain/dtos/results/violations_rules"
	"github.com/KoNekoD/go-deptrac/pkg/domain/events"
	"github.com/KoNekoD/go-deptrac/pkg/domain/utils"
)

// EventHelper - Utility class for managing adding violations that could be skipped.
type EventHelper struct {
	UnmatchedSkippedViolation map[string][]string
	SkippedViolations         map[string][]string
	LayerProvider             *LayerProvider
}

func NewEventHelper(skippedViolations map[string][]string, layerProvider *LayerProvider) *EventHelper {
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

	matched := len(skippedViolation) > 0 && utils.InArray(dependent, skippedViolation)
	if !matched {
		return false
	}

	// remove unmatched if exists
	unmatchedSkippedViolationDeonder, ok := e.UnmatchedSkippedViolation[depender]
	if ok && utils.InArray(dependent, unmatchedSkippedViolationDeonder) {
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

func (e *EventHelper) AddSkippableViolation(event *events.ProcessEvent, analysisResult *results.AnalysisResult, dependentLayer string, violationCreatingRule violations_rules.ViolationCreatingInterface) {
	if e.shouldViolationBeSkipped(event.Dependency.GetDepender().ToString(), event.Dependency.GetDependent().ToString()) {
		analysisResult.AddRule(violations_rules.NewSkippedViolation(event.Dependency, event.DependerLayer, dependentLayer))
	} else {
		analysisResult.AddRule(violations_rules.NewViolation(event.Dependency, event.DependerLayer, dependentLayer, violationCreatingRule))
	}
}

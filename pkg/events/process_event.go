package events

import (
	"github.com/KoNekoD/go-deptrac/pkg/domain/dtos/dependencies"
	"github.com/KoNekoD/go-deptrac/pkg/domain/dtos/tokens_references"
	"github.com/KoNekoD/go-deptrac/pkg/rules"
)

// ProcessEvent - Event that is triggered on every found dependency_contract. Used to apply rules on the found dependencies.
type ProcessEvent struct {
	Dependency         dependencies.DependencyInterface
	DependerReference  tokens_references.TokenReferenceInterface
	DependerLayer      string
	DependentReference tokens_references.TokenReferenceInterface
	DependentLayers    map[string]bool
	result             *rules.AnalysisResult
}

func NewProcessEvent(
	dependency dependencies.DependencyInterface,
	dependerReference tokens_references.TokenReferenceInterface,
	dependerLayer string,
	dependentReference tokens_references.TokenReferenceInterface,
	dependentLayers map[string]bool,
	result *rules.AnalysisResult,
) *ProcessEvent {
	return &ProcessEvent{
		Dependency:         dependency,
		DependerReference:  dependerReference,
		DependerLayer:      dependerLayer,
		DependentReference: dependentReference,
		DependentLayers:    dependentLayers,
		result:             result,
	}
}

func (e *ProcessEvent) GetResult() *rules.AnalysisResult {
	return e.result
}

func (e *ProcessEvent) ReplaceResult(result *rules.AnalysisResult) {
	e.result = result
}

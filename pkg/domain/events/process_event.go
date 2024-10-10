package events

import (
	"github.com/KoNekoD/go-deptrac/pkg/domain/dtos/analysis_results"
	"github.com/KoNekoD/go-deptrac/pkg/domain/dtos/dependencies"
	"github.com/KoNekoD/go-deptrac/pkg/domain/dtos/tokens_references"
)

// ProcessEvent - Event that is triggered on every found dependency_contract. Used to apply rules on the found dependencies.
type ProcessEvent struct {
	Dependency         dependencies.DependencyInterface
	DependerReference  tokens_references.TokenReferenceInterface
	DependerLayer      string
	DependentReference tokens_references.TokenReferenceInterface
	DependentLayers    map[string]bool
	result             *analysis_results.AnalysisResult
}

func NewProcessEvent(
	dependency dependencies.DependencyInterface,
	dependerReference tokens_references.TokenReferenceInterface,
	dependerLayer string,
	dependentReference tokens_references.TokenReferenceInterface,
	dependentLayers map[string]bool,
	result *analysis_results.AnalysisResult,
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

func (e *ProcessEvent) GetResult() *analysis_results.AnalysisResult {
	return e.result
}

func (e *ProcessEvent) ReplaceResult(result *analysis_results.AnalysisResult) {
	e.result = result
}

package events

import (
	"github.com/KoNekoD/go-deptrac/pkg/dependencies"
	"github.com/KoNekoD/go-deptrac/pkg/rules"
	"github.com/KoNekoD/go-deptrac/pkg/tokens"
)

// ProcessEvent - Event that is triggered on every found dependency_contract. Used to apply rules on the found dependencies.
type ProcessEvent struct {
	Dependency         dependencies.DependencyInterface
	DependerReference  tokens.TokenReferenceInterface
	DependerLayer      string
	DependentReference tokens.TokenReferenceInterface
	DependentLayers    map[string]bool
	result             *rules.AnalysisResult
}

func NewProcessEvent(
	dependency dependencies.DependencyInterface,
	dependerReference tokens.TokenReferenceInterface,
	dependerLayer string,
	dependentReference tokens.TokenReferenceInterface,
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

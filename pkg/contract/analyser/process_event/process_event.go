package process_event

import (
	"github.com/KoNekoD/go-deptrac/pkg/contract/analyser/analysis_result"
	"github.com/KoNekoD/go-deptrac/pkg/contract/ast"
	Dependency2 "github.com/KoNekoD/go-deptrac/pkg/contract/dependency"
)

// ProcessEvent - Event that is triggered on every found dependency. Used to apply rules on the found dependencies.
type ProcessEvent struct {
	Dependency         Dependency2.DependencyInterface
	DependerReference  ast.TokenReferenceInterface
	DependerLayer      string
	DependentReference ast.TokenReferenceInterface
	DependentLayers    map[string]bool
	result             *analysis_result.AnalysisResult
}

func NewProcessEvent(
	dependency Dependency2.DependencyInterface,
	dependerReference ast.TokenReferenceInterface,
	dependerLayer string,
	dependentReference ast.TokenReferenceInterface,
	dependentLayers map[string]bool,
	result *analysis_result.AnalysisResult,
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

func (e *ProcessEvent) GetResult() *analysis_result.AnalysisResult {
	return e.result
}

func (e *ProcessEvent) ReplaceResult(result *analysis_result.AnalysisResult) {
	e.result = result
}

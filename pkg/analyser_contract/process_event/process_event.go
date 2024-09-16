package process_event

import (
	"github.com/KoNekoD/go-deptrac/pkg/analyser_contract/analysis_result"
	"github.com/KoNekoD/go-deptrac/pkg/ast_contract"
	Dependency2 "github.com/KoNekoD/go-deptrac/pkg/dependency_contract"
)

// ProcessEvent - Event that is triggered on every found dependency_contract. Used to apply rules on the found dependencies.
type ProcessEvent struct {
	Dependency         Dependency2.DependencyInterface
	DependerReference  ast_contract.TokenReferenceInterface
	DependerLayer      string
	DependentReference ast_contract.TokenReferenceInterface
	DependentLayers    map[string]bool
	result             *analysis_result.AnalysisResult
}

func NewProcessEvent(
	dependency Dependency2.DependencyInterface,
	dependerReference ast_contract.TokenReferenceInterface,
	dependerLayer string,
	dependentReference ast_contract.TokenReferenceInterface,
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

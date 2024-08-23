package process_event

import (
	"fmt"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Ast/TokenReferenceInterface"
	Dependency2 "github.com/KoNekoD/go-deptrac/pkg/src/contract/Dependency/DependencyInterface"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/analyser/analysis_result"
)

// ProcessEvent - Event that is triggered on every found dependency. Used to apply rules on the found dependencies.
type ProcessEvent struct {
	Dependency         Dependency2.DependencyInterface
	DependerReference  TokenReferenceInterface.TokenReferenceInterface
	DependerLayer      string
	DependentReference TokenReferenceInterface.TokenReferenceInterface
	DependentLayers    map[string]bool
	result             *analysis_result.AnalysisResult
}

func NewProcessEvent(
	dependency Dependency2.DependencyInterface,
	dependerReference TokenReferenceInterface.TokenReferenceInterface,
	dependerLayer string,
	dependentReference TokenReferenceInterface.TokenReferenceInterface,
	dependentLayers map[string]bool,
	result *analysis_result.AnalysisResult,
) *ProcessEvent {
	if dependerLayer == "File" && dependentLayers["Ast"] {
		fmt.Println()
	}

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

package ProcessEvent

import (
	"fmt"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Analyser/AnalysisResult"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Ast/TokenReferenceInterface"
	Dependency2 "github.com/KoNekoD/go-deptrac/pkg/src/contract/Dependency/DependencyInterface"
)

// ProcessEvent - Event that is triggered on every found dependency. Used to apply rules on the found dependencies.
type ProcessEvent struct {
	Dependency         Dependency2.DependencyInterface
	DependerReference  TokenReferenceInterface.TokenReferenceInterface
	DependerLayer      string
	DependentReference TokenReferenceInterface.TokenReferenceInterface
	DependentLayers    map[string]bool
	result             *AnalysisResult.AnalysisResult
}

func NewProcessEvent(
	dependency Dependency2.DependencyInterface,
	dependerReference TokenReferenceInterface.TokenReferenceInterface,
	dependerLayer string,
	dependentReference TokenReferenceInterface.TokenReferenceInterface,
	dependentLayers map[string]bool,
	result *AnalysisResult.AnalysisResult,
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

func (e *ProcessEvent) GetResult() *AnalysisResult.AnalysisResult {
	return e.result
}

func (e *ProcessEvent) ReplaceResult(result *AnalysisResult.AnalysisResult) {
	e.result = result
}

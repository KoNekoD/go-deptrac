package analyser_core

import (
	"fmt"
	astContract "github.com/KoNekoD/go-deptrac/pkg/ast_contract"
	"github.com/KoNekoD/go-deptrac/pkg/ast_core"
	dependency_core2 "github.com/KoNekoD/go-deptrac/pkg/dependency_core"
	layer_contract2 "github.com/KoNekoD/go-deptrac/pkg/layer_contract"
)

type AnalyserException struct {
	Message  string
	Previous error
}

func (a *AnalyserException) Error() string {
	if a.Previous != nil {
		return fmt.Sprintf("%s\n%s", a.Message, a.Previous.Error())
	} else {
		return a.Message
	}
}

func NewInvalidEmitterConfiguration(e *dependency_core2.InvalidEmitterConfigurationException) *AnalyserException {
	return &AnalyserException{Message: "Invalid emitter configuration.", Previous: e}
}

func NewUnrecognizedToken(e *dependency_core2.UnrecognizedTokenException) *AnalyserException {
	return &AnalyserException{Message: "Unrecognized token.", Previous: e}
}

func NewInvalidLayerDefinition(e *layer_contract2.InvalidLayerDefinitionException) *AnalyserException {
	return &AnalyserException{Message: "Invalid layer_contract definition.", Previous: e}
}

func NewInvalidCollectorDefinition(e *layer_contract2.InvalidCollectorDefinitionException) *AnalyserException {
	return &AnalyserException{Message: "Invalid collector definition.", Previous: e}
}

func NewFailedAstParsing(e *ast_core.AstException) *AnalyserException {
	return &AnalyserException{Message: "Failed Ast parsing.", Previous: e}
}

func NewCouldNotParseFile(e *astContract.CouldNotParseFileException) *AnalyserException {
	return &AnalyserException{Message: "Could not parse file_supportive.", Previous: e}
}

func NewCircularReference(e *layer_contract2.CircularReferenceException) *AnalyserException {
	return &AnalyserException{Message: "Circular layer_contract dependency_contract.", Previous: e}
}

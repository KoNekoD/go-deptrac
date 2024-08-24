package analyser

import (
	"fmt"
	astContract "github.com/KoNekoD/go-deptrac/pkg/contract/ast"
	"github.com/KoNekoD/go-deptrac/pkg/contract/layer"
	"github.com/KoNekoD/go-deptrac/pkg/core/ast"
	"github.com/KoNekoD/go-deptrac/pkg/core/dependency"
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

func NewInvalidEmitterConfiguration(e *dependency.InvalidEmitterConfigurationException) *AnalyserException {
	return &AnalyserException{Message: "Invalid emitter configuration.", Previous: e}
}

func NewUnrecognizedToken(e *dependency.UnrecognizedTokenException) *AnalyserException {
	return &AnalyserException{Message: "Unrecognized token.", Previous: e}
}

func NewInvalidLayerDefinition(e *layer.InvalidLayerDefinitionException) *AnalyserException {
	return &AnalyserException{Message: "Invalid layer definition.", Previous: e}
}

func NewInvalidCollectorDefinition(e *layer.InvalidCollectorDefinitionException) *AnalyserException {
	return &AnalyserException{Message: "Invalid collector definition.", Previous: e}
}

func NewFailedAstParsing(e *ast.AstException) *AnalyserException {
	return &AnalyserException{Message: "Failed Ast parsing.", Previous: e}
}

func NewCouldNotParseFile(e *astContract.CouldNotParseFileException) *AnalyserException {
	return &AnalyserException{Message: "Could not parse file.", Previous: e}
}

func NewCircularReference(e *layer.CircularReferenceException) *AnalyserException {
	return &AnalyserException{Message: "Circular layer dependency.", Previous: e}
}

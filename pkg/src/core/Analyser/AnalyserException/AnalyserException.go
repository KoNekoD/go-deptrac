package AnalyserException

import (
	"fmt"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Ast/CouldNotParseFileException"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Layer/CircularReferenceException"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Layer/InvalidCollectorDefinitionException"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Layer/InvalidLayerDefinitionException"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/Ast/AstException"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/Dependency/InvalidEmitterConfigurationException"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/Dependency/UnrecognizedTokenException"
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

func NewInvalidEmitterConfiguration(e *InvalidEmitterConfigurationException.InvalidEmitterConfigurationException) *AnalyserException {
	return &AnalyserException{Message: "Invalid emitter configuration.", Previous: e}
}

func NewUnrecognizedToken(e *UnrecognizedTokenException.UnrecognizedTokenException) *AnalyserException {
	return &AnalyserException{Message: "Unrecognized token.", Previous: e}
}

func NewInvalidLayerDefinition(e *InvalidLayerDefinitionException.InvalidLayerDefinitionException) *AnalyserException {
	return &AnalyserException{Message: "Invalid layer definition.", Previous: e}
}

func NewInvalidCollectorDefinition(e *InvalidCollectorDefinitionException.InvalidCollectorDefinitionException) *AnalyserException {
	return &AnalyserException{Message: "Invalid collector definition.", Previous: e}
}

func NewFailedAstParsing(e *AstException.AstException) *AnalyserException {
	return &AnalyserException{Message: "Failed Ast parsing.", Previous: e}
}

func NewCouldNotParseFile(e *CouldNotParseFileException.CouldNotParseFileException) *AnalyserException {
	return &AnalyserException{Message: "Could not parse file.", Previous: e}
}

func NewCircularReference(e *CircularReferenceException.CircularReferenceException) *AnalyserException {
	return &AnalyserException{Message: "Circular layer dependency.", Previous: e}
}

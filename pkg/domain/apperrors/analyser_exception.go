package apperrors

import (
	"fmt"
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

func NewInvalidEmitterConfiguration(e *InvalidEmitterConfigurationException) *AnalyserException {
	return &AnalyserException{Message: "Invalid emitter configuration.", Previous: e}
}

func NewUnrecognizedToken(e *UnrecognizedTokenException) *AnalyserException {
	return &AnalyserException{Message: "Unrecognized token.", Previous: e}
}

func NewInvalidLayerDefinition(e *InvalidLayerDefinitionException) *AnalyserException {
	return &AnalyserException{Message: "Invalid layer_contract definition.", Previous: e}
}

func NewInvalidCollectorDefinition(e *InvalidCollectorDefinitionException) *AnalyserException {
	return &AnalyserException{Message: "Invalid collector definition.", Previous: e}
}

func NewFailedAstParsing(e *AstException) *AnalyserException {
	return &AnalyserException{Message: "Failed Ast parsing.", Previous: e}
}

func NewCouldNotParseFile(e *CouldNotParseFileException) *AnalyserException {
	return &AnalyserException{Message: "Could not parse file_supportive.", Previous: e}
}

func NewCircularReference(e *CircularReferenceException) *AnalyserException {
	return &AnalyserException{Message: "Circular layer_contract dependency_contract.", Previous: e}
}

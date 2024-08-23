package layer

import "fmt"

// InvalidLayerDefinitionException - Thrown when the configuration of a particular layer is not valid. Use this exception when writing custom collectors.
type InvalidLayerDefinitionException struct {
	Message string
}

func (e InvalidLayerDefinitionException) Error() string {
	return e.Message
}

func NewInvalidLayerDefinitionExceptionMissingName() *InvalidLayerDefinitionException {
	return &InvalidLayerDefinitionException{
		Message: "Could not resolve layer definition. The field \"name\" is required for all layers.",
	}
}

func NewInvalidLayerDefinitionExceptionDuplicateName(layerName string) *InvalidLayerDefinitionException {
	return &InvalidLayerDefinitionException{
		Message: fmt.Sprintf("The layer name \"%s\" is already in use. Names must be unique.", layerName),
	}
}

func NewInvalidLayerDefinitionExceptionCollectorRequired(layerName string) *InvalidLayerDefinitionException {
	return &InvalidLayerDefinitionException{
		Message: fmt.Sprintf("The layer \"%s\" is empty. You must assign at least 1 collector to a layer.", layerName),
	}
}

func NewInvalidLayerDefinitionExceptionLayerRequired() *InvalidLayerDefinitionException {
	return &InvalidLayerDefinitionException{
		Message: "Layer configuration is empty. You need to define at least 1 layer.",
	}
}

func NewInvalidLayerDefinitionExceptionCircularTokenReference(tokenName string) *InvalidLayerDefinitionException {
	return &InvalidLayerDefinitionException{
		Message: fmt.Sprintf("Circular dependency between layers detected. Token \"%s\" could not be resolved.", tokenName),
	}
}

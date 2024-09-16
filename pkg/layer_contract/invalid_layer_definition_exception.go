package layer_contract

import "fmt"

// InvalidLayerDefinitionException - Thrown when the configuration of a particular layer_contract is not valid. Use this exception when writing custom collectors.
type InvalidLayerDefinitionException struct {
	Message string
}

func (e InvalidLayerDefinitionException) Error() string {
	return e.Message
}

func NewInvalidLayerDefinitionExceptionMissingName() *InvalidLayerDefinitionException {
	return &InvalidLayerDefinitionException{
		Message: "Could not resolve layer_contract definition. The field \"name\" is required for all layers.",
	}
}

func NewInvalidLayerDefinitionExceptionDuplicateName(layerName string) *InvalidLayerDefinitionException {
	return &InvalidLayerDefinitionException{
		Message: fmt.Sprintf("The layer_contract name \"%s\" is already in use. Names must be unique.", layerName),
	}
}

func NewInvalidLayerDefinitionExceptionCollectorRequired(layerName string) *InvalidLayerDefinitionException {
	return &InvalidLayerDefinitionException{
		Message: fmt.Sprintf("The layer_contract \"%s\" is empty. You must assign at least 1 collector to a layer_contract.", layerName),
	}
}

func NewInvalidLayerDefinitionExceptionLayerRequired() *InvalidLayerDefinitionException {
	return &InvalidLayerDefinitionException{
		Message: "Layer configuration is empty. You need to define at least 1 layer_contract.",
	}
}

func NewInvalidLayerDefinitionExceptionCircularTokenReference(tokenName string) *InvalidLayerDefinitionException {
	return &InvalidLayerDefinitionException{
		Message: fmt.Sprintf("Circular dependency_contract between layers detected. Token \"%s\" could not be resolved.", tokenName),
	}
}

package apperrors

import (
	"fmt"
)

type InvalidEmitterConfigurationException struct {
	Message string
}

func (e InvalidEmitterConfigurationException) Error() string {
	return e.Message
}

func NewInvalidEmitterConfigurationException(message string) *InvalidEmitterConfigurationException {
	return &InvalidEmitterConfigurationException{Message: message}
}

func NewInvalidEmitterConfigurationExceptionCouldNotLocate(emitterType string) *InvalidEmitterConfigurationException {
	return NewInvalidEmitterConfigurationException(fmt.Sprintf("Could not locate emitter type '%s' in the DI container.", emitterType))
}

func NewInvalidEmitterConfigurationExceptionIsNotEmitter(emitterType string, emitter interface{}) *InvalidEmitterConfigurationException {
	return NewInvalidEmitterConfigurationException(fmt.Sprintf("Type \"%s\" is not valid emitter (expected \"%s\", but is \"%T\").", emitterType, "DependencyEmitterInterface", emitter))
}

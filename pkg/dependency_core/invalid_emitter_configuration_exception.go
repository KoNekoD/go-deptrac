package dependency_core

import (
	"fmt"
	"github.com/KoNekoD/go-deptrac/pkg/config_contract"
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

func NewInvalidEmitterConfigurationExceptionCouldNotLocate(emitterType config_contract.EmitterType) *InvalidEmitterConfigurationException {
	return NewInvalidEmitterConfigurationException(fmt.Sprintf("Could not locate emitter type '%s' in the DI container.", emitterType))
}

func NewInvalidEmitterConfigurationExceptionIsNotEmitter(emitterType config_contract.EmitterType, emitter interface{}) *InvalidEmitterConfigurationException {
	return NewInvalidEmitterConfigurationException(fmt.Sprintf("Type \"%s\" is not valid emitter (expected \"%s\", but is \"%T\").", emitterType, "DependencyEmitterInterface", emitter))
}

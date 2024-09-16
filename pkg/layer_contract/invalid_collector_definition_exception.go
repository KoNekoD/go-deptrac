package layer_contract

import (
	"fmt"
	"github.com/KoNekoD/go-deptrac/pkg/config_contract"
	"reflect"
	"strings"
)

// InvalidCollectorDefinitionException - Thrown when the configuration of a particular collector is not valid. Use this exception when writing custom collectors.
type InvalidCollectorDefinitionException struct {
	Message  string
	Previous error
}

func (e InvalidCollectorDefinitionException) Error() string {
	if e.Previous != nil {
		return fmt.Sprintf("%s\n%s", e.Message, e.Previous.Error())
	} else {
		return e.Message
	}
}

func NewInvalidCollectorDefinitionExceptionMissingType() *InvalidCollectorDefinitionException {
	return &InvalidCollectorDefinitionException{
		Message: "Could not resolve collector definition because of missing \"type\" field.",
	}
}

func NewInvalidCollectorDefinitionExceptionUnsupportedType(collectorType config_contract.CollectorType, supportedTypes []config_contract.CollectorType, previous error) *InvalidCollectorDefinitionException {
	supportedTypesStrings := make([]string, len(supportedTypes))

	for i, supportedType := range supportedTypes {
		supportedTypesStrings[i] = string(supportedType)
	}

	return &InvalidCollectorDefinitionException{
		Message:  fmt.Sprintf("Could not find a collector for type \"%s\". Supported types: \"%s\".", collectorType, strings.Join(supportedTypesStrings, "\", \"")),
		Previous: previous,
	}

}

func NewInvalidCollectorDefinitionExceptionUnsupportedClass(id string, collector interface{}) *InvalidCollectorDefinitionException {
	return &InvalidCollectorDefinitionException{
		Message: fmt.Sprintf("Type \"%s\" is not valid collector (expected \"%s\", but is \"%s\").", id, "Qossmic\\Deptrac\\Contract\\Layer\\CollectorInterface", reflect.TypeOf(collector)),
	}
}

func NewInvalidCollectorDefinitionExceptionInvalidCollectorConfiguration(message string) *InvalidCollectorDefinitionException {
	return &InvalidCollectorDefinitionException{
		Message: message,
	}
}

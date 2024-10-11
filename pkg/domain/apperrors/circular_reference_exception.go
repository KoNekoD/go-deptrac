package apperrors

import (
	"fmt"
	"strings"
)

// CircularReferenceException - when there are circular dependencies between layers.
type CircularReferenceException struct {
	Message string
}

func (c CircularReferenceException) Error() string {
	return c.Message
}

func NewCircularReferenceExceptionFromCircularLayerDependency(layer string, others []string) CircularReferenceException {
	return CircularReferenceException{Message: fmt.Sprintf("Circular ruleset dependency_contract for layer_contract %s depending on: %s", layer, strings.Join(others, "->"))}
}

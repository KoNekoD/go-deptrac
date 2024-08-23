package dependency

import (
	"fmt"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/ast"
)

type UnrecognizedTokenException struct {
	Message string
}

func NewUnrecognizedTokenException(message string) *UnrecognizedTokenException {
	return &UnrecognizedTokenException{Message: message}
}

func (e UnrecognizedTokenException) Error() string {
	return e.Message
}

func (e UnrecognizedTokenException) NewCannotCreateReference(token ast.TokenInterface) *UnrecognizedTokenException {
	return &UnrecognizedTokenException{Message: fmt.Sprintf("Cannot create TokenReference for token '%T'", token)}
}

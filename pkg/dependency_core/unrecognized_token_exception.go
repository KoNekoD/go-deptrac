package dependency_core

import (
	"fmt"
	"github.com/KoNekoD/go-deptrac/pkg/ast_contract"
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

func (e UnrecognizedTokenException) NewCannotCreateReference(token ast_contract.TokenInterface) *UnrecognizedTokenException {
	return &UnrecognizedTokenException{Message: fmt.Sprintf("Cannot create TokenReference for token '%T'", token)}
}

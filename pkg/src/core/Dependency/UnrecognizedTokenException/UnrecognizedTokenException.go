package UnrecognizedTokenException

import (
	"fmt"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Ast/TokenInterface"
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

func (e UnrecognizedTokenException) NewCannotCreateReference(token TokenInterface.TokenInterface) *UnrecognizedTokenException {
	return &UnrecognizedTokenException{Message: fmt.Sprintf("Cannot create TokenReference for token '%T'", token)}
}

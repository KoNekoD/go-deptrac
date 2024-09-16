package ast_core

import (
	"fmt"
	"github.com/KoNekoD/go-deptrac/pkg/input_collector_core"
)

type AstException struct {
	Message  string
	Previous error
}

func (a *AstException) Error() string {
	if a.Previous != nil {
		return fmt.Sprintf("%s\n%s", a.Message, a.Previous.Error())
	} else {
		return a.Message
	}
}

func NewCouldNotCollectFiles(exception *input_collector_core.InputException) *AstException {
	return &AstException{Message: "Could not create AstMap.", Previous: exception}
}

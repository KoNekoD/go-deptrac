package FunctionNameCollector

import (
	"fmt"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Ast/TokenReferenceInterface"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Layer/InvalidCollectorDefinitionException"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/Ast/AstMap"
	"github.com/KoNekoD/go-deptrac/pkg/util"
)

type FunctionNameCollector struct{}

func NewFunctionNameCollector() *FunctionNameCollector {
	return &FunctionNameCollector{}
}

func (c FunctionNameCollector) Satisfy(config map[string]interface{}, reference TokenReferenceInterface.TokenReferenceInterface) (bool, error) {
	if _, ok := reference.(*AstMap.FunctionReference); !ok {
		return false, nil
	}

	pattern, err := c.GetPattern(config)
	if err != nil {
		return false, err
	}

	tokenName := reference.GetToken().(*AstMap.FunctionToken)

	return tokenName.Match(pattern), nil
}

func (c FunctionNameCollector) GetPattern(config map[string]interface{}) (string, error) {
	if !util.MapKeyExists(config, "value") || !util.MapKeyIsString(config, "value") {
		return "", InvalidCollectorDefinitionException.NewInvalidCollectorDefinitionExceptionInvalidCollectorConfiguration("FunctionNameCollector needs the regex configuration.")
	}

	return fmt.Sprintf("/%s/i", config["value"].(string)), nil
}

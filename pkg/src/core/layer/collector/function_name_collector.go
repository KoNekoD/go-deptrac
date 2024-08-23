package collector

import (
	"fmt"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/ast"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/layer"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/ast/ast_map"
	"github.com/KoNekoD/go-deptrac/pkg/util"
)

type FunctionNameCollector struct{}

func NewFunctionNameCollector() *FunctionNameCollector {
	return &FunctionNameCollector{}
}

func (c FunctionNameCollector) Satisfy(config map[string]interface{}, reference ast.TokenReferenceInterface) (bool, error) {
	if _, ok := reference.(*ast_map.FunctionReference); !ok {
		return false, nil
	}

	pattern, err := c.GetPattern(config)
	if err != nil {
		return false, err
	}

	tokenName := reference.GetToken().(*ast_map.FunctionToken)

	return tokenName.Match(pattern), nil
}

func (c FunctionNameCollector) GetPattern(config map[string]interface{}) (string, error) {
	if !util.MapKeyExists(config, "value") || !util.MapKeyIsString(config, "value") {
		return "", layer.NewInvalidCollectorDefinitionExceptionInvalidCollectorConfiguration("FunctionNameCollector needs the regex configuration.")
	}

	return fmt.Sprintf("/%s/i", config["value"].(string)), nil
}

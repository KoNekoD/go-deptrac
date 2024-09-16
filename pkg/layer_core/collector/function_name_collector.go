package collector

import (
	"fmt"
	"github.com/KoNekoD/go-deptrac/pkg/ast_contract"
	ast_map2 "github.com/KoNekoD/go-deptrac/pkg/ast_core/ast_map"
	"github.com/KoNekoD/go-deptrac/pkg/layer_contract"
	"github.com/KoNekoD/go-deptrac/pkg/util"
)

type FunctionNameCollector struct{}

func NewFunctionNameCollector() *FunctionNameCollector {
	return &FunctionNameCollector{}
}

func (c FunctionNameCollector) Satisfy(config map[string]interface{}, reference ast_contract.TokenReferenceInterface) (bool, error) {
	if _, ok := reference.(*ast_map2.FunctionReference); !ok {
		return false, nil
	}

	pattern, err := c.GetPattern(config)
	if err != nil {
		return false, err
	}

	tokenName := reference.GetToken().(*ast_map2.FunctionToken)

	return tokenName.Match(pattern), nil
}

func (c FunctionNameCollector) GetPattern(config map[string]interface{}) (string, error) {
	if !util.MapKeyExists(config, "value") || !util.MapKeyIsString(config, "value") {
		return "", layer_contract.NewInvalidCollectorDefinitionExceptionInvalidCollectorConfiguration("FunctionNameCollector needs the regex configuration.")
	}

	return fmt.Sprintf("/%s/i", config["value"].(string)), nil
}

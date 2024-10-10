package collectors

import (
	"fmt"
	"github.com/KoNekoD/go-deptrac/pkg/domain/apperrors"
	"github.com/KoNekoD/go-deptrac/pkg/domain/utils"
	"github.com/KoNekoD/go-deptrac/pkg/references"
	"github.com/KoNekoD/go-deptrac/pkg/tokens"
)

type FunctionNameCollector struct{}

func NewFunctionNameCollector() *FunctionNameCollector {
	return &FunctionNameCollector{}
}

func (c FunctionNameCollector) Satisfy(config map[string]interface{}, reference tokens.TokenReferenceInterface) (bool, error) {
	if _, ok := reference.(*references.FunctionReference); !ok {
		return false, nil
	}

	pattern, err := c.GetPattern(config)
	if err != nil {
		return false, err
	}

	tokenName := reference.GetToken().(*tokens.FunctionToken)

	return tokenName.Match(pattern), nil
}

func (c FunctionNameCollector) GetPattern(config map[string]interface{}) (string, error) {
	if !utils.MapKeyExists(config, "value") || !utils.MapKeyIsString(config, "value") {
		return "", apperrors.NewInvalidCollectorDefinitionInvalidCollectorConfiguration("FunctionNameCollector needs the regex configuration.")
	}

	return fmt.Sprintf("/%s/i", config["value"].(string)), nil
}

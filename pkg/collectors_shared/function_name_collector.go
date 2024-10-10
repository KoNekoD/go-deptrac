package collectors_shared

import (
	"fmt"
	"github.com/KoNekoD/go-deptrac/pkg/domain/apperrors"
	"github.com/KoNekoD/go-deptrac/pkg/domain/dtos/tokens"
	tokens_references2 "github.com/KoNekoD/go-deptrac/pkg/domain/dtos/tokens_references"
	"github.com/KoNekoD/go-deptrac/pkg/domain/utils"
)

type FunctionNameCollector struct{}

func NewFunctionNameCollector() *FunctionNameCollector {
	return &FunctionNameCollector{}
}

func (c FunctionNameCollector) Satisfy(config map[string]interface{}, reference tokens_references2.TokenReferenceInterface) (bool, error) {
	if _, ok := reference.(*tokens_references2.FunctionReference); !ok {
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

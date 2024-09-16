package collector

import (
	"github.com/KoNekoD/go-deptrac/pkg/ast_contract"
	ast_map2 "github.com/KoNekoD/go-deptrac/pkg/ast_core/ast_map"
	"github.com/KoNekoD/go-deptrac/pkg/layer_contract"
	"github.com/KoNekoD/go-deptrac/pkg/util"
)

type ClassNameRegexCollector struct {
	*RegexCollector
}

func NewClassNameRegexCollector() *ClassNameRegexCollector {
	return &ClassNameRegexCollector{
		RegexCollector: NewRegexCollector(),
	}
}

func (c ClassNameRegexCollector) Satisfy(config map[string]interface{}, reference ast_contract.TokenReferenceInterface) (bool, error) {
	if _, ok := reference.(*ast_map2.ClassLikeReference); !ok {
		return false, nil
	}

	validatedPattern, err := c.GetValidatedPattern(config, c.GetPattern)
	if err != nil {
		return false, err
	}

	return reference.GetToken().(*ast_map2.ClassLikeToken).Match(validatedPattern), nil
}

func (c ClassNameRegexCollector) GetPattern(config map[string]interface{}) (string, error) {
	if !util.MapKeyExists(config, "value") || !util.MapKeyIsString(config, "value") {
		return "", layer_contract.NewInvalidCollectorDefinitionExceptionInvalidCollectorConfiguration("ClassNameRegexCollector needs the regex configuration.")
	}

	return config["value"].(string), nil
}

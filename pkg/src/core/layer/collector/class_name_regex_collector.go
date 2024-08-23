package collector

import (
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Ast/TokenReferenceInterface"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Layer/InvalidCollectorDefinitionException"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/ast/ast_map"
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

func (c ClassNameRegexCollector) Satisfy(config map[string]interface{}, reference TokenReferenceInterface.TokenReferenceInterface) (bool, error) {
	if _, ok := reference.(*ast_map.ClassLikeReference); !ok {
		return false, nil
	}

	validatedPattern, err := c.GetValidatedPattern(config, c.GetPattern)
	if err != nil {
		return false, err
	}

	return reference.GetToken().(*ast_map.ClassLikeToken).Match(validatedPattern), nil
}

func (c ClassNameRegexCollector) GetPattern(config map[string]interface{}) (string, error) {
	if !util.MapKeyExists(config, "value") || !util.MapKeyIsString(config, "value") {
		return "", InvalidCollectorDefinitionException.NewInvalidCollectorDefinitionExceptionInvalidCollectorConfiguration("ClassNameRegexCollector needs the regex configuration.")
	}

	return config["value"].(string), nil
}

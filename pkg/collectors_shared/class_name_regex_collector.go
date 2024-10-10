package collectors_shared

import (
	"github.com/KoNekoD/go-deptrac/pkg/domain/apperrors"
	"github.com/KoNekoD/go-deptrac/pkg/domain/utils"
	"github.com/KoNekoD/go-deptrac/pkg/references"
	"github.com/KoNekoD/go-deptrac/pkg/tokens"
)

type ClassNameRegexCollector struct {
	*RegexCollector
}

func NewClassNameRegexCollector() *ClassNameRegexCollector {
	return &ClassNameRegexCollector{
		RegexCollector: NewRegexCollector(),
	}
}

func (c ClassNameRegexCollector) Satisfy(config map[string]interface{}, reference tokens.TokenReferenceInterface) (bool, error) {
	if _, ok := reference.(*references.ClassLikeReference); !ok {
		return false, nil
	}

	validatedPattern, err := c.GetValidatedPattern(config, c.GetPattern)
	if err != nil {
		return false, err
	}

	return reference.GetToken().(*tokens.ClassLikeToken).Match(validatedPattern), nil
}

func (c ClassNameRegexCollector) GetPattern(config map[string]interface{}) (string, error) {
	if !utils.MapKeyExists(config, "value") || !utils.MapKeyIsString(config, "value") {
		return "", apperrors.NewInvalidCollectorDefinitionInvalidCollectorConfiguration("ClassNameRegexCollector needs the regex configuration.")
	}

	return config["value"].(string), nil
}

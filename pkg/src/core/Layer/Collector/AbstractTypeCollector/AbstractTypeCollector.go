package AbstractTypeCollector

import (
	"fmt"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Ast/TokenReferenceInterface"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Layer/InvalidCollectorDefinitionException"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/Ast/AstMap"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/Layer/Collector/RegexCollector"
)

type AbstractTypeCollector struct {
	*RegexCollector.RegexCollector
}

func NewAbstractTypeCollector() *AbstractTypeCollector {
	return &AbstractTypeCollector{
		RegexCollector: RegexCollector.NewRegexCollector(),
	}
}

func (c *AbstractTypeCollector) GetType() AstMap.ClassLikeType {
	panic("Not implemented")
}

func (c *AbstractTypeCollector) Satisfy(config map[string]interface{}, reference TokenReferenceInterface.TokenReferenceInterface) (bool, error) {
	v, ok := reference.(*AstMap.ClassLikeReference)
	if !ok {
		return false, nil
	}

	pattern, err := c.GetValidatedPattern(config, c.GetPattern)
	if err != nil {
		return false, err
	}

	isClassLike := AstMap.TypeClasslike == c.GetType()
	isSameType := *v.Type == c.GetType()

	return (isClassLike || isSameType) && v.GetToken().(*AstMap.ClassLikeToken).Match(pattern), nil
}

func (c *AbstractTypeCollector) GetPattern(config map[string]interface{}) (string, error) {
	if _, ok := config["value"]; !ok {
		if _, ok2 := config["value"].(string); !ok2 {
			return "", InvalidCollectorDefinitionException.NewInvalidCollectorDefinitionExceptionInvalidCollectorConfiguration(fmt.Sprintf("Collector \"%s\" needs the regex configuration", c.GetType().ToString()))
		}
	}
	return fmt.Sprintf("/%s/i", config["value"].(string)), nil
}

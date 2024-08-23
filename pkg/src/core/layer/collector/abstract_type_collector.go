package collector

import (
	"fmt"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/ast"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/layer"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/ast/ast_map"
)

type AbstractTypeCollector struct {
	*RegexCollector
}

func NewAbstractTypeCollector() *AbstractTypeCollector {
	return &AbstractTypeCollector{
		RegexCollector: NewRegexCollector(),
	}
}

func (c *AbstractTypeCollector) GetType() ast_map.ClassLikeType {
	panic("Not implemented")
}

func (c *AbstractTypeCollector) Satisfy(config map[string]interface{}, reference ast.TokenReferenceInterface) (bool, error) {
	v, ok := reference.(*ast_map.ClassLikeReference)
	if !ok {
		return false, nil
	}

	pattern, err := c.GetValidatedPattern(config, c.GetPattern)
	if err != nil {
		return false, err
	}

	isClassLike := ast_map.TypeClasslike == c.GetType()
	isSameType := *v.Type == c.GetType()

	return (isClassLike || isSameType) && v.GetToken().(*ast_map.ClassLikeToken).Match(pattern), nil
}

func (c *AbstractTypeCollector) GetPattern(config map[string]interface{}) (string, error) {
	if _, ok := config["value"]; !ok {
		if _, ok2 := config["value"].(string); !ok2 {
			return "", layer.NewInvalidCollectorDefinitionExceptionInvalidCollectorConfiguration(fmt.Sprintf("Collector \"%s\" needs the regex configuration", c.GetType().ToString()))
		}
	}
	return fmt.Sprintf("/%s/i", config["value"].(string)), nil
}

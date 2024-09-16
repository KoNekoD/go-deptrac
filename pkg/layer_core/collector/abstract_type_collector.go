package collector

import (
	"fmt"
	"github.com/KoNekoD/go-deptrac/pkg/ast_contract"
	ast_map2 "github.com/KoNekoD/go-deptrac/pkg/ast_core/ast_map"
	"github.com/KoNekoD/go-deptrac/pkg/layer_contract"
)

type AbstractTypeCollector struct {
	*RegexCollector
}

func NewAbstractTypeCollector() *AbstractTypeCollector {
	return &AbstractTypeCollector{
		RegexCollector: NewRegexCollector(),
	}
}

func (c *AbstractTypeCollector) GetType() ast_map2.ClassLikeType {
	panic("Not implemented")
}

func (c *AbstractTypeCollector) Satisfy(config map[string]interface{}, reference ast_contract.TokenReferenceInterface) (bool, error) {
	v, ok := reference.(*ast_map2.ClassLikeReference)
	if !ok {
		return false, nil
	}

	pattern, err := c.GetValidatedPattern(config, c.GetPattern)
	if err != nil {
		return false, err
	}

	isClassLike := ast_map2.TypeClasslike == c.GetType()
	isSameType := *v.Type == c.GetType()

	return (isClassLike || isSameType) && v.GetToken().(*ast_map2.ClassLikeToken).Match(pattern), nil
}

func (c *AbstractTypeCollector) GetPattern(config map[string]interface{}) (string, error) {
	if _, ok := config["value"]; !ok {
		if _, ok2 := config["value"].(string); !ok2 {
			return "", layer_contract.NewInvalidCollectorDefinitionExceptionInvalidCollectorConfiguration(fmt.Sprintf("Collector \"%s\" needs the regex configuration", c.GetType().ToString()))
		}
	}
	return fmt.Sprintf("/%s/i", config["value"].(string)), nil
}

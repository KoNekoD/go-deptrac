package collectors

import (
	"fmt"
	"github.com/KoNekoD/go-deptrac/pkg/domain/apperrors"
	"github.com/KoNekoD/go-deptrac/pkg/references"
	"github.com/KoNekoD/go-deptrac/pkg/tokens"
	"github.com/KoNekoD/go-deptrac/pkg/types"
)

type AbstractTypeCollector struct {
	*RegexCollector
}

func NewAbstractTypeCollector() *AbstractTypeCollector {
	return &AbstractTypeCollector{
		RegexCollector: NewRegexCollector(),
	}
}

func (c *AbstractTypeCollector) GetType() types.ClassLikeType {
	panic("Not implemented")
}

func (c *AbstractTypeCollector) Satisfy(config map[string]interface{}, reference tokens.TokenReferenceInterface) (bool, error) {
	v, ok := reference.(*references.ClassLikeReference)
	if !ok {
		return false, nil
	}

	pattern, err := c.GetValidatedPattern(config, c.GetPattern)
	if err != nil {
		return false, err
	}

	isClassLike := types.TypeClasslike == c.GetType()
	isSameType := *v.Type == c.GetType()

	return (isClassLike || isSameType) && v.GetToken().(*tokens.ClassLikeToken).Match(pattern), nil
}

func (c *AbstractTypeCollector) GetPattern(config map[string]interface{}) (string, error) {
	if _, ok := config["value"]; !ok {
		if _, ok2 := config["value"].(string); !ok2 {
			return "", apperrors.NewInvalidCollectorDefinitionInvalidCollectorConfiguration(fmt.Sprintf("Collector \"%s\" needs the regex configuration", c.GetType().ToString()))
		}
	}
	return fmt.Sprintf("/%s/i", config["value"].(string)), nil
}

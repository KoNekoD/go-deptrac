package collectors_shared

import (
	"fmt"
	"github.com/KoNekoD/go-deptrac/pkg/domain/apperrors"
	tokens2 "github.com/KoNekoD/go-deptrac/pkg/domain/dtos/tokens"
	tokens_references2 "github.com/KoNekoD/go-deptrac/pkg/domain/dtos/tokens_references"
	"github.com/KoNekoD/go-deptrac/pkg/domain/enums"
)

type AbstractTypeCollector struct {
	*RegexCollector
}

func NewAbstractTypeCollector() *AbstractTypeCollector {
	return &AbstractTypeCollector{
		RegexCollector: NewRegexCollector(),
	}
}

func (c *AbstractTypeCollector) GetType() enums.ClassLikeType {
	panic("Not implemented")
}

func (c *AbstractTypeCollector) Satisfy(config map[string]interface{}, reference tokens_references2.TokenReferenceInterface) (bool, error) {
	v, ok := reference.(*tokens_references2.ClassLikeReference)
	if !ok {
		return false, nil
	}

	pattern, err := c.GetValidatedPattern(config, c.GetPattern)
	if err != nil {
		return false, err
	}

	isClassLike := enums.TypeClasslike == c.GetType()
	isSameType := *v.Type == c.GetType()

	return (isClassLike || isSameType) && v.GetToken().(*tokens2.ClassLikeToken).Match(pattern), nil
}

func (c *AbstractTypeCollector) GetPattern(config map[string]interface{}) (string, error) {
	if _, ok := config["value"]; !ok {
		if _, ok2 := config["value"].(string); !ok2 {
			return "", apperrors.NewInvalidCollectorDefinitionInvalidCollectorConfiguration(fmt.Sprintf("Collector \"%s\" needs the regex configuration", c.GetType().ToString()))
		}
	}
	return fmt.Sprintf("/%s/i", config["value"].(string)), nil
}

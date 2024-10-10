package collectors_shared

import (
	"fmt"
	"github.com/KoNekoD/go-deptrac/pkg/domain/apperrors"
	tokens_references2 "github.com/KoNekoD/go-deptrac/pkg/domain/dtos/tokens_references"
	"github.com/KoNekoD/go-deptrac/pkg/domain/utils"
	"github.com/KoNekoD/go-deptrac/pkg/parsers"
)

type MethodCollector struct {
	*RegexCollector
	astParser *parsers.NikicPhpParser
}

func NewMethodCollector(astParser *parsers.NikicPhpParser) *MethodCollector {
	return &MethodCollector{
		RegexCollector: NewRegexCollector(),
		astParser:      astParser,
	}
}

func (c *MethodCollector) Satisfy(config map[string]interface{}, reference tokens_references2.TokenReferenceInterface) (bool, error) {
	if _, ok := reference.(*tokens_references2.ClassLikeReference); !ok {
		return false, nil
	}

	pattern, err := c.GetValidatedPattern(config, c.GetPattern)
	if err != nil {
		return false, err
	}

	classLike := c.astParser.GetNodeForClassLikeReference(reference.(*tokens_references2.ClassLikeReference))
	if classLike == nil {
		return false, nil
	}

	// TODO: Implement methods getting
	panic("TODO" + pattern)
	//        foreach ($classLike->getMethods() as $classMethod) {
	//            if (1 === \preg_match($pattern, (string) $classMethod->name)) {
	//                return \true;
	//            }
	//        }
	//        return \false;
}

func (c *MethodCollector) GetPattern(config map[string]interface{}) (string, error) {
	if !utils.MapKeyExists(config, "value") || !utils.MapKeyIsString(config, "value") {
		return "", apperrors.NewInvalidCollectorDefinitionInvalidCollectorConfiguration("MethodCollector needs the name configuration.")
	}
	return fmt.Sprintf("/%s/i", config["value"].(string)), nil
}

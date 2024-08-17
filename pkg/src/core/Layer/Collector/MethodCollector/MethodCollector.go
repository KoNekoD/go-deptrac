package MethodCollector

import (
	"fmt"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Ast/TokenReferenceInterface"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Layer/InvalidCollectorDefinitionException"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/Ast/AstMap"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/Ast/Parser/NikicPhpParser/NikicPhpParser"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/Layer/Collector/RegexCollector"
	"github.com/KoNekoD/go-deptrac/pkg/util"
)

type MethodCollector struct {
	*RegexCollector.RegexCollector
	astParser *NikicPhpParser.NikicPhpParser
}

func NewMethodCollector(astParser *NikicPhpParser.NikicPhpParser) *MethodCollector {
	return &MethodCollector{
		RegexCollector: RegexCollector.NewRegexCollector(),
		astParser:      astParser,
	}
}

func (c *MethodCollector) Satisfy(config map[string]interface{}, reference TokenReferenceInterface.TokenReferenceInterface) (bool, error) {
	if _, ok := reference.(*AstMap.ClassLikeReference); !ok {
		return false, nil
	}

	pattern, err := c.GetValidatedPattern(config, c.GetPattern)
	if err != nil {
		return false, err
	}

	classLike := c.astParser.GetNodeForClassLikeReference(reference.(*AstMap.ClassLikeReference))
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
	if !util.MapKeyExists(config, "value") || !util.MapKeyIsString(config, "value") {
		return "", InvalidCollectorDefinitionException.NewInvalidCollectorDefinitionExceptionInvalidCollectorConfiguration("MethodCollector needs the name configuration.")
	}
	return fmt.Sprintf("/%s/i", config["value"].(string)), nil
}

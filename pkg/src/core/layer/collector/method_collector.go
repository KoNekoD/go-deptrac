package collector

import (
	"fmt"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Layer/InvalidCollectorDefinitionException"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/ast"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/ast/ast_map"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/ast/parser/nikic_php_parser"
	"github.com/KoNekoD/go-deptrac/pkg/util"
)

type MethodCollector struct {
	*RegexCollector
	astParser *nikic_php_parser.NikicPhpParser
}

func NewMethodCollector(astParser *nikic_php_parser.NikicPhpParser) *MethodCollector {
	return &MethodCollector{
		RegexCollector: NewRegexCollector(),
		astParser:      astParser,
	}
}

func (c *MethodCollector) Satisfy(config map[string]interface{}, reference ast.TokenReferenceInterface) (bool, error) {
	if _, ok := reference.(*ast_map.ClassLikeReference); !ok {
		return false, nil
	}

	pattern, err := c.GetValidatedPattern(config, c.GetPattern)
	if err != nil {
		return false, err
	}

	classLike := c.astParser.GetNodeForClassLikeReference(reference.(*ast_map.ClassLikeReference))
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

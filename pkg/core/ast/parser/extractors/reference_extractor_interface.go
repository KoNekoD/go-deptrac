package extractors

import (
	"github.com/KoNekoD/go-deptrac/pkg/core/ast/ast_map"
	"github.com/KoNekoD/go-deptrac/pkg/core/ast/parser"
	"go/ast"
)

type ReferenceExtractorInterface interface {
	ProcessNode(node ast.Node, referenceBuilder ast_map.ReferenceBuilderInterface, typeScope *parser.TypeScope)
}

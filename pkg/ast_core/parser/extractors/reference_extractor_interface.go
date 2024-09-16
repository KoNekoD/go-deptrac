package extractors

import (
	"github.com/KoNekoD/go-deptrac/pkg/ast_core/ast_map"
	"github.com/KoNekoD/go-deptrac/pkg/ast_core/parser"
	"go/ast"
)

type ReferenceExtractorInterface interface {
	ProcessNode(node ast.Node, referenceBuilder ast_map.ReferenceBuilderInterface, typeScope *parser.TypeScope)
}

package ReferenceExtractorInterface

import (
	"github.com/KoNekoD/go-deptrac/pkg/src/core/Ast/AstMap"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/Ast/Parser/TypeScope"
	"go/ast"
)

type ReferenceExtractorInterface interface {
	ProcessNode(node ast.Node, referenceBuilder AstMap.ReferenceBuilderInterface, typeScope *TypeScope.TypeScope)
}

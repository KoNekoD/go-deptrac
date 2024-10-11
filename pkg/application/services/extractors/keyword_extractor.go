package extractors

import (
	"github.com/KoNekoD/go-deptrac/pkg/application/services/references_builders"
	"github.com/KoNekoD/go-deptrac/pkg/application/services/types"
	"github.com/KoNekoD/go-deptrac/pkg/domain/utils"
	"go/ast"
)

type KeywordExtractor struct {
	typeResolver *types.TypeResolver
}

func NewKeywordExtractor(typeResolver *types.TypeResolver) *KeywordExtractor {
	return &KeywordExtractor{typeResolver: typeResolver}
}

func (e *KeywordExtractor) ProcessNode(node ast.Node, referenceBuilder references_builders.ReferenceBuilderInterface, typeScope *types.TypeScope) {
	if assertTypedNode, ok := node.(*ast.TypeAssertExpr); ok {
		for _, classLikeName := range e.typeResolver.ResolvePHPParserTypes(typeScope, assertTypedNode.Type) {
			referenceBuilder.Instanceof(classLikeName, utils.GetLineByPosition(typeScope.FilePath, int(assertTypedNode.Type.Pos())))
		}
	}
}

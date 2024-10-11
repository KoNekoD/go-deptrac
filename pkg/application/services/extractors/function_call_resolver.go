package extractors

import (
	"github.com/KoNekoD/go-deptrac/pkg/application/services/references_builders"
	"github.com/KoNekoD/go-deptrac/pkg/application/services/types"
	"github.com/KoNekoD/go-deptrac/pkg/domain/utils"
	"go/ast"
)

type FunctionCallResolver struct {
	typeResolver *types.TypeResolver
}

func NewFunctionCallResolver(typeResolver *types.TypeResolver) *FunctionCallResolver {
	return &FunctionCallResolver{typeResolver: typeResolver}
}

func (e *FunctionCallResolver) ProcessNode(node ast.Node, referenceBuilder references_builders.ReferenceBuilderInterface, typeScope *types.TypeScope) {
	typedNode, ok := node.(*ast.CallExpr)
	if !ok {
		return
	}

	for _, classLikeName := range e.typeResolver.ResolvePHPParserTypes(typeScope, typedNode.Fun) {
		referenceBuilder.UnresolvedFunctionCall(classLikeName, utils.GetLineByPosition(typeScope.FilePath, int(typedNode.Fun.Pos())))
	}
}

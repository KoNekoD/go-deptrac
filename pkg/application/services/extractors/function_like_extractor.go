package extractors

import (
	"github.com/KoNekoD/go-deptrac/pkg/application/services/references_builders"
	"github.com/KoNekoD/go-deptrac/pkg/application/services/types"
	"github.com/KoNekoD/go-deptrac/pkg/domain/utils"
	"go/ast"
)

type FunctionLikeExtractor struct {
	typeResolver *types.TypeResolver
}

func NewFunctionLikeExtractor(typeResolver *types.TypeResolver) *FunctionLikeExtractor {
	return &FunctionLikeExtractor{typeResolver: typeResolver}
}

func (e *FunctionLikeExtractor) ProcessNode(node ast.Node, referenceBuilder references_builders.ReferenceBuilderInterface, typeScope *types.TypeScope) {
	typedNode, ok := node.(*ast.FuncType)
	if !ok {
		return
	}

	for _, field := range typedNode.Params.List {
		if nil == field.Type {
			continue
		}

		for _, classLikeName := range e.typeResolver.ResolvePHPParserTypes(typeScope, field.Type) {
			pos := int(field.Type.Pos())

			referenceBuilder.Parameter(classLikeName, utils.GetLineByPosition(typeScope.FilePath, pos))
		}
	}

	for _, returnType := range typedNode.Results.List {
		if nil == returnType.Type {
			continue
		}

		for _, classLikeName := range e.typeResolver.ResolvePHPParserTypes(typeScope, returnType.Type) {
			pos := int(returnType.Type.Pos())

			referenceBuilder.ReturnType(classLikeName, utils.GetLineByPosition(typeScope.FilePath, pos))
		}
	}
}

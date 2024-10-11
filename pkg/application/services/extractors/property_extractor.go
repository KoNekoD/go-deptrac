package extractors

import (
	"github.com/KoNekoD/go-deptrac/pkg/application/services/references_builders"
	"github.com/KoNekoD/go-deptrac/pkg/application/services/types"
	"github.com/KoNekoD/go-deptrac/pkg/domain/utils"
	"go/ast"
)

type PropertyExtractor struct {
	typeResolver *types.TypeResolver
}

func NewPropertyExtractor(typeResolver *types.TypeResolver) *PropertyExtractor {
	return &PropertyExtractor{typeResolver: typeResolver}
}

func (e *PropertyExtractor) ProcessNode(node ast.Node, referenceBuilder references_builders.ReferenceBuilderInterface, typeScope *types.TypeScope) {
	typedStructNode, ok := node.(*ast.StructType)
	if !ok {
		return
	}

	for _, field := range typedStructNode.Fields.List {
		if nil == field.Type {
			continue
		}

		for _, classLikeName := range e.typeResolver.ResolvePHPParserTypes(typeScope, field.Type) {
			referenceBuilder.Variable(classLikeName, utils.GetLineByPosition(typeScope.FilePath, int(field.Type.Pos())))
		}
	}
}

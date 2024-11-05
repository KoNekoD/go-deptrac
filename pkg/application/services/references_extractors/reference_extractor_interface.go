package references_extractors

import (
	"github.com/KoNekoD/go-deptrac/pkg/application/services/references_builders"
	"github.com/KoNekoD/go-deptrac/pkg/application/services/types"
	"go/ast"
)

// TODO: НУЖНО РЕАЛИЗОВАТЬ ВСЕ ЭКСТРАКТОРЫ И УБРАТЬ В НИХ ЧАСТЬ КОДА ИЗ ВИЗИЬОРА И ТАЙП РЕЗОЛВЕРА - ВЫЧЕСЛЕНИЕ ЗАВИСИМОСТЕЙ-- ДЕЛО ReferenceExtractorInterface!

type ReferenceExtractorInterface interface {
	ProcessNode(node ast.Node, referenceBuilder references_builders.ReferenceBuilderInterface, typeScope *types.TypeScope)
}

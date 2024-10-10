package extractors

import (
	"github.com/KoNekoD/go-deptrac/pkg/references_builders"
	"github.com/KoNekoD/go-deptrac/pkg/types"
	"go/ast"
)

// TODO: НУЖНО РЕАЛИЗОВАТЬ ВСЕ ЭКСТРАКТОРЫ И УБРАТЬ В НИХ ЧАСТЬ КОДА ИЗ ВИЗИЬОРА И ТАЙП РЕЗОЛВЕРА - ВЫЧЕСЛЕНИЕ ЗАВИСИМОСТЕЙ-- ДЕЛО ReferenceExtractorInterface!

type ReferenceExtractorInterface interface {
	ProcessNode(node ast.Node, referenceBuilder references_builders.ReferenceBuilderInterface, typeScope *types.TypeScope)
}

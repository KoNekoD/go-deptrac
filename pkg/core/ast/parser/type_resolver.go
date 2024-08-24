package parser

import (
	"github.com/KoNekoD/go-deptrac/pkg/core/ast/parser/nikic_php_parser/node_namer"
	"go/ast"
	"go/parser"
	"go/token"
	"golang.org/x/exp/maps"
	"strings"
)

type TypeResolver struct {
	nodeNamer *node_namer.NodeNamer
}

func NewTypeResolver(nodeNamer *node_namer.NodeNamer) *TypeResolver {
	return &TypeResolver{
		nodeNamer: nodeNamer,
	}
}

func (r *TypeResolver) ResolvePHPParserTypes(typeScope *TypeScope, nodes ...ast.Expr) []string {
	types := make([]string, 0)
	for _, node := range nodes {
		resolved := r.resolvePHPParserType(typeScope, node)

		for _, s := range resolved {
			if s == "" {
				panic("1")
			}
		}

		types = append(types, resolved...)
	}

	return types
}

func (r *TypeResolver) resolvePHPParserType(scope *TypeScope, node ast.Expr) []string {
	switch v := node.(type) {
	case *ast.StarExpr:
		return r.resolveStarExpr(scope, v)
	case *ast.Ident:
		return make([]string, 0)
	case *ast.SelectorExpr:
		return r.resolveSelectorExpr(scope, v)
	case *ast.MapType:
		return r.resolveMapType(scope, v)
	case *ast.ArrayType:
		return r.resolveArrayType(scope, v)
	case *ast.Ellipsis:
		return r.resolvePHPParserType(scope, v.Elt)
	case *ast.InterfaceType:
		return r.resolveInterfaceType(scope, v)
	case *ast.FuncType:
		return r.resolveFieldList(scope, v.Results, v.TypeParams, v.Results)
	default:
		panic("5")
	}
}

func (r *TypeResolver) resolveStarExpr(scope *TypeScope, expr *ast.StarExpr) []string {
	return r.resolvePHPParserType(scope, expr.X)
}

func (r *TypeResolver) resolveIdent(scope *TypeScope, v *ast.Ident) []string {
	resolved := scope.GetUse(v.Name)

	if resolved == nil {
		// TODO: Добавить проверку, что такой модуль есть в go.mod или давать ошибку
		return make([]string, 0) // Костыль
	}

	return []string{*resolved}
}

func (r *TypeResolver) resolveSelectorExpr(scope *TypeScope, v *ast.SelectorExpr) []string {
	if ident, ok := v.X.(*ast.Ident); ok {
		xResolved := r.resolveIdent(scope, ident)

		if len(xResolved) == 0 {
			// TODO: Добавить проверку на наличие в go.mod
			return make([]string, 0) // TODO: ВАЖНО! Сейчас ПОЛНОСТЬЮ исключены ВСЕ внешние пакеты, нужно это исправить
		}

		if len(xResolved) != 1 {
			panic("impossible") // TODO: Possible when by import declared another package name
		}

		selResolved := v.Sel.Name

		xResolvedZero := xResolved[0]

		pathValidate := strings.Replace(xResolvedZero, "github.com/KoNekoD/go-deptrac/", "/home/mizuki/Documents/dev/KoNekoD/go-deptrac/", 1)
		parseDir, err := parser.ParseDir(token.NewFileSet(), pathValidate, nil, 0)
		if len(maps.Keys(parseDir)) != 1 {
			// TODO: Add checking in go.mod, если там нет такого модуля - ошибка
			return make([]string, 0) // Костыль
		}
		if err != nil {
			panic(err)
		}
		foundFileName := ""
		firstKey := maps.Keys(parseDir)[0]
		for filename, file := range parseDir[firstKey].Files {
			if file.Scope.Lookup(selResolved) != nil {
				foundFileName = filename
				break
			}
		}

		if foundFileName == "" {
			r.resolveIdent(scope, ident)
			panic("2")
		}

		// Validate
		pkgstrctnme, err := r.nodeNamer.GetPackageStructName(foundFileName, selResolved)

		if err != nil {
			panic(err)
		}

		return []string{pkgstrctnme} // TODO: Rework to Namer
	}

	return r.resolvePHPParserType(scope, v.X)
}

func (r *TypeResolver) resolveMapType(scope *TypeScope, v *ast.MapType) []string {
	return r.ResolvePHPParserTypes(scope, v.Key, v.Value)
}

func (r *TypeResolver) resolveArrayType(scope *TypeScope, v *ast.ArrayType) []string {
	if v.Len == nil {
		return r.resolvePHPParserType(scope, v.Elt)
	}

	return r.ResolvePHPParserTypes(scope, v.Len, v.Elt)
}

func (r *TypeResolver) resolveInterfaceType(scope *TypeScope, v *ast.InterfaceType) []string {
	if v.Methods.List == nil {
		return make([]string, 0)
	}
	return r.resolveFields(scope, v.Methods.List...)
}

func (r *TypeResolver) resolveFieldList(scope *TypeScope, list ...*ast.FieldList) []string {
	resolved := make([]string, 0)

	for _, fieldList := range list {
		if fieldList == nil {
			continue
		}
		resolved = append(resolved, r.resolveFields(scope, fieldList.List...)...)
	}

	return resolved
}

func (r *TypeResolver) resolveFields(scope *TypeScope, fields ...*ast.Field) []string {
	resolved := make([]string, 0)
	for _, field := range fields {
		if field.Names == nil {
			resolved = append(resolved, r.resolvePHPParserType(scope, field.Type)...)
		} else {
			for _, name := range field.Names {
				resolvedUse := scope.GetUse(name.Name)

				if resolvedUse == nil {
					panic("4")
				}

				resolved = append(resolved, *resolvedUse)
			}
		}
	}

	return resolved
}

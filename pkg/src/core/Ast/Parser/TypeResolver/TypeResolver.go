package TypeResolver

import (
	"github.com/KoNekoD/go-deptrac/pkg/src/core/Ast/Parser/TypeScope"
	"go/ast"
)

type TypeResolver struct {
}

func NewTypeResolver() *TypeResolver {
	return &TypeResolver{}
}

func (r *TypeResolver) ResolvePHPParserTypes(typeScope *TypeScope.TypeScope, nodes ...ast.Expr) []string {
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

func (r *TypeResolver) resolvePHPParserType(scope *TypeScope.TypeScope, node ast.Expr) []string {
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

func (r *TypeResolver) resolveStarExpr(scope *TypeScope.TypeScope, expr *ast.StarExpr) []string {
	return r.resolvePHPParserType(scope, expr.X)
}

func (r *TypeResolver) resolveIdent(scope *TypeScope.TypeScope, v *ast.Ident) []string {
	resolved := scope.GetUse(v.Name)

	if resolved == nil {
		return make([]string, 0)
	}

	return []string{*resolved}
}

func (r *TypeResolver) resolveSelectorExpr(scope *TypeScope.TypeScope, v *ast.SelectorExpr) []string {
	if ident, ok := v.X.(*ast.Ident); ok {
		return r.resolveIdent(scope, ident)
	}

	return r.resolvePHPParserType(scope, v.X)
}

func (r *TypeResolver) resolveMapType(scope *TypeScope.TypeScope, v *ast.MapType) []string {
	return r.ResolvePHPParserTypes(scope, v.Key, v.Value)
}

func (r *TypeResolver) resolveArrayType(scope *TypeScope.TypeScope, v *ast.ArrayType) []string {
	if v.Len == nil {
		return r.resolvePHPParserType(scope, v.Elt)
	}

	return r.ResolvePHPParserTypes(scope, v.Len, v.Elt)
}

func (r *TypeResolver) resolveInterfaceType(scope *TypeScope.TypeScope, v *ast.InterfaceType) []string {
	if v.Methods.List == nil {
		return make([]string, 0)
	}
	return r.resolveFields(scope, v.Methods.List...)
}

func (r *TypeResolver) resolveFieldList(scope *TypeScope.TypeScope, list ...*ast.FieldList) []string {
	resolved := make([]string, 0)

	for _, fieldList := range list {
		if fieldList == nil {
			continue
		}
		resolved = append(resolved, r.resolveFields(scope, fieldList.List...)...)
	}

	return resolved
}

func (r *TypeResolver) resolveFields(scope *TypeScope.TypeScope, fields ...*ast.Field) []string {
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

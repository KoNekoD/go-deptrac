package utils

import "go/ast"

func IsStruct(object *ast.Object) bool {
	if object.Kind != ast.Typ {
		return false // Only types can be structs
	}
	decl, okDecl := object.Decl.(*ast.TypeSpec) // Is checked by object.Kind
	if !okDecl {
		return false
	}

	_, ok := decl.Type.(*ast.StructType) // Final check if type is struct
	return ok
}

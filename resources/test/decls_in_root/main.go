package decls_in_root

// ast.Decl - Tok - IMPORT (75)
import "fmt"

// TestConst - ast.Decl - Tok - CONST (64)
const TestConst = 1

// TestStruct - ast.Decl - Tok - TYPE (84)
type TestStruct struct{}

// testVar1 - ast.Decl - Tok - VAR (85)
var testVar1 = "1"

// Func1 - ast.FuncDecl - Recv is nil
func Func1() {
	fmt.Print()
}

// name - ast.FuncDecl - Recv relates to TestStruct
func (s TestStruct) name() {
	fmt.Println(testVar1, TestConst)
}

package reference_visitors

import (
	"fmt"
	"github.com/KoNekoD/go-deptrac/pkg/application/services/references_builders"
	"github.com/KoNekoD/go-deptrac/pkg/application/services/references_extractors"
	"github.com/KoNekoD/go-deptrac/pkg/application/services/types"
	_ "github.com/KoNekoD/go-deptrac/resources"
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func TestFileReferenceVisitorOk(t *testing.T) {
	file := "pkg/core/ast_contract/parser/nikic_php_parser/nikic_php_parser.go"

	nodes, err := parser.ParseFile(token.NewFileSet(), file, nil, 0)

	if err != nil {
		t.Error(err)
	}

	referenceExtractorInterfaces := make([]references_extractors.ReferenceExtractorInterface, 0)

	fileReferenceVisitor := NewFileReferenceVisitor(
		references_builders.CreateFileReferenceBuilder(file),
		types.NewTypeResolver(
			nil,
		),
		nil,
		referenceExtractorInterfaces...,
	)

	ast.Walk(fileReferenceVisitor, nodes)

	fmt.Println()
}

// TestDeclsInRoot - Decls is only ast.Decl or ast.FuncDecl
func TestDeclsInRoot(t *testing.T) {
	file, err := parser.ParseFile(token.NewFileSet(), "resources/test/decls_in_root/main.go", nil, 0)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(file)
}

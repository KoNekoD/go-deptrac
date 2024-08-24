package nikic_php_parser

import (
	"fmt"
	"github.com/KoNekoD/go-deptrac/pkg/core/ast/ast_map"
	parser2 "github.com/KoNekoD/go-deptrac/pkg/core/ast/parser"
	"github.com/KoNekoD/go-deptrac/pkg/core/ast/parser/extractors"
	_ "github.com/KoNekoD/go-deptrac/resources"
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func TestFileReferenceVisitorOk(t *testing.T) {
	file := "pkg/core/ast/parser/nikic_php_parser/nikic_php_parser.go"

	nodes, err := parser.ParseFile(token.NewFileSet(), file, nil, 0)

	if err != nil {
		t.Error(err)
	}

	referenceExtractorInterfaces := make([]extractors.ReferenceExtractorInterface, 0)

	fileReferenceVisitor := NewFileReferenceVisitor(
		ast_map.CreateFileReferenceBuilder(file),
		parser2.NewTypeResolver(
			nil,
		),
		nil,
		referenceExtractorInterfaces...,
	)

	ast.Walk(fileReferenceVisitor, nodes)

	fmt.Println()
}

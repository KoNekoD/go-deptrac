package nikic_php_parser

import (
	"github.com/KoNekoD/go-deptrac/pkg/core/ast/ast_map"
	_ "github.com/KoNekoD/go-deptrac/resources"
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func TestFileReferenceVisitorOk(t *testing.T) {
	file := "pkg/ast/parser.go"

	nodes, err := parser.ParseFile(token.NewFileSet(), file, nil, 0)

	if err != nil {
		t.Error(err)
	}

	fileReferenceVisitor := NewFileReferenceVisitor(
		ast_map.CreateFileReferenceBuilder(file),
		nil,
		nil,
	)

	ast.Walk(fileReferenceVisitor, nodes)
}

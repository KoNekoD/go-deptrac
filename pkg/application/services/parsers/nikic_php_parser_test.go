package parsers

import (
	"go/parser"
	"go/token"
	"testing"
)

func TestParsePackage(t *testing.T) {
	dir, err := parser.ParseDir(token.NewFileSet(), "/home/mizuki/Documents/dev/KoNekoD/go_deptrac/pkg/Core/Ast/AstMap", nil, 0)
	if err != nil {
		t.Error(err)
	}

	for _, v := range dir {
		t.Log(v)
	}
}

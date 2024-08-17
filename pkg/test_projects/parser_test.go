package test_projects

import "testing"

func TestOk(t *testing.T) {
	p := NewFileParser()

	file, err := p.ParseFile("/home/mizuki/Documents/dev/KoNekoD/go-deptrac/pkg/src/Core/Ast/AstLoader/AstLoader.go")
	if err != nil {
		t.Error(err)
	}

	t.Log(file)
}

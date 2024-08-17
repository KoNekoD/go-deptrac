package AstMapExtractor

import (
	"github.com/KoNekoD/go-deptrac/pkg/src/core/Ast/AstLoader"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/InputCollector/FileInputCollector"
	"github.com/KoNekoD/go-deptrac/pkg/test_projects"
	"os"
	"testing"
)

func TestAstMapExtractorExtractWorkedFine(t *testing.T) {
	paths := []string{
		"analyser",
		"ast",
		"config",
		"dumper",
		"layer",
		"output_formatter",
		"result",
		"util",
	}
	var excluded = make([]string, 0)

	wd, err := os.Getwd()

	if err != nil {
		t.Error(err)
	}

	basePath := wd + "/pkg/"

	fileInputCollector, err := FileInputCollector.NewFileInputCollector(paths, excluded, basePath)

	if err != nil {
		t.Error(err)
	}

	astLoader := AstLoader.NewAstLoader(test_projects.NewFileParser(), nil)

	e := NewAstMapExtractor(fileInputCollector, astLoader)

	extract, err := e.Extract()
	if err != nil {
		t.Error(err)
	}

	if extract == nil {
		t.Error("extract is nil")
	}
}

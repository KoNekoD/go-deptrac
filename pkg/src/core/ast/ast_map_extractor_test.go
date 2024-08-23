package ast

import (
	"github.com/KoNekoD/go-deptrac/pkg/src/core/ast/parser"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/ast/parser/cache"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/ast/parser/extractors"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/ast/parser/nikic_php_parser"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/input_collector"
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

	fileInputCollector, err := input_collector.NewFileInputCollector(paths, excluded, basePath)

	if err != nil {
		t.Error(err)
	}

	astLoader := NewAstLoader(
		nikic_php_parser.NewNikicPhpParser(
			cache.NewAstFileReferenceInMemoryCache(),
			parser.NewTypeResolver(),
			[]extractors.ReferenceExtractorInterface{},
		),
		nil,
	)

	e := NewAstMapExtractor(fileInputCollector, astLoader)

	extract, err := e.Extract()
	if err != nil {
		t.Error(err)
	}

	if extract == nil {
		t.Error("extract is nil")
	}
}

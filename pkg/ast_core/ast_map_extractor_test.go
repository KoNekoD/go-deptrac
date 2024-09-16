package ast_core

import (
	"github.com/KoNekoD/go-deptrac/pkg/ast_core/parser"
	"github.com/KoNekoD/go-deptrac/pkg/ast_core/parser/cache"
	"github.com/KoNekoD/go-deptrac/pkg/ast_core/parser/extractors"
	"github.com/KoNekoD/go-deptrac/pkg/ast_core/parser/nikic_php_parser"
	"github.com/KoNekoD/go-deptrac/pkg/input_collector_core"
	"os"
	"testing"
)

func TestAstMapExtractorExtractWorkedFine(t *testing.T) {
	paths := []string{
		"analyser_contract",
		"ast_contract",
		"config_contract",
		"dumper",
		"layer_contract",
		"output_formatter_contract",
		"result_contract",
		"util",
	}
	var excluded = make([]string, 0)

	wd, err := os.Getwd()

	if err != nil {
		t.Error(err)
	}

	basePath := wd + "/pkg/"

	fileInputCollector, err := input_collector_core.NewFileInputCollector(paths, excluded, basePath)

	if err != nil {
		t.Error(err)
	}

	astLoader := NewAstLoader(
		nikic_php_parser.NewNikicPhpParser(
			cache.NewAstFileReferenceInMemoryCache(),
			parser.NewTypeResolver(
				nil,
			),
			nil,
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

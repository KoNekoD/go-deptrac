package ast_map

import (
	"github.com/KoNekoD/go-deptrac/pkg/application/services/input_collectors"
	"github.com/KoNekoD/go-deptrac/pkg/extractors"
	"github.com/KoNekoD/go-deptrac/pkg/parsers"
	"github.com/KoNekoD/go-deptrac/pkg/types"
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

	fileInputCollector, err := input_collectors.NewFileInputCollector(paths, excluded, basePath)

	if err != nil {
		t.Error(err)
	}

	astLoader := NewAstLoader(
		parsers.NewNikicPhpParser(
			NewAstFileReferenceInMemoryCache(),
			types.NewTypeResolver(
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

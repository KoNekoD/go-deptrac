package ast_map

import (
	"github.com/KoNekoD/go-deptrac/pkg/application/services/input_collectors"
	"github.com/KoNekoD/go-deptrac/pkg/domain/dtos/ast_map"
)

type AstMapExtractor struct {
	inputCollector input_collectors.InputCollector
	astLoader      *AstLoader
	astMapCache    *ast_map.AstMap
}

func NewAstMapExtractor(inputCollector input_collectors.InputCollector, astLoader *AstLoader) *AstMapExtractor {
	return &AstMapExtractor{
		inputCollector: inputCollector,
		astLoader:      astLoader,
		astMapCache:    nil,
	}
}

func (e *AstMapExtractor) Extract() (*ast_map.AstMap, error) {
	if e.astMapCache == nil {
		collected, err := e.inputCollector.Collect()

		if err != nil {
			return nil, err
		}

		createdAstMap, errCreateAstMap := e.astLoader.CreateAstMap(collected)
		if errCreateAstMap != nil {
			return nil, errCreateAstMap
		}

		e.astMapCache = createdAstMap
	}

	return e.astMapCache, nil
}

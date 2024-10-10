package ast_map

import (
	"github.com/KoNekoD/go-deptrac/pkg/collectors_shared"
)

type AstMapExtractor struct {
	inputCollector collectors_shared.InputCollectorInterface
	astLoader      *AstLoader
	astMapCache    *AstMap
}

func NewAstMapExtractor(inputCollector collectors_shared.InputCollectorInterface, astLoader *AstLoader) *AstMapExtractor {
	return &AstMapExtractor{
		inputCollector: inputCollector,
		astLoader:      astLoader,
		astMapCache:    nil,
	}
}

func (e *AstMapExtractor) Extract() (*AstMap, error) {
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

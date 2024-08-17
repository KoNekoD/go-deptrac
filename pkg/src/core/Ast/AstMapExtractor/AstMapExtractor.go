package AstMapExtractor

import (
	"github.com/KoNekoD/go-deptrac/pkg/src/core/Ast/AstLoader"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/Ast/AstMap"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/InputCollector/InputCollectorInterface"
)

type AstMapExtractor struct {
	inputCollector InputCollectorInterface.InputCollectorInterface
	astLoader      *AstLoader.AstLoader
	astMapCache    *AstMap.AstMap
}

func NewAstMapExtractor(inputCollector InputCollectorInterface.InputCollectorInterface, astLoader *AstLoader.AstLoader) *AstMapExtractor {
	return &AstMapExtractor{
		inputCollector: inputCollector,
		astLoader:      astLoader,
		astMapCache:    nil,
	}
}

func (e *AstMapExtractor) Extract() (*AstMap.AstMap, error) {
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

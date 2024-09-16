package ast_core

import (
	ast_contract2 "github.com/KoNekoD/go-deptrac/pkg/ast_contract"
	ast_map2 "github.com/KoNekoD/go-deptrac/pkg/ast_core/ast_map"
	"github.com/KoNekoD/go-deptrac/pkg/ast_core/parser"
	"github.com/KoNekoD/go-deptrac/pkg/dependency_injection_supportive/event_dispatcher/event_dispatcher_interface"
)

type AstLoader struct {
	parser          parser.ParserInterface
	eventDispatcher event_dispatcher_interface.EventDispatcherInterface
}

func NewAstLoader(parser parser.ParserInterface, eventDispatcher event_dispatcher_interface.EventDispatcherInterface) *AstLoader {
	return &AstLoader{
		parser:          parser,
		eventDispatcher: eventDispatcher,
	}
}

func (l *AstLoader) CreateAstMap(files []string) (*ast_map2.AstMap, error) {
	references := make([]*ast_map2.FileReference, 0)

	err := l.eventDispatcher.DispatchEvent(ast_contract2.NewPreCreateAstMapEvent(len(files)))
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		reference, err := l.parser.ParseFile(file)

		if err != nil {
			err := l.eventDispatcher.DispatchEvent(ast_contract2.NewAstFileSyntaxErrorEvent(file, err.Error()))
			if err != nil {
				return nil, err
			}

			continue
		}

		references = append(references, reference)

		errDispatchAnalysed := l.eventDispatcher.DispatchEvent(ast_contract2.NewAstFileAnalysedEvent(file))
		if errDispatchAnalysed != nil {
			return nil, errDispatchAnalysed
		}
	}

	astMap := ast_map2.NewAstMap(references)

	errDispatchPostCreateMap := l.eventDispatcher.DispatchEvent(ast_contract2.NewPostCreateAstMapEvent())
	if errDispatchPostCreateMap != nil {
		return nil, errDispatchPostCreateMap
	}

	return astMap, nil
}

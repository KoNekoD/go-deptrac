package ast

import (
	"github.com/KoNekoD/go-deptrac/pkg/contract/ast"
	"github.com/KoNekoD/go-deptrac/pkg/core/ast/ast_map"
	"github.com/KoNekoD/go-deptrac/pkg/core/ast/parser"
	"github.com/KoNekoD/go-deptrac/pkg/supportive/dependency_injection/event_dispatcher/event_dispatcher_interface"
)

type AstLoader struct {
	parser          parser.ParserInterface
	eventDispatcher util.EventDispatcherInterface
}

func NewAstLoader(parser parser.ParserInterface, eventDispatcher util.EventDispatcherInterface) *AstLoader {
	return &AstLoader{
		parser:          parser,
		eventDispatcher: eventDispatcher,
	}
}

func (l *AstLoader) CreateAstMap(files []string) (*ast_map.AstMap, error) {
	references := make([]*ast_map.FileReference, 0)

	err := l.eventDispatcher.DispatchEvent(ast.NewPreCreateAstMapEvent(len(files)))
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		reference, err := l.parser.ParseFile(file)

		if err != nil {
			err := l.eventDispatcher.DispatchEvent(ast.NewAstFileSyntaxErrorEvent(file, err.Error()))
			if err != nil {
				return nil, err
			}

			continue
		}

		references = append(references, reference)

		errDispatchAnalysed := l.eventDispatcher.DispatchEvent(ast.NewAstFileAnalysedEvent(file))
		if errDispatchAnalysed != nil {
			return nil, errDispatchAnalysed
		}
	}

	astMap := ast_map.NewAstMap(references)

	errDispatchPostCreateMap := l.eventDispatcher.DispatchEvent(ast.NewPostCreateAstMapEvent())
	if errDispatchPostCreateMap != nil {
		return nil, errDispatchPostCreateMap
	}

	return astMap, nil
}

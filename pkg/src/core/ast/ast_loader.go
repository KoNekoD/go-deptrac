package ast

import (
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Ast/AstFileAnalysedEvent"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Ast/AstFileSyntaxErrorEvent"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Ast/PostCreateAstMapEvent"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Ast/PreCreateAstMapEvent"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/ast/ast_map"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/ast/parser"
	"github.com/KoNekoD/go-deptrac/pkg/src/supportive/DependencyInjection/EventDispatcher/EventDispatcherInterface"
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

	err := l.eventDispatcher.DispatchEvent(PreCreateAstMapEvent.NewPreCreateAstMapEvent(len(files)))
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		reference, err := l.parser.ParseFile(file)

		if err != nil {
			err := l.eventDispatcher.DispatchEvent(AstFileSyntaxErrorEvent.NewAstFileSyntaxErrorEvent(file, err.Error()))
			if err != nil {
				return nil, err
			}

			continue
		}

		references = append(references, reference)

		errDispatchAnalysed := l.eventDispatcher.DispatchEvent(AstFileAnalysedEvent.NewAstFileAnalysedEvent(file))
		if errDispatchAnalysed != nil {
			return nil, errDispatchAnalysed
		}
	}

	astMap := ast_map.NewAstMap(references)

	errDispatchPostCreateMap := l.eventDispatcher.DispatchEvent(PostCreateAstMapEvent.NewPostCreateAstMapEvent())
	if errDispatchPostCreateMap != nil {
		return nil, errDispatchPostCreateMap
	}

	return astMap, nil
}

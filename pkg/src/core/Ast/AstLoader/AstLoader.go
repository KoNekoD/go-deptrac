package AstLoader

import (
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Ast/AstFileAnalysedEvent"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Ast/AstFileSyntaxErrorEvent"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Ast/PostCreateAstMapEvent"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Ast/PreCreateAstMapEvent"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/Ast/AstMap"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/Ast/Parser/ParserInterface"
	"github.com/KoNekoD/go-deptrac/pkg/src/supportive/DependencyInjection/EventDispatcher/EventDispatcherInterface"
)

type AstLoader struct {
	parser          ParserInterface.ParserInterface
	eventDispatcher util.EventDispatcherInterface
}

func NewAstLoader(parser ParserInterface.ParserInterface, eventDispatcher util.EventDispatcherInterface) *AstLoader {
	return &AstLoader{
		parser:          parser,
		eventDispatcher: eventDispatcher,
	}
}

func (l *AstLoader) CreateAstMap(files []string) (*AstMap.AstMap, error) {
	references := make([]*AstMap.FileReference, 0)

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

	astMap := AstMap.NewAstMap(references)

	errDispatchPostCreateMap := l.eventDispatcher.DispatchEvent(PostCreateAstMapEvent.NewPostCreateAstMapEvent())
	if errDispatchPostCreateMap != nil {
		return nil, errDispatchPostCreateMap
	}

	return astMap, nil
}

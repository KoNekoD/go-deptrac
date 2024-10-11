package ast_map

import (
	"github.com/KoNekoD/go-deptrac/pkg/application/event_dispatchers"
	"github.com/KoNekoD/go-deptrac/pkg/application/services/parsers"
	"github.com/KoNekoD/go-deptrac/pkg/domain/dtos/ast_maps"
	"github.com/KoNekoD/go-deptrac/pkg/domain/dtos/tokens_references"
	"github.com/KoNekoD/go-deptrac/pkg/domain/events"
)

type AstLoader struct {
	parser          parsers.ParserInterface
	eventDispatcher event_dispatchers.EventDispatcherInterface
}

func NewAstLoader(parser parsers.ParserInterface, eventDispatcher event_dispatchers.EventDispatcherInterface) *AstLoader {
	return &AstLoader{
		parser:          parser,
		eventDispatcher: eventDispatcher,
	}
}

func (l *AstLoader) CreateAstMap(files []string) (*ast_maps.AstMap, error) {
	references := make([]*tokens_references.FileReference, 0)

	err := l.eventDispatcher.DispatchEvent(events.NewPreCreateAstMapEvent(len(files)))
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		reference, err := l.parser.ParseFile(file)
		if err != nil {
			err := l.eventDispatcher.DispatchEvent(events.NewAstFileSyntaxErrorEvent(file, err.Error()))
			if err != nil {
				return nil, err
			}

			continue
		}

		references = append(references, reference)

		errDispatchAnalysed := l.eventDispatcher.DispatchEvent(events.NewAstFileAnalysedEvent(file))
		if errDispatchAnalysed != nil {
			return nil, errDispatchAnalysed
		}
	}

	astMap := ast_maps.NewAstMap(references)

	errDispatchPostCreateMap := l.eventDispatcher.DispatchEvent(events.NewPostCreateAstMapEvent())
	if errDispatchPostCreateMap != nil {
		return nil, errDispatchPostCreateMap
	}

	return astMap, nil
}

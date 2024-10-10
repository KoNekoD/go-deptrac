package ast_map

import (
	"github.com/KoNekoD/go-deptrac/pkg/events"
	"github.com/KoNekoD/go-deptrac/pkg/parsers"
	"github.com/KoNekoD/go-deptrac/pkg/references"
)

type AstLoader struct {
	parser          parsers.ParserInterface
	eventDispatcher events.EventDispatcherInterface
}

func NewAstLoader(parser parsers.ParserInterface, eventDispatcher events.EventDispatcherInterface) *AstLoader {
	return &AstLoader{
		parser:          parser,
		eventDispatcher: eventDispatcher,
	}
}

func (l *AstLoader) CreateAstMap(files []string) (*AstMap, error) {
	references := make([]*references.FileReference, 0)

	err := l.eventDispatcher.DispatchEvent(NewPreCreateAstMapEvent(len(files)))
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		reference, err := l.parser.ParseFile(file)
		if err != nil {
			err := l.eventDispatcher.DispatchEvent(NewAstFileSyntaxErrorEvent(file, err.Error()))
			if err != nil {
				return nil, err
			}

			continue
		}

		references = append(references, reference)

		errDispatchAnalysed := l.eventDispatcher.DispatchEvent(NewAstFileAnalysedEvent(file))
		if errDispatchAnalysed != nil {
			return nil, errDispatchAnalysed
		}
	}

	astMap := NewAstMap(references)

	errDispatchPostCreateMap := l.eventDispatcher.DispatchEvent(NewPostCreateAstMapEvent())
	if errDispatchPostCreateMap != nil {
		return nil, errDispatchPostCreateMap
	}

	return astMap, nil
}

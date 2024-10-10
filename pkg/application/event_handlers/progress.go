package event_handlers

import (
	"github.com/KoNekoD/go-deptrac/pkg/domain/events"
	"github.com/KoNekoD/go-deptrac/pkg/infrastructure/services"
)

type Progress struct {
	output services.OutputInterface
}

func NewProgress(output services.OutputInterface) *Progress {
	return &Progress{output: output}
}

func (s *Progress) HandleEvent(rawEvent interface{}, stopPropagation func()) error {
	switch event := rawEvent.(type) {
	case *events.PreCreateAstMapEvent:
		s.output.GetStyle().ProgressStart(event.ExpectedFileCount)
	case *events.PostCreateAstMapEvent:
		err := s.output.GetStyle().ProgressFinish()
		if err != nil {
			return err
		}
	case *events.AstFileAnalysedEvent:
		err := s.output.GetStyle().ProgressAdvance(services.ProgressAdvanceDefault)
		if err != nil {
			return err
		}
	}

	return nil
}

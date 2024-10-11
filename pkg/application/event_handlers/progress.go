package event_handlers

import (
	services2 "github.com/KoNekoD/go-deptrac/pkg/application/services"
	"github.com/KoNekoD/go-deptrac/pkg/domain/events"
)

type Progress struct {
	output services2.OutputInterface
}

func NewProgress(output services2.OutputInterface) *Progress {
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
		err := s.output.GetStyle().ProgressAdvance(services2.ProgressAdvanceDefault)
		if err != nil {
			return err
		}
	}

	return nil
}

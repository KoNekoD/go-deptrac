package subscribers

import (
	"github.com/KoNekoD/go-deptrac/pkg/ast_map"
	"github.com/KoNekoD/go-deptrac/pkg/domain/events"
	"github.com/KoNekoD/go-deptrac/pkg/results"
)

type ProgressSubscriber struct {
	output results.OutputInterface
}

func NewProgressSubscriber(output results.OutputInterface) *ProgressSubscriber {
	return &ProgressSubscriber{
		output: output,
	}
}

func (s *ProgressSubscriber) InvokeEventSubscriber(rawEvent interface{}, stopPropagation func()) error {
	switch event := rawEvent.(type) {
	case *ast_map.PreCreateAstMapEvent:
		s.output.GetStyle().ProgressStart(event.ExpectedFileCount)
	case *ast_map.PostCreateAstMapEvent:
		err := s.output.GetStyle().ProgressFinish()
		if err != nil {
			return err
		}
	case *events.AstFileAnalysedEvent:
		err := s.output.GetStyle().ProgressAdvance(results.ProgressAdvanceDefault)
		if err != nil {
			return err
		}
	}

	return nil
}

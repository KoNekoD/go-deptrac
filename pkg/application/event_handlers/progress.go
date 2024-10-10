package event_handlers

import (
	"github.com/KoNekoD/go-deptrac/pkg"
	"github.com/KoNekoD/go-deptrac/pkg/ast_map"
	"github.com/KoNekoD/go-deptrac/pkg/domain/events"
)

type Progress struct {
	output pkg.OutputInterface
}

func NewProgress(output pkg.OutputInterface) *Progress {
	return &Progress{output: output}
}

func (s *Progress) HandleEvent(rawEvent interface{}, stopPropagation func()) error {
	switch event := rawEvent.(type) {
	case *ast_map.PreCreateAstMapEvent:
		s.output.GetStyle().ProgressStart(event.ExpectedFileCount)
	case *ast_map.PostCreateAstMapEvent:
		err := s.output.GetStyle().ProgressFinish()
		if err != nil {
			return err
		}
	case *events.AstFileAnalysedEvent:
		err := s.output.GetStyle().ProgressAdvance(pkg.ProgressAdvanceDefault)
		if err != nil {
			return err
		}
	}

	return nil
}

package subscriber

import (
	"github.com/KoNekoD/go-deptrac/pkg/contract/ast"
	"github.com/KoNekoD/go-deptrac/pkg/contract/output_formatter"
)

type ProgressSubscriber struct {
	output output_formatter.OutputInterface
}

func NewProgressSubscriber(output output_formatter.OutputInterface) *ProgressSubscriber {
	return &ProgressSubscriber{
		output: output,
	}
}

func (s *ProgressSubscriber) InvokeEventSubscriber(rawEvent interface{}, stopPropagation func()) error {
	switch event := rawEvent.(type) {
	case *ast.PreCreateAstMapEvent:
		s.output.GetStyle().ProgressStart(event.ExpectedFileCount)
	case *ast.PostCreateAstMapEvent:
		err := s.output.GetStyle().ProgressFinish()
		if err != nil {
			return err
		}
	case *ast.AstFileAnalysedEvent:
		err := s.output.GetStyle().ProgressAdvance(output_formatter.ProgressAdvanceDefault)
		if err != nil {
			return err
		}
	}

	return nil
}

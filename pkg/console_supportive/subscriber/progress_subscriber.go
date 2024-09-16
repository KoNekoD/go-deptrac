package subscriber

import (
	ast_contract2 "github.com/KoNekoD/go-deptrac/pkg/ast_contract"
	output_formatter_contract2 "github.com/KoNekoD/go-deptrac/pkg/output_formatter_contract"
)

type ProgressSubscriber struct {
	output output_formatter_contract2.OutputInterface
}

func NewProgressSubscriber(output output_formatter_contract2.OutputInterface) *ProgressSubscriber {
	return &ProgressSubscriber{
		output: output,
	}
}

func (s *ProgressSubscriber) InvokeEventSubscriber(rawEvent interface{}, stopPropagation func()) error {
	switch event := rawEvent.(type) {
	case *ast_contract2.PreCreateAstMapEvent:
		s.output.GetStyle().ProgressStart(event.ExpectedFileCount)
	case *ast_contract2.PostCreateAstMapEvent:
		err := s.output.GetStyle().ProgressFinish()
		if err != nil {
			return err
		}
	case *ast_contract2.AstFileAnalysedEvent:
		err := s.output.GetStyle().ProgressAdvance(output_formatter_contract2.ProgressAdvanceDefault)
		if err != nil {
			return err
		}
	}

	return nil
}

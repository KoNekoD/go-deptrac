package ProgressSubscriber

import (
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Ast/AstFileAnalysedEvent"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Ast/PostCreateAstMapEvent"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Ast/PreCreateAstMapEvent"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/OutputFormatter/OutputInterface"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/OutputFormatter/OutputStyleInterface"
)

type ProgressSubscriber struct {
	output OutputInterface.OutputInterface
}

func NewProgressSubscriber(output OutputInterface.OutputInterface) *ProgressSubscriber {
	return &ProgressSubscriber{
		output: output,
	}
}

func (s *ProgressSubscriber) InvokeEventSubscriber(rawEvent interface{}, stopPropagation func()) error {
	switch event := rawEvent.(type) {
	case *PreCreateAstMapEvent.PreCreateAstMapEvent:
		s.output.GetStyle().ProgressStart(event.ExpectedFileCount)
	case *PostCreateAstMapEvent.PostCreateAstMapEvent:
		err := s.output.GetStyle().ProgressFinish()
		if err != nil {
			return err
		}
	case *AstFileAnalysedEvent.AstFileAnalysedEvent:
		err := s.output.GetStyle().ProgressAdvance(OutputStyleInterface.ProgressAdvanceDefault)
		if err != nil {
			return err
		}
	}

	return nil
}

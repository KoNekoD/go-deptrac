package ConsoleSubscriber

import (
	"fmt"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Ast/AstFileAnalysedEvent"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Ast/AstFileSyntaxErrorEvent"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Ast/PostCreateAstMapEvent"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Ast/PreCreateAstMapEvent"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Dependency/PostEmitEvent"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Dependency/PostFlattenEvent"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Dependency/PreEmitEvent"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Dependency/PreFlattenEvent"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/OutputFormatter/OutputInterface"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/OutputFormatter/OutputStyleInterface"
	"github.com/KoNekoD/go-deptrac/pkg/src/supportive/TimeStopwatch"
)

type ConsoleSubscriber struct {
	output    OutputInterface.OutputInterface
	stopwatch *TimeStopwatch.Stopwatch
}

func NewConsoleSubscriber(output OutputInterface.OutputInterface, stopwatch *TimeStopwatch.Stopwatch) *ConsoleSubscriber {
	return &ConsoleSubscriber{
		output:    output,
		stopwatch: stopwatch,
	}
}

func (s *ConsoleSubscriber) InvokeEventSubscriber(rawEvent interface{}, stopPropagation func()) error {
	switch event := rawEvent.(type) {
	case *PreCreateAstMapEvent.PreCreateAstMapEvent:
		if s.output.IsVerbose() {
			err := s.stopwatchStart("ast")
			if err != nil {
				return err
			}
			s.output.WriteLineFormatted(OutputStyleInterface.StringOrArrayOfStrings{String: fmt.Sprintf("Start to create an AstMap for <info>%d</> Files.", event.ExpectedFileCount)})
		}
	case *PostCreateAstMapEvent.PostCreateAstMapEvent:
		if s.output.IsVerbose() {
			s.printMessageWithTime("ast", "<info>AstMap created in %01.2f sec.</>", "<info>AstMap created.</>")
		}
	case *AstFileAnalysedEvent.AstFileAnalysedEvent:
		if s.output.IsVerbose() {
			s.output.WriteLineFormatted(OutputStyleInterface.StringOrArrayOfStrings{String: fmt.Sprintf("Parsing File %s", event.File)})
		}
	case *AstFileSyntaxErrorEvent.AstFileSyntaxErrorEvent:
		s.output.WriteLineFormatted(OutputStyleInterface.StringOrArrayOfStrings{String: fmt.Sprintf("\nSyntax Error on File %s\n<error>%s</>\n", event.File, event.SyntaxError)})
	case *PreEmitEvent.PreEmitEvent:
		if s.output.IsVerbose() {
			err := s.stopwatchStart("deps")
			if err != nil {
				return err
			}
			s.output.WriteLineFormatted(OutputStyleInterface.StringOrArrayOfStrings{String: fmt.Sprintf("start emitting dependencies <info>%s</>", event.EmitterName)})
		}
	case *PostEmitEvent.PostEmitEvent:
		if s.output.IsVerbose() {
			s.printMessageWithTime("deps", "<info>Dependencies emitted in %01.2f sec.</>", "<info>Dependencies emitted.</>")
		}
	case *PreFlattenEvent.PreFlattenEvent:
		if s.output.IsVerbose() {
			err := s.stopwatchStart("flatten")
			if err != nil {
				return err
			}
			s.output.WriteLineFormatted(OutputStyleInterface.StringOrArrayOfStrings{String: "start flatten dependencies"})
		}
	case *PostFlattenEvent.PostFlattenEvent:
		if s.output.IsVerbose() {
			s.printMessageWithTime("flatten", "<info>Dependencies flattened in %01.2f sec.</>", "<info>Dependencies flattened.</>")
		}
	}

	return nil
}

func (s *ConsoleSubscriber) stopwatchStart(event string) error {
	err := s.stopwatch.Start(event)
	if err != nil {
		return err
	}

	return nil
}

func (s *ConsoleSubscriber) printMessageWithTime(event string, messageWithTime string, messageWithoutTime string) {
	period, err := s.stopwatch.Stop(event)

	if err != nil {
		s.output.WriteLineFormatted(OutputStyleInterface.StringOrArrayOfStrings{String: messageWithoutTime})
		return
	}

	s.output.WriteLineFormatted(OutputStyleInterface.StringOrArrayOfStrings{String: fmt.Sprintf(messageWithTime, period.ToSeconds())})
}

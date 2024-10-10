package subscribers

import (
	"fmt"
	"github.com/KoNekoD/go-deptrac/pkg/ast_map"
	events2 "github.com/KoNekoD/go-deptrac/pkg/domain/events"
	"github.com/KoNekoD/go-deptrac/pkg/domain/stopwatch"
	"github.com/KoNekoD/go-deptrac/pkg/emitters"
	"github.com/KoNekoD/go-deptrac/pkg/results"
)

type ConsoleSubscriber struct {
	output    results.OutputInterface
	stopwatch *stopwatch.Stopwatch
}

func NewConsoleSubscriber(output results.OutputInterface, stopwatch *stopwatch.Stopwatch) *ConsoleSubscriber {
	return &ConsoleSubscriber{
		output:    output,
		stopwatch: stopwatch,
	}
}

func (s *ConsoleSubscriber) InvokeEventSubscriber(rawEvent interface{}, stopPropagation func()) error {
	switch event := rawEvent.(type) {
	case *ast_map.PreCreateAstMapEvent:
		if s.output.IsVerbose() {
			err := s.stopwatchStart("ast_contract")
			if err != nil {
				return err
			}
			s.output.WriteLineFormatted(results.StringOrArrayOfStrings{String: fmt.Sprintf("Start to create an AstMap for <info>%d</> Files.", event.ExpectedFileCount)})
		}
	case *ast_map.PostCreateAstMapEvent:
		if s.output.IsVerbose() {
			s.printMessageWithTime("ast_contract", "<info>AstMap created in %01.2f sec.</>", "<info>AstMap created.</>")
		}
	case *events2.AstFileAnalysedEvent:
		if s.output.IsVerbose() {
			s.output.WriteLineFormatted(results.StringOrArrayOfStrings{String: fmt.Sprintf("Parsing File %s", event.File)})
		}
	case *ast_map.AstFileSyntaxErrorEvent:
		s.output.WriteLineFormatted(results.StringOrArrayOfStrings{String: fmt.Sprintf("\nSyntax Error on File %s\n<error>%s</>\n", event.File, event.SyntaxError)})
	case *emitters.PreEmitEvent:
		if s.output.IsVerbose() {
			err := s.stopwatchStart("deps")
			if err != nil {
				return err
			}
			s.output.WriteLineFormatted(results.StringOrArrayOfStrings{String: fmt.Sprintf("start emitting dependencies <info>%s</>", event.EmitterName)})
		}
	case *emitters.PostEmitEvent:
		if s.output.IsVerbose() {
			s.printMessageWithTime("deps", "<info>Dependencies emitted in %01.2f sec.</>", "<info>Dependencies emitted.</>")
		}
	case *events2.PreFlattenEvent:
		if s.output.IsVerbose() {
			err := s.stopwatchStart("flatten")
			if err != nil {
				return err
			}
			s.output.WriteLineFormatted(results.StringOrArrayOfStrings{String: "start flatten dependencies"})
		}
	case *events2.PostFlattenEvent:
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
		s.output.WriteLineFormatted(results.StringOrArrayOfStrings{String: messageWithoutTime})
		return
	}

	s.output.WriteLineFormatted(results.StringOrArrayOfStrings{String: fmt.Sprintf(messageWithTime, period.ToSeconds())})
}

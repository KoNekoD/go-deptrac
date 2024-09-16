package subscriber

import (
	"fmt"
	ast_contract2 "github.com/KoNekoD/go-deptrac/pkg/ast_contract"
	dependency_contract2 "github.com/KoNekoD/go-deptrac/pkg/dependency_contract"
	output_formatter_contract2 "github.com/KoNekoD/go-deptrac/pkg/output_formatter_contract"
	"github.com/KoNekoD/go-deptrac/pkg/time_stopwatch_supportive"
)

type ConsoleSubscriber struct {
	output    output_formatter_contract2.OutputInterface
	stopwatch *time_stopwatch_supportive.Stopwatch
}

func NewConsoleSubscriber(output output_formatter_contract2.OutputInterface, stopwatch *time_stopwatch_supportive.Stopwatch) *ConsoleSubscriber {
	return &ConsoleSubscriber{
		output:    output,
		stopwatch: stopwatch,
	}
}

func (s *ConsoleSubscriber) InvokeEventSubscriber(rawEvent interface{}, stopPropagation func()) error {
	switch event := rawEvent.(type) {
	case *ast_contract2.PreCreateAstMapEvent:
		if s.output.IsVerbose() {
			err := s.stopwatchStart("ast_contract")
			if err != nil {
				return err
			}
			s.output.WriteLineFormatted(output_formatter_contract2.StringOrArrayOfStrings{String: fmt.Sprintf("Start to create an AstMap for <info>%d</> Files.", event.ExpectedFileCount)})
		}
	case *ast_contract2.PostCreateAstMapEvent:
		if s.output.IsVerbose() {
			s.printMessageWithTime("ast_contract", "<info>AstMap created in %01.2f sec.</>", "<info>AstMap created.</>")
		}
	case *ast_contract2.AstFileAnalysedEvent:
		if s.output.IsVerbose() {
			s.output.WriteLineFormatted(output_formatter_contract2.StringOrArrayOfStrings{String: fmt.Sprintf("Parsing File %s", event.File)})
		}
	case *ast_contract2.AstFileSyntaxErrorEvent:
		s.output.WriteLineFormatted(output_formatter_contract2.StringOrArrayOfStrings{String: fmt.Sprintf("\nSyntax Error on File %s\n<error>%s</>\n", event.File, event.SyntaxError)})
	case *dependency_contract2.PreEmitEvent:
		if s.output.IsVerbose() {
			err := s.stopwatchStart("deps")
			if err != nil {
				return err
			}
			s.output.WriteLineFormatted(output_formatter_contract2.StringOrArrayOfStrings{String: fmt.Sprintf("start emitting dependencies <info>%s</>", event.EmitterName)})
		}
	case *dependency_contract2.PostEmitEvent:
		if s.output.IsVerbose() {
			s.printMessageWithTime("deps", "<info>Dependencies emitted in %01.2f sec.</>", "<info>Dependencies emitted.</>")
		}
	case *dependency_contract2.PreFlattenEvent:
		if s.output.IsVerbose() {
			err := s.stopwatchStart("flatten")
			if err != nil {
				return err
			}
			s.output.WriteLineFormatted(output_formatter_contract2.StringOrArrayOfStrings{String: "start flatten dependencies"})
		}
	case *dependency_contract2.PostFlattenEvent:
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
		s.output.WriteLineFormatted(output_formatter_contract2.StringOrArrayOfStrings{String: messageWithoutTime})
		return
	}

	s.output.WriteLineFormatted(output_formatter_contract2.StringOrArrayOfStrings{String: fmt.Sprintf(messageWithTime, period.ToSeconds())})
}

package event_handlers

import (
	"fmt"
	"github.com/KoNekoD/go-deptrac/pkg/domain/events"
	"github.com/KoNekoD/go-deptrac/pkg/domain/stopwatch"
	"github.com/KoNekoD/go-deptrac/pkg/infrastructure/services"
)

type Console struct {
	output    services.OutputInterface
	stopwatch *stopwatch.Stopwatch
}

func NewConsole(output services.OutputInterface, stopwatch *stopwatch.Stopwatch) *Console {
	return &Console{
		output:    output,
		stopwatch: stopwatch,
	}
}

func (s *Console) HandleEvent(rawEvent interface{}, stopPropagation func()) error {
	switch event := rawEvent.(type) {
	case *events.PreCreateAstMapEvent:
		if s.output.IsVerbose() {
			err := s.stopwatchStart("ast_contract")
			if err != nil {
				return err
			}
			s.output.WriteLineFormatted(services.StringOrArrayOfStrings{String: fmt.Sprintf("Start to create an AstMap for <info>%d</> Files.", event.ExpectedFileCount)})
		}
	case *events.PostCreateAstMapEvent:
		if s.output.IsVerbose() {
			s.printMessageWithTime("ast_contract", "<info>AstMap created in %01.2f sec.</>", "<info>AstMap created.</>")
		}
	case *events.AstFileAnalysedEvent:
		if s.output.IsVerbose() {
			s.output.WriteLineFormatted(services.StringOrArrayOfStrings{String: fmt.Sprintf("Parsing File %s", event.File)})
		}
	case *events.AstFileSyntaxErrorEvent:
		s.output.WriteLineFormatted(services.StringOrArrayOfStrings{String: fmt.Sprintf("\nSyntax Error on File %s\n<error>%s</>\n", event.File, event.SyntaxError)})
	case *events.PreEmitEvent:
		if s.output.IsVerbose() {
			err := s.stopwatchStart("deps")
			if err != nil {
				return err
			}
			s.output.WriteLineFormatted(services.StringOrArrayOfStrings{String: fmt.Sprintf("start emitting dependencies <info>%s</>", event.EmitterName)})
		}
	case *events.PostEmitEvent:
		if s.output.IsVerbose() {
			s.printMessageWithTime("deps", "<info>Dependencies emitted in %01.2f sec.</>", "<info>Dependencies emitted.</>")
		}
	case *events.PreFlattenEvent:
		if s.output.IsVerbose() {
			err := s.stopwatchStart("flatten")
			if err != nil {
				return err
			}
			s.output.WriteLineFormatted(services.StringOrArrayOfStrings{String: "start flatten dependencies"})
		}
	case *events.PostFlattenEvent:
		if s.output.IsVerbose() {
			s.printMessageWithTime("flatten", "<info>Dependencies flattened in %01.2f sec.</>", "<info>Dependencies flattened.</>")
		}
	}

	return nil
}

func (s *Console) stopwatchStart(event string) error {
	err := s.stopwatch.Start(event)
	if err != nil {
		return err
	}

	return nil
}

func (s *Console) printMessageWithTime(event string, messageWithTime string, messageWithoutTime string) {
	period, err := s.stopwatch.Stop(event)

	if err != nil {
		s.output.WriteLineFormatted(services.StringOrArrayOfStrings{String: messageWithoutTime})
		return
	}

	s.output.WriteLineFormatted(services.StringOrArrayOfStrings{String: fmt.Sprintf(messageWithTime, period.ToSeconds())})
}

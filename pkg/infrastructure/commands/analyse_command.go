package commands

import (
	"github.com/KoNekoD/go-deptrac/pkg/application/event_dispatchers"
	"github.com/KoNekoD/go-deptrac/pkg/application/event_handlers"
	"github.com/KoNekoD/go-deptrac/pkg/domain/dtos/commands_options"
	"github.com/KoNekoD/go-deptrac/pkg/infrastructure/services"
	"github.com/KoNekoD/go-deptrac/pkg/infrastructure/services/formatters"
	"github.com/KoNekoD/go-deptrac/pkg/infrastructure/services/runners"
)

// AnalyseCommand - Analyses your project using the provided depfile
type AnalyseCommand struct {
	runner             *runners.AnalyseRunner
	dispatcher         event_dispatchers.EventDispatcherInterface
	formatterProvider  *formatters.FormatterProvider
	verboseBoolFlag    bool
	debugBoolFlag      bool
	consoleSubscriber  *event_handlers.Console
	progressSubscriber *event_handlers.Progress
	analyseOptions     *commands_options.AnalyseOptions
}

func NewAnalyseCommand(runner *runners.AnalyseRunner, dispatcher event_dispatchers.EventDispatcherInterface, formatterProvider *formatters.FormatterProvider, verboseBoolFlag bool, debugBoolFlag bool, consoleSubscriber *event_handlers.Console, progressSubscriber *event_handlers.Progress, analyseOptions *commands_options.AnalyseOptions) *AnalyseCommand {
	return &AnalyseCommand{
		runner:             runner,
		dispatcher:         dispatcher,
		formatterProvider:  formatterProvider,
		verboseBoolFlag:    verboseBoolFlag,
		debugBoolFlag:      debugBoolFlag,
		consoleSubscriber:  consoleSubscriber,
		progressSubscriber: progressSubscriber,
		analyseOptions:     analyseOptions,
	}
}

func (c *AnalyseCommand) Run() error {
	symfonyOutput := services.NewSymfonyOutput(formatters.NewStyle(c.verboseBoolFlag, c.debugBoolFlag))

	// Moved to services
	//event_subscriber_interface_map_reg.RegForAnalyseCommand(c.consoleSubscriber, c.progressSubscriber, !options.NoProgress)

	err := c.runner.Run(c.analyseOptions, symfonyOutput)
	if err != nil {
		return err
	}

	return nil
}

package commands

import (
	event_handlers2 "github.com/KoNekoD/go-deptrac/pkg/application/event_handlers"
	"github.com/KoNekoD/go-deptrac/pkg/domain/dtos/options"
	"github.com/KoNekoD/go-deptrac/pkg/infrastructure/app"
	"github.com/KoNekoD/go-deptrac/pkg/infrastructure/services"
	"github.com/KoNekoD/go-deptrac/pkg/infrastructure/services/dispatchers"
	formatters2 "github.com/KoNekoD/go-deptrac/pkg/infrastructure/services/formatters"
)

// AnalyseCommand - Analyses your project using the provided depfile
type AnalyseCommand struct {
	runner             *app.AnalyseRunner
	dispatcher         dispatchers.EventDispatcherInterface
	formatterProvider  *formatters2.FormatterProvider
	verboseBoolFlag    bool
	debugBoolFlag      bool
	consoleSubscriber  *event_handlers2.Console
	progressSubscriber *event_handlers2.Progress
	analyseOptions     *options.AnalyseOptions
}

func NewAnalyseCommand(runner *app.AnalyseRunner, dispatcher dispatchers.EventDispatcherInterface, formatterProvider *formatters2.FormatterProvider, verboseBoolFlag bool, debugBoolFlag bool, consoleSubscriber *event_handlers2.Console, progressSubscriber *event_handlers2.Progress, analyseOptions *options.AnalyseOptions) *AnalyseCommand {
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
	symfonyOutput := services.NewSymfonyOutput(formatters2.NewStyle(c.verboseBoolFlag, c.debugBoolFlag))

	// Moved to services
	//event_subscriber_interface_map_reg.RegForAnalyseCommand(c.consoleSubscriber, c.progressSubscriber, !options.NoProgress)

	err := c.runner.Run(c.analyseOptions, symfonyOutput)
	if err != nil {
		return err
	}

	return nil
}

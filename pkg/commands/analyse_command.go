package commands

import (
	"github.com/KoNekoD/go-deptrac/pkg/app"
	event_handlers2 "github.com/KoNekoD/go-deptrac/pkg/application/event_handlers"
	"github.com/KoNekoD/go-deptrac/pkg/dispatchers"
	"github.com/KoNekoD/go-deptrac/pkg/formatters"
	"github.com/KoNekoD/go-deptrac/pkg/results"
	"github.com/KoNekoD/go-deptrac/pkg/rules"
)

// AnalyseCommand - Analyses your project using the provided depfile
type AnalyseCommand struct {
	runner             *app.AnalyseRunner
	dispatcher         dispatchers.EventDispatcherInterface
	formatterProvider  *formatters.FormatterProvider
	verboseBoolFlag    bool
	debugBoolFlag      bool
	consoleSubscriber  *event_handlers2.Console
	progressSubscriber *event_handlers2.Progress
	analyseOptions     *rules.AnalyseOptions
}

func NewAnalyseCommand(runner *app.AnalyseRunner, dispatcher dispatchers.EventDispatcherInterface, formatterProvider *formatters.FormatterProvider, verboseBoolFlag bool, debugBoolFlag bool, consoleSubscriber *event_handlers2.Console, progressSubscriber *event_handlers2.Progress, analyseOptions *rules.AnalyseOptions) *AnalyseCommand {
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
	symfonyOutput := results.NewSymfonyOutput(formatters.NewStyle(c.verboseBoolFlag, c.debugBoolFlag))

	// Moved to services
	//event_subscriber_interface_map_reg.RegForAnalyseCommand(c.consoleSubscriber, c.progressSubscriber, !options.NoProgress)

	err := c.runner.Run(c.analyseOptions, symfonyOutput)
	if err != nil {
		return err
	}

	return nil
}

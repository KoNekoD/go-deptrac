package commands

import (
	"github.com/KoNekoD/go-deptrac/pkg/events"
	"github.com/KoNekoD/go-deptrac/pkg/formatters"
	"github.com/KoNekoD/go-deptrac/pkg/results"
	"github.com/KoNekoD/go-deptrac/pkg/rules"
	"github.com/KoNekoD/go-deptrac/pkg/runners"
	"github.com/KoNekoD/go-deptrac/pkg/subscribers"
)

// AnalyseCommand - Analyses your project using the provided depfile
type AnalyseCommand struct {
	runner             *runners.AnalyseRunner
	dispatcher         events.EventDispatcherInterface
	formatterProvider  *formatters.FormatterProvider
	verboseBoolFlag    bool
	debugBoolFlag      bool
	consoleSubscriber  *subscribers.ConsoleSubscriber
	progressSubscriber *subscribers.ProgressSubscriber
	analyseOptions     *rules.AnalyseOptions
}

func NewAnalyseCommand(runner *runners.AnalyseRunner, dispatcher events.EventDispatcherInterface, formatterProvider *formatters.FormatterProvider, verboseBoolFlag bool, debugBoolFlag bool, consoleSubscriber *subscribers.ConsoleSubscriber, progressSubscriber *subscribers.ProgressSubscriber, analyseOptions *rules.AnalyseOptions) *AnalyseCommand {
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

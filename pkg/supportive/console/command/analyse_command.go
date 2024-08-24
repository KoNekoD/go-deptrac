package command

import (
	"github.com/KoNekoD/go-deptrac/pkg/supportive/console/subscriber"
	"github.com/KoNekoD/go-deptrac/pkg/supportive/console/symfony"
	"github.com/KoNekoD/go-deptrac/pkg/supportive/dependency_injection/event_dispatcher/event_dispatcher_interface"
	output_formatter2 "github.com/KoNekoD/go-deptrac/pkg/supportive/output_formatter"
)

// AnalyseCommand - Analyses your project using the provided depfile
type AnalyseCommand struct {
	runner             *AnalyseRunner
	dispatcher         event_dispatcher_interface.EventDispatcherInterface
	formatterProvider  *output_formatter2.FormatterProvider
	verboseBoolFlag    bool
	debugBoolFlag      bool
	consoleSubscriber  *subscriber.ConsoleSubscriber
	progressSubscriber *subscriber.ProgressSubscriber
	analyseOptions     *AnalyseOptions
}

func NewAnalyseCommand(runner *AnalyseRunner, dispatcher event_dispatcher_interface.EventDispatcherInterface, formatterProvider *output_formatter2.FormatterProvider, verboseBoolFlag bool, debugBoolFlag bool, consoleSubscriber *subscriber.ConsoleSubscriber, progressSubscriber *subscriber.ProgressSubscriber, analyseOptions *AnalyseOptions) *AnalyseCommand {
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
	symfonyOutput := symfony.NewSymfonyOutput(symfony.NewStyle(c.verboseBoolFlag, c.debugBoolFlag))

	// Moved to services
	//event_subscriber_interface_map_reg.RegForAnalyseCommand(c.consoleSubscriber, c.progressSubscriber, !options.NoProgress)

	err := c.runner.Run(c.analyseOptions, symfonyOutput)
	if err != nil {
		return err
	}

	return nil
}

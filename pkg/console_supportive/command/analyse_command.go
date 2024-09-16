package command

import (
	subscriber2 "github.com/KoNekoD/go-deptrac/pkg/console_supportive/subscriber"
	symfony2 "github.com/KoNekoD/go-deptrac/pkg/console_supportive/symfony"
	"github.com/KoNekoD/go-deptrac/pkg/dependency_injection_supportive/event_dispatcher/event_dispatcher_interface"
	output_formatter2 "github.com/KoNekoD/go-deptrac/pkg/output_formatter_supportive"
)

// AnalyseCommand - Analyses your project using the provided depfile
type AnalyseCommand struct {
	runner             *AnalyseRunner
	dispatcher         event_dispatcher_interface.EventDispatcherInterface
	formatterProvider  *output_formatter2.FormatterProvider
	verboseBoolFlag    bool
	debugBoolFlag      bool
	consoleSubscriber  *subscriber2.ConsoleSubscriber
	progressSubscriber *subscriber2.ProgressSubscriber
	analyseOptions     *AnalyseOptions
}

func NewAnalyseCommand(runner *AnalyseRunner, dispatcher event_dispatcher_interface.EventDispatcherInterface, formatterProvider *output_formatter2.FormatterProvider, verboseBoolFlag bool, debugBoolFlag bool, consoleSubscriber *subscriber2.ConsoleSubscriber, progressSubscriber *subscriber2.ProgressSubscriber, analyseOptions *AnalyseOptions) *AnalyseCommand {
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
	symfonyOutput := symfony2.NewSymfonyOutput(symfony2.NewStyle(c.verboseBoolFlag, c.debugBoolFlag))

	// Moved to services
	//event_subscriber_interface_map_reg.RegForAnalyseCommand(c.consoleSubscriber, c.progressSubscriber, !options.NoProgress)

	err := c.runner.Run(c.analyseOptions, symfonyOutput)
	if err != nil {
		return err
	}

	return nil
}

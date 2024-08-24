package command

import (
	"flag"
	"fmt"
	"github.com/KoNekoD/go-deptrac/pkg/contract/output_formatter"
	"github.com/KoNekoD/go-deptrac/pkg/supportive/console"
	"github.com/KoNekoD/go-deptrac/pkg/supportive/console/subscriber"
	"github.com/KoNekoD/go-deptrac/pkg/supportive/console/symfony"
	"github.com/KoNekoD/go-deptrac/pkg/supportive/dependency_injection/event_dispatcher/event_dispatcher_interface"
	"github.com/KoNekoD/go-deptrac/pkg/supportive/dependency_injection/event_subscriber_interface_map/event_subscriber_interface_map_reg"
	output_formatter2 "github.com/KoNekoD/go-deptrac/pkg/supportive/output_formatter"
	"strings"
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
}

func NewAnalyseCommand(runner *AnalyseRunner, dispatcher event_dispatcher_interface.EventDispatcherInterface, formatterProvider *output_formatter2.FormatterProvider, verboseBoolFlag bool, debugBoolFlag bool, consoleSubscriber *subscriber.ConsoleSubscriber, progressSubscriber *subscriber.ProgressSubscriber) *AnalyseCommand {
	return &AnalyseCommand{
		runner:             runner,
		dispatcher:         dispatcher,
		formatterProvider:  formatterProvider,
		verboseBoolFlag:    verboseBoolFlag,
		debugBoolFlag:      debugBoolFlag,
		consoleSubscriber:  consoleSubscriber,
		progressSubscriber: progressSubscriber,
	}
}

func (c *AnalyseCommand) Run() error {
	knownFormattersStr := make([]string, 0)
	for _, formatterType := range c.formatterProvider.GetKnownFormatters() {
		knownFormattersStr = append(knownFormattersStr, fmt.Sprintf("\"%s\"", formatterType))
	}

	var (
		formatterUsagePossible = strings.Join(knownFormattersStr, ", ")
		formatterUsage         = fmt.Sprintf("Format in which to print the result of the analysis. Possible: [\"%s\"]", formatterUsagePossible)
		formatter              = flag.String("formatter", string(output_formatter.Table), formatterUsage)

		output          = flag.String("output", "", "Output file path for formatter (if applicable)")
		noProgress      = flag.Bool("no-progress", false, "Do not show progress bar")
		reportSkipped   = flag.Bool("report-skipped", false, "Report skipped violations")
		reportUncovered = flag.Bool("report-uncovered", false, "Report uncovered dependencies")
		failOnUncovered = flag.Bool("fail-on-uncovered", false, "Fails if any uncovered dependency is found")
	)

	symfonyOutput := symfony.NewSymfonyOutput(symfony.NewStyle(c.verboseBoolFlag, c.debugBoolFlag))

	if formatter == nil {
		formatterTmp := string(getDefaultFormatter())
		formatter = &formatterTmp
	}

	options := NewAnalyseOptions(
		nil != noProgress && *noProgress == true,
		*formatter,
		output,
		nil != reportSkipped && *reportSkipped == true,
		nil != reportUncovered && *reportUncovered == true,
		nil != failOnUncovered && *failOnUncovered == true,
	)

	event_subscriber_interface_map_reg.RegForAnalyseCommand(c.consoleSubscriber, c.progressSubscriber, !options.NoProgress)

	err := c.runner.Run(options, symfonyOutput)
	if err != nil {
		return err
	}

	return nil
}

func getDefaultFormatter() output_formatter.OutputFormatterType {
	if console.NewEnv().GetEnv("GITHUB_ACTIONS") != "" {
		return output_formatter2.NewGithubActionsOutputFormatter().GetName()
	}
	return output_formatter2.NewTableOutputFormatter().GetName()
}

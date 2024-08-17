package AnalyseCommand

import (
	"flag"
	"fmt"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/OutputFormatter/OutputFormatterInterface/OutputFormatterType"
	"github.com/KoNekoD/go-deptrac/pkg/src/supportive/Console/Command/AnalyseOptions"
	"github.com/KoNekoD/go-deptrac/pkg/src/supportive/Console/Command/AnalyseRunner"
	"github.com/KoNekoD/go-deptrac/pkg/src/supportive/Console/Env"
	"github.com/KoNekoD/go-deptrac/pkg/src/supportive/Console/Subscriber/ConsoleSubscriber"
	"github.com/KoNekoD/go-deptrac/pkg/src/supportive/Console/Subscriber/ProgressSubscriber"
	"github.com/KoNekoD/go-deptrac/pkg/src/supportive/Console/Symfony/Style"
	"github.com/KoNekoD/go-deptrac/pkg/src/supportive/Console/Symfony/SymfonyOutput"
	"github.com/KoNekoD/go-deptrac/pkg/src/supportive/DependencyInjection/EventDispatcher/EventDispatcherInterface"
	"github.com/KoNekoD/go-deptrac/pkg/src/supportive/DependencyInjection/EventSubscriberInterfaceMap/EventSubscriberInterfaceMapReg"
	"github.com/KoNekoD/go-deptrac/pkg/src/supportive/OutputFormatter/FormatterProvider"
	"github.com/KoNekoD/go-deptrac/pkg/src/supportive/OutputFormatter/GithubActionsOutputFormatter"
	"github.com/KoNekoD/go-deptrac/pkg/src/supportive/OutputFormatter/TableOutputFormatter"
	"strings"
)

// AnalyseCommand - Analyses your project using the provided depfile
type AnalyseCommand struct {
	runner             *AnalyseRunner.AnalyseRunner
	dispatcher         util.EventDispatcherInterface
	formatterProvider  *FormatterProvider.FormatterProvider
	verboseBoolFlag    bool
	debugBoolFlag      bool
	consoleSubscriber  *ConsoleSubscriber.ConsoleSubscriber
	progressSubscriber *ProgressSubscriber.ProgressSubscriber
}

func NewAnalyseCommand(runner *AnalyseRunner.AnalyseRunner, dispatcher util.EventDispatcherInterface, formatterProvider *FormatterProvider.FormatterProvider, verboseBoolFlag bool, debugBoolFlag bool, consoleSubscriber *ConsoleSubscriber.ConsoleSubscriber, progressSubscriber *ProgressSubscriber.ProgressSubscriber) *AnalyseCommand {
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
		formatter              = flag.String("formatter", string(OutputFormatterType.Table), formatterUsage)

		output          = flag.String("output", "", "Output file path for formatter (if applicable)")
		noProgress      = flag.Bool("no-progress", false, "Do not show progress bar")
		reportSkipped   = flag.Bool("report-skipped", false, "Report skipped violations")
		reportUncovered = flag.Bool("report-uncovered", false, "Report uncovered dependencies")
		failOnUncovered = flag.Bool("fail-on-uncovered", false, "Fails if any uncovered dependency is found")
	)

	symfonyOutput := SymfonyOutput.NewSymfonyOutput(Style.NewStyle(c.verboseBoolFlag, c.debugBoolFlag))

	if formatter == nil {
		formatterTmp := string(getDefaultFormatter())
		formatter = &formatterTmp
	}

	options := AnalyseOptions.NewAnalyseOptions(
		nil != noProgress && *noProgress == true,
		*formatter,
		output,
		nil != reportSkipped && *reportSkipped == true,
		nil != reportUncovered && *reportUncovered == true,
		nil != failOnUncovered && *failOnUncovered == true,
	)

	EventSubscriberInterfaceMapReg.RegForAnalyseCommand(c.consoleSubscriber, c.progressSubscriber, !options.NoProgress)

	err := c.runner.Run(options, symfonyOutput)
	if err != nil {
		return err
	}

	return nil
}

func getDefaultFormatter() OutputFormatterType.OutputFormatterType {
	if Env.NewEnv().GetEnv("GITHUB_ACTIONS") != "" {
		return GithubActionsOutputFormatter.NewGithubActionsOutputFormatter().GetName()
	}
	return TableOutputFormatter.NewTableOutputFormatter().GetName()
}

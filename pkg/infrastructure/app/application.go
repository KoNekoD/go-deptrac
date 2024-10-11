package app

import (
	"flag"
	"github.com/KoNekoD/go-deptrac/pkg/domain/apperrors"
	"github.com/KoNekoD/go-deptrac/pkg/infrastructure/di"
	"github.com/gookit/color"
	"os"
	"slices"
)

const Version = "@git-version@"

type Application struct {
	defaultCommand string
}

func NewApplication() *Application {
	return &Application{}
}

func (a *Application) Run() {
	_, err := a.doRun()
	if err != nil {
		color.Printf("\n<fg=167;bg=232>" + err.Error() + "</>\n")
	}
}

func (a *Application) getDefaultInputDefinition() {
	// $definition = parent::getDefaultInputDefinition();

	// return $definition;
}

const DirectorySeparator = "/"

func (a *Application) doRun() (int, error) {
	currentWorkingDirectory, err := os.Getwd()
	if err != nil {
		return 0, apperrors.NewCannotGetCurrentWorkingDirectoryExceptionCannotGetCWD()
	}

	// try {
	//     $input->bind($this->getDefinition());
	// } catch (ExceptionInterface) {
	//     // Errors must be ignored, full binding/validation happens later when the command is known.
	// }

	// if (null === $input->getArgument('command') && \true === $input->getOption('version')) {
	//     return parent::doRun($input, $output);
	// }

	var help bool
	flag.BoolVar(
		&help,
		"help",
		false,
		"Display help for the given command. When no command is given display help for the <info>analyse</> command",
	)
	var noCacheArgument bool
	flag.BoolVar(
		&noCacheArgument,
		"no-cache",
		false,
		"Disable caching mechanisms (wins over --cache-file_supportive)",
	)
	var clearCache bool
	flag.BoolVar(
		&clearCache,
		"clear-cache",
		false,
		"Clears cache file_supportive before run",
	)
	var cacheFile string
	flag.StringVar(
		&cacheFile,
		"cache-file_supportive",
		"",
		"Location where cache file_supportive will be stored",
	)
	var configFile string
	flag.StringVar(
		&configFile,
		"config",
		currentWorkingDirectory+DirectorySeparator+"deptrac.yaml",
		"Location of Depfile containing the configuration",
	)
	flag.Parse()

	var (
		commandArgument = flag.Arg(0)
	)

	config := currentWorkingDirectory + DirectorySeparator + "deptrac.yaml"
	if configFile != "" {
		config = configFile
	}

	cache := cacheFile

	factory := di.NewServiceContainerBuilder(currentWorkingDirectory)

	if !slices.Contains([]string{"init", "list", "help", "completion"}, commandArgument) {
		factory, err = factory.WithConfig(config)
		if err != nil {
			return 0, err
		}
	}

	noCache := false
	if noCacheArgument == true {
		noCache = true
	}

	var factoryBuildCache *string
	if !noCache {
		factoryBuildCache = &cache
	}

	_, err = factory.Build(factoryBuildCache, clearCache)
	if err != nil {
		if help == true {
			a.setDefaultCommand("help")
		} else {
			return 0, err
		}
	}

	switch commandArgument {
	case "analyse":
		err := factory.GetContainer().AnalyseCommand.Run()
		if err != nil {
			return 1, err
		}
	}

	// return parent::doRun($input, $output);
	return 0, nil
}

func (a *Application) setDefaultCommand(command string) {
	a.defaultCommand = command
}

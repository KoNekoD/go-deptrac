package Application

import (
	"flag"
	"github.com/KoNekoD/go-deptrac/pkg/src/supportive/Console/CannotGetCurrentWorkingDirectoryException"
	"github.com/KoNekoD/go-deptrac/pkg/src/supportive/DependencyInjection"
	"github.com/gookit/color"
	"os"
	"slices"
)

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
		return 0, CannotGetCurrentWorkingDirectoryException.NewCannotGetCurrentWorkingDirectoryExceptionCannotGetCWD()
	}

	// try {
	//     $input->bind($this->getDefinition());
	// } catch (ExceptionInterface) {
	//     // Errors must be ignored, full binding/validation happens later when the command is known.
	// }

	// if (null === $input->getArgument('command') && \true === $input->getOption('version')) {
	//     return parent::doRun($input, $output);
	// }

	var (
		commandArgument = flag.Arg(0)
	)

	var (
		help            = flag.Bool("help", false, "Display help for the given command. When no command is given display help for the <info>analyse</> command")
		noCacheArgument = flag.Bool("no-cache", false, "Disable caching mechanisms (wins over --cache-file)")
		clearCache      = flag.Bool("clear-cache", false, "Clears cache file before run")
		cacheFile       = flag.String("cache-file", "", "Location where cache file will be stored")
		configFile      = flag.String("config-file", currentWorkingDirectory+DirectorySeparator+"deptrac.yaml", "Location of Depfile containing the configuration")
	)

	config := currentWorkingDirectory + DirectorySeparator + "deptrac.yaml"
	if configFile != nil {
		config = *configFile
	}

	cache := cacheFile

	factory := DependencyInjection.NewServiceContainerBuilder(currentWorkingDirectory)

	if !slices.Contains([]string{"init", "list", "help", "completion"}, commandArgument) {
		factory, err = factory.WithConfig(config)
		if err != nil {
			return 0, err
		}
	}

	noCache := false
	if noCacheArgument != nil && *noCacheArgument == true {
		noCache = true
	}

	var factoryBuildCache *string
	if !noCache {
		factoryBuildCache = cache
	}

	_, err = factory.Build(factoryBuildCache, clearCache != nil && *clearCache == true)
	if err != nil {
		if help != nil && *help == true {
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
package di

import (
	"flag"
	"fmt"
	"github.com/KoNekoD/go-deptrac/pkg/application/event_dispatchers"
	"github.com/KoNekoD/go-deptrac/pkg/application/event_handlers"
	applicationServices "github.com/KoNekoD/go-deptrac/pkg/application/services"
	"github.com/KoNekoD/go-deptrac/pkg/application/services/analysers"
	"github.com/KoNekoD/go-deptrac/pkg/application/services/ast_file_reference_cache"
	"github.com/KoNekoD/go-deptrac/pkg/application/services/ast_map"
	"github.com/KoNekoD/go-deptrac/pkg/application/services/collectors_resolvers"
	"github.com/KoNekoD/go-deptrac/pkg/application/services/dependencies_collectors"
	"github.com/KoNekoD/go-deptrac/pkg/application/services/emitters"
	"github.com/KoNekoD/go-deptrac/pkg/application/services/input_collectors"
	"github.com/KoNekoD/go-deptrac/pkg/application/services/layers_resolvers"
	"github.com/KoNekoD/go-deptrac/pkg/application/services/parsers"
	"github.com/KoNekoD/go-deptrac/pkg/application/services/references_extractors"
	"github.com/KoNekoD/go-deptrac/pkg/application/services/types"
	"github.com/KoNekoD/go-deptrac/pkg/domain/dtos/commands_options"
	"github.com/KoNekoD/go-deptrac/pkg/domain/enums"
	"github.com/KoNekoD/go-deptrac/pkg/domain/events"
	domainServices "github.com/KoNekoD/go-deptrac/pkg/domain/services"
	"github.com/KoNekoD/go-deptrac/pkg/domain/stopwatch"
	"github.com/KoNekoD/go-deptrac/pkg/infrastructure/commands"
	"github.com/KoNekoD/go-deptrac/pkg/infrastructure/services"
	"github.com/KoNekoD/go-deptrac/pkg/infrastructure/services/formatters"
	"github.com/KoNekoD/go-deptrac/pkg/infrastructure/services/runners"
	"github.com/elliotchance/orderedmap/v2"
	"os"
	"strings"
)

func getDefaultFormatter() enums.OutputFormatterType {
	if os.Getenv("GITHUB_ACTIONS") != "" {
		return formatters.NewGithubActionsOutputFormatter().GetName()
	}
	return formatters.NewTableOutputFormatter().GetName()
}

func Services(builder *ContainerBuilder) error {

	cacheableFileSubscriber := builder.CacheableFileSubscriber
	builderConfiguration := builder.Configuration
	projectDirectory := builder.ProjectDirectory
	verboseBoolFlag := flag.Bool("verbose", true, "Verbose mode")
	debugBoolFlag := flag.Bool("debug", false, "Debug mode")
	style := formatters.NewStyle(
		verboseBoolFlag != nil && *verboseBoolFlag == true,
		debugBoolFlag != nil && *debugBoolFlag == true,
	)
	symfonyOutput := services.NewSymfonyOutput(style)

	timeStopwatch := stopwatch.NewStopwatch()

	nodeNamer := domainServices.NewNodeNamer(projectDirectory)

	/*
	 * Utilities
	 */
	eventDispatcher := event_dispatchers.NewEventDispatcher(debugBoolFlag != nil && *debugBoolFlag == true)

	fileInputCollector, err := input_collectors.NewFileInputCollector(
		builderConfiguration.Paths,
		builderConfiguration.ExcludeFiles,
		projectDirectory,
	)
	if err != nil {
		return err
	}

	ymlFileLoader := services.NewYmlFileLoader()
	dumper := domainServices.NewDumper("/deptrac_template.yaml")

	/*
	 * AST
	 */
	astFileReferenceInMemoryCache := ast_file_reference_cache.NewAstFileReferenceInMemoryCache()
	if builder.AstFileReferenceCacheInterface == nil {
		builder.AstFileReferenceCacheInterface = astFileReferenceInMemoryCache
	}
	typeResolver := types.NewTypeResolver(nodeNamer)
	referenceExtractors := []references_extractors.ReferenceExtractorInterface{
		references_extractors.NewFunctionLikeExtractor(typeResolver),
		references_extractors.NewPropertyExtractor(typeResolver),
		references_extractors.NewKeywordExtractor(typeResolver),
		references_extractors.NewFunctionCallResolver(typeResolver),
	}
	nikicPhpParser := parsers.NewNikicPhpParser(builder.AstFileReferenceCacheInterface, typeResolver, nodeNamer, referenceExtractors)
	parserInterface := nikicPhpParser
	astLoader := ast_map.NewAstLoader(parserInterface, eventDispatcher)

	/*
	 * Dependency
	 */
	dependencyEmitters := map[enums.EmitterType]emitters.DependencyEmitterInterface{
		enums.EmitterTypeClassToken:               emitters.NewClassDependencyEmitter(),
		enums.EmitterTypeClassSuperGlobalToken:    emitters.NewClassSuperglobalDependencyEmitter(),
		enums.EmitterTypeFileToken:                emitters.NewFileDependencyEmitter(),
		enums.EmitterTypeFunctionToken:            emitters.NewFunctionDependencyEmitter(),
		enums.EmitterTypeFunctionCall:             emitters.NewFunctionCallDependencyEmitter(),
		enums.EmitterTypeFunctionSuperGlobalToken: emitters.NewFunctionSuperglobalDependencyEmitter(),
		enums.EmitterTypeUseToken:                 emitters.NewUsesDependencyEmitter(),
	}
	inheritanceFlattener := applicationServices.NewInheritanceFlattener()
	dependencyResolver := applicationServices.NewDependencyResolver(builderConfiguration.Analyser, dependencyEmitters, inheritanceFlattener, eventDispatcher)
	tokenResolver := applicationServices.NewTokenResolver()

	astMapExtractor := ast_map.NewAstMapExtractor(fileInputCollector, astLoader)

	layerProvider := applicationServices.NewLayerProvider(builderConfiguration.Rulesets)
	eventHelper := applicationServices.NewEventHelper(builderConfiguration.SkipViolations, layerProvider)

	/*
	 * Events (before first possible event)
	 */
	/*
	 * Events
	 */
	event_dispatchers.Map = orderedmap.NewOrderedMap[string, *orderedmap.OrderedMap[int, []event_dispatchers.EventHandlerInterface]]()

	// Events
	uncoveredDependentHandler := event_handlers.NewUncoveredDependent(builderConfiguration.IgnoreUncoveredInternalStructs)
	matchingLayersHandler := event_handlers.NewMatchingLayers()
	allowDependencyHandler := event_handlers.NewAllowDependency()
	consoleSubscriber := event_handlers.NewConsole(symfonyOutput, timeStopwatch)
	progressSubscriber := event_handlers.NewProgress(symfonyOutput)
	dependsOnDisallowedLayer := event_handlers.NewDependsOnDisallowedLayer(eventHelper)
	dependsOnPrivateLayer := event_handlers.NewDependsOnPrivateLayer(eventHelper)
	dependsOnInternalToken := event_handlers.NewDependsOnInternalToken(eventHelper, builderConfiguration.Analyser)
	unmatchedSkippedViolations := event_handlers.NewUnmatchedSkippedViolations(eventHelper)

	processEvent := &events.ProcessEvent{}
	postProcessEvent := &events.PostProcessEvent{}
	preCreateAstMapEvent := &events.PreCreateAstMapEvent{}
	postCreateAstMapEvent := &events.PostCreateAstMapEvent{}
	// Events Handlers
	// TODO: Тут надо реализовать глобальный хук на параметры deptrac чтобы сделать что-то вида "param('skip_violations')"
	event_dispatchers.Reg(processEvent, allowDependencyHandler, -100)
	event_dispatchers.Reg(processEvent, dependsOnPrivateLayer, -3)
	event_dispatchers.Reg(processEvent, dependsOnInternalToken, -2)
	event_dispatchers.Reg(processEvent, dependsOnDisallowedLayer, -1)
	event_dispatchers.Reg(processEvent, matchingLayersHandler, 1)
	event_dispatchers.Reg(processEvent, uncoveredDependentHandler, 2)
	event_dispatchers.Reg(postProcessEvent, unmatchedSkippedViolations, event_dispatchers.DefaultPriority)
	if cacheableFileSubscriber != nil {
		event_dispatchers.Reg(preCreateAstMapEvent, cacheableFileSubscriber, event_dispatchers.DefaultPriority)
		event_dispatchers.Reg(postCreateAstMapEvent, cacheableFileSubscriber, event_dispatchers.DefaultPriority)
	}

	/*
	 * OutputFormatter
	 */
	outputFormatter := map[enums.OutputFormatterType]formatters.OutputFormatterInterface{
		enums.Table:         formatters.NewTableOutputFormatter(),
		enums.GithubActions: formatters.NewGithubActionsOutputFormatter(),
		// TODO:
		// $services->set(ConsoleOutputFormatter::class)->tag('output_formatter_contract');
		// $services->set(JUnitOutputFormatter::class)->tag('output_formatter_contract');
		// $services->set(XMLOutputFormatter::class)->tag('output_formatter_contract');
		// $services->set(BaselineOutputFormatter::class)->tag('output_formatter_contract');
		// $services->set(JsonOutputFormatter::class)->tag('output_formatter_contract');
		// $services->set(GraphVizOutputDisplayFormatter::class)->tag('output_formatter_contract');
		// $services->set(GraphVizOutputImageFormatter::class)->tag('output_formatter_contract');
		// $services->set(GraphVizOutputDotFormatter::class)->tag('output_formatter_contract');
		// $services->set(GraphVizOutputHtmlFormatter::class)->tag('output_formatter_contract');
		// $services->set(CodeclimateOutputFormatter::class)->tag('output_formatter_contract');
		// $services->set(MermaidJSOutputFormatter::class)->tag('output_formatter_contract');
	}
	formatterProvider := formatters.NewFormatterProvider(outputFormatter)
	formatterConfiguration := formatters.NewFormatterConfiguration(builderConfiguration.Formatters)

	//
	knownFormattersStr := make([]string, 0)
	for _, formatterType := range formatterProvider.GetKnownFormatters() {
		knownFormattersStr = append(knownFormattersStr, fmt.Sprintf("\"%s\"", formatterType))
	}
	var (
		formatterUsagePossible = strings.Join(knownFormattersStr, ", ")
		formatterUsage         = fmt.Sprintf("Format in which to print the result_contract of the analysis. Possible: [\"%s\"]", formatterUsagePossible)
		formatter              = flag.String("formatter", string(enums.Table), formatterUsage)
		output                 = flag.String("output", "", "Output file_supportive path for formatter (if applicable)")
		noProgress             = flag.Bool("no-progress", false, "Do not show progress bar")
		reportSkipped          = flag.Bool("report-skipped", false, "Report skipped violations")
		reportUncovered        = flag.Bool("report-uncovered", false, "Report uncovered dependencies")
		failOnUncovered        = flag.Bool("fail-on-uncovered", false, "Fails if any uncovered dependency_contract is found")
	)

	if formatter == nil {
		formatterTmp := string(getDefaultFormatter())
		formatter = &formatterTmp
	}

	analyseOptions := commands_options.NewAnalyseOptions(
		nil != noProgress && *noProgress == true,
		*formatter,
		output,
		nil != reportSkipped && *reportSkipped == true,
		nil != reportUncovered && *reportUncovered == true,
		nil != failOnUncovered && *failOnUncovered == true,
	)
	RegForAnalyseCommand(consoleSubscriber, progressSubscriber, !analyseOptions.NoProgress)
	//

	/*
	 * LayerConfig
	 */
	inheritanceLevelCollector, err := dependencies_collectors.NewInheritanceLevelCollector(astMapExtractor)
	if err != nil {
		return err
	}
	inheritsCollector, err := dependencies_collectors.NewInheritsCollector(astMapExtractor)
	if err != nil {
		return err
	}
	usesCollector, err := dependencies_collectors.NewUsesCollector(astMapExtractor)
	if err != nil {
		return err
	}
	collectorProvider := applicationServices.NewCollectorProvider()
	collectorResolver := collectors_resolvers.NewCollectorResolver(collectorProvider)
	layerResolver := layers_resolvers.NewLayerResolver(collectorResolver, builderConfiguration.Layers)
	collectors := map[enums.CollectorType]domainServices.CollectorInterface{
		//AttributeCollector
		enums.CollectorTypeTypeBool:           dependencies_collectors.NewBoolCollector(collectorResolver),
		enums.CollectorTypeTypeClass:          dependencies_collectors.NewClassCollector(),
		enums.CollectorTypeTypeClasslike:      dependencies_collectors.NewClassLikeCollector(),
		enums.CollectorTypeTypeClassNameRegex: dependencies_collectors.NewClassNameRegexCollector(),
		//CollectorType.TypeTagValueRegex: TagValueRegexCollector.NewTagValueRegexCollector(),
		enums.CollectorTypeTypeDirectory: dependencies_collectors.NewDirectoryCollector(),
		//CollectorType.TypeExtends: ExtendsCollector.NewExtendsCollector(collectorResolver),
		enums.CollectorTypeTypeFunctionName: dependencies_collectors.NewFunctionNameCollector(),
		enums.CollectorTypeTypeGlob:         dependencies_collectors.NewGlobCollector(projectDirectory),
		//ImplementsCollector
		enums.CollectorTypeTypeInheritance: inheritanceLevelCollector,
		enums.CollectorTypeTypeInterface:   dependencies_collectors.NewInterfaceCollector(),
		enums.CollectorTypeTypeInherits:    inheritsCollector,
		enums.CollectorTypeTypeLayer:       dependencies_collectors.NewLayerCollector(layerResolver),
		enums.CollectorTypeTypeMethod:      dependencies_collectors.NewMethodCollector(nikicPhpParser),
		enums.CollectorTypeTypeTrait:       dependencies_collectors.NewTraitCollector(),
		enums.CollectorTypeTypeUses:        usesCollector,
		//CollectorType.TypePhpInternal: PhpInternalCollector
		enums.CollectorTypeTypeComposer: dependencies_collectors.NewComposerCollector(),
	}
	collectorProvider.Set(collectors)

	/*
	 * SetAnalyser
	 */
	dependencyLayersAnalyser := analysers.NewDependencyLayersAnalyser(astMapExtractor, dependencyResolver, tokenResolver, layerResolver, eventDispatcher)
	tokenInLayerAnalyser := analysers.NewTokenInLayerAnalyser(astMapExtractor, tokenResolver, layerResolver, builderConfiguration.Analyser)
	layerForTokenAnalyser := analysers.NewLayerForTokenAnalyser(astMapExtractor, tokenResolver, layerResolver)
	unassignedTokenAnalyser := analysers.NewUnassignedTokenAnalyser(astMapExtractor, tokenResolver, layerResolver, builderConfiguration.Analyser)
	layerDependenciesAnalyser := analysers.NewLayerDependenciesAnalyser(astMapExtractor, tokenResolver, dependencyResolver, layerResolver)
	rulesetUsageAnalyser := analysers.NewRulesetUsageAnalyser(layerProvider, layerResolver, astMapExtractor, dependencyResolver, tokenResolver, builderConfiguration.Layers)

	/*
	 * Console
	 */
	analyseRunner := runners.NewAnalyseRunner(dependencyLayersAnalyser, formatterProvider)
	analyseCommand := commands.NewAnalyseCommand(analyseRunner, eventDispatcher, formatterProvider, *verboseBoolFlag, *debugBoolFlag, consoleSubscriber, progressSubscriber, analyseOptions)

	// TODO: other commands
	// $services->set(InitCommand::class)->autowire()->tag('console_supportive.command');
	// $services->set(ChangedFilesRunner::class)->autowire();
	// $services->set(ChangedFilesCommand::class)->autowire()->tag('console_supportive.command');
	// $services->set(DebugLayerRunner::class)->autowire()->args(['$layers' => param('layers')]);
	// $services->set(DebugLayerCommand::class)->autowire()->tag('console_supportive.command');
	// $services->set(DebugTokenRunner::class)->autowire();
	// $services->set(DebugTokenCommand::class)->autowire()->tag('console_supportive.command');
	// $services->set(DebugUnassignedRunner::class)->autowire();
	// $services->set(DebugUnassignedCommand::class)->autowire()->tag('console_supportive.command');
	// $services->set(DebugDependenciesRunner::class)->autowire();
	// $services->set(DebugDependenciesCommand::class)->autowire()->tag('console_supportive.command');
	// $services->set(DebugUnusedRunner::class)->autowire();
	// $services->set(DebugUnusedCommand::class)->autowire()->tag('console_supportive.command');

	builder.VerboseBoolFlag = verboseBoolFlag
	builder.DebugBoolFlag = debugBoolFlag
	builder.Style = style
	builder.SymfonyOutput = symfonyOutput
	builder.TimeStopwatch = timeStopwatch
	builder.EventDispatcher = eventDispatcher
	builder.FileInputCollector = fileInputCollector
	builder.YmlFileLoader = ymlFileLoader
	builder.Dumper = dumper
	builder.AstFileReferenceInMemoryCache = astFileReferenceInMemoryCache
	builder.TypeResolver = typeResolver
	builder.ReferenceExtractors = referenceExtractors
	builder.NikicPhpParser = nikicPhpParser
	builder.ParserInterface = parserInterface
	builder.AstLoader = astLoader
	builder.InheritanceFlattener = inheritanceFlattener
	builder.DependencyResolver = dependencyResolver
	builder.TokenResolver = tokenResolver
	builder.AstMapExtractor = astMapExtractor
	builder.CollectorResolver = collectorResolver
	builder.LayerResolver = layerResolver
	builder.CollectorProvider = collectorProvider
	builder.UncoveredDependentHandler = uncoveredDependentHandler
	builder.MatchingLayersHandler = matchingLayersHandler
	builder.LayerProvider = layerProvider
	builder.AllowDependencyHandler = allowDependencyHandler
	builder.DependsOnDisallowedLayer = dependsOnDisallowedLayer
	builder.EventHelper = eventHelper
	builder.DependsOnPrivateLayer = dependsOnPrivateLayer
	builder.DependsOnInternalToken = dependsOnInternalToken
	builder.UnmatchedSkippedViolations = unmatchedSkippedViolations
	builder.DependencyLayersAnalyser = dependencyLayersAnalyser
	builder.TokenInLayerAnalyser = tokenInLayerAnalyser
	builder.LayerForTokenAnalyser = layerForTokenAnalyser
	builder.UnassignedTokenAnalyser = unassignedTokenAnalyser
	builder.LayerDependenciesAnalyser = layerDependenciesAnalyser
	builder.RulesetUsageAnalyser = rulesetUsageAnalyser
	builder.ConsoleSubscriber = consoleSubscriber
	builder.ProgressSubscriber = progressSubscriber
	builder.FormatterProvider = formatterProvider
	builder.FormatterConfiguration = formatterConfiguration
	builder.AnalyseRunner = analyseRunner
	builder.AnalyseCommand = analyseCommand
	builder.NodeNamer = nodeNamer
	builder.AnalyseOptions = analyseOptions

	return nil
}

func RegForAnalyseCommand(consoleSubscriber *event_handlers.Console, progressSubscriber *event_handlers.Progress, withProgress bool) {
	processEvent := &events.ProcessEvent{}
	postProcessEvent := &events.PostProcessEvent{}
	preCreateAstMapEvent := &events.PreCreateAstMapEvent{}
	postCreateAstMapEvent := &events.PostCreateAstMapEvent{}
	astFileAnalysedEvent := &events.AstFileAnalysedEvent{}
	astFileSyntaxErrorEvent := &events.AstFileSyntaxErrorEvent{}
	preEmitEvent := &events.PreEmitEvent{}
	postEmitEvent := &events.PostEmitEvent{}
	preFlattenEvent := &events.PreFlattenEvent{}
	postFlattenEvent := &events.PostFlattenEvent{}

	event_dispatchers.Reg(preCreateAstMapEvent, consoleSubscriber, event_dispatchers.DefaultPriority)
	event_dispatchers.Reg(postCreateAstMapEvent, consoleSubscriber, event_dispatchers.DefaultPriority)
	event_dispatchers.Reg(processEvent, consoleSubscriber, event_dispatchers.DefaultPriority)
	event_dispatchers.Reg(postProcessEvent, consoleSubscriber, event_dispatchers.DefaultPriority)
	event_dispatchers.Reg(astFileAnalysedEvent, consoleSubscriber, event_dispatchers.DefaultPriority)
	event_dispatchers.Reg(astFileSyntaxErrorEvent, consoleSubscriber, event_dispatchers.DefaultPriority)
	event_dispatchers.Reg(preEmitEvent, consoleSubscriber, event_dispatchers.DefaultPriority)
	event_dispatchers.Reg(postEmitEvent, consoleSubscriber, event_dispatchers.DefaultPriority)
	event_dispatchers.Reg(preFlattenEvent, consoleSubscriber, event_dispatchers.DefaultPriority)
	event_dispatchers.Reg(postFlattenEvent, consoleSubscriber, event_dispatchers.DefaultPriority)

	if withProgress {
		event_dispatchers.Reg(preCreateAstMapEvent, progressSubscriber, event_dispatchers.DefaultPriority)
		event_dispatchers.Reg(postCreateAstMapEvent, progressSubscriber, 1)
		event_dispatchers.Reg(astFileAnalysedEvent, progressSubscriber, event_dispatchers.DefaultPriority)
	}
}

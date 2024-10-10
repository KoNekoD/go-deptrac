package app

import (
	"flag"
	"fmt"
	"github.com/KoNekoD/go-deptrac/pkg"
	"github.com/KoNekoD/go-deptrac/pkg/analysers"
	event_handlers2 "github.com/KoNekoD/go-deptrac/pkg/application/event_handlers"
	services2 "github.com/KoNekoD/go-deptrac/pkg/application/services"
	"github.com/KoNekoD/go-deptrac/pkg/application/services/ast_file_reference_cache"
	ast_map2 "github.com/KoNekoD/go-deptrac/pkg/application/services/ast_map"
	"github.com/KoNekoD/go-deptrac/pkg/application/services/collectors_resolvers"
	"github.com/KoNekoD/go-deptrac/pkg/application/services/dependencies_collectors"
	"github.com/KoNekoD/go-deptrac/pkg/application/services/emitters"
	"github.com/KoNekoD/go-deptrac/pkg/application/services/input_collectors"
	"github.com/KoNekoD/go-deptrac/pkg/application/services/layers_resolvers"
	"github.com/KoNekoD/go-deptrac/pkg/application/services/parsers"
	"github.com/KoNekoD/go-deptrac/pkg/application/services/types"
	"github.com/KoNekoD/go-deptrac/pkg/commands"
	"github.com/KoNekoD/go-deptrac/pkg/dispatchers"
	enums2 "github.com/KoNekoD/go-deptrac/pkg/domain/enums"
	events2 "github.com/KoNekoD/go-deptrac/pkg/domain/events"
	"github.com/KoNekoD/go-deptrac/pkg/domain/services"
	"github.com/KoNekoD/go-deptrac/pkg/domain/stopwatch"
	"github.com/KoNekoD/go-deptrac/pkg/domain/utils"
	"github.com/KoNekoD/go-deptrac/pkg/extractors"
	"github.com/KoNekoD/go-deptrac/pkg/formatters"
	"github.com/elliotchance/orderedmap/v2"
	"os"
	"strings"
)

func getDefaultFormatter() enums2.OutputFormatterType {
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
	symfonyOutput := pkg.NewSymfonyOutput(style)

	timeStopwatch := stopwatch.NewStopwatch()

	nodeNamer := services.NewNodeNamer(projectDirectory)

	/*
	 * Utilities
	 */
	eventDispatcher := dispatchers.NewEventDispatcher(debugBoolFlag != nil && *debugBoolFlag == true)

	fileInputCollector, err := input_collectors.NewFileInputCollector(
		builderConfiguration.Paths,
		builderConfiguration.ExcludeFiles,
		projectDirectory,
	)
	if err != nil {
		return err
	}

	ymlFileLoader := pkg.NewYmlFileLoader()
	dumper := utils.NewDumper("/deptrac_template.yaml")

	/*
	 * AST
	 */
	astFileReferenceInMemoryCache := ast_file_reference_cache.NewAstFileReferenceInMemoryCache()
	if builder.AstFileReferenceCacheInterface == nil {
		builder.AstFileReferenceCacheInterface = astFileReferenceInMemoryCache
	}
	typeResolver := types.NewTypeResolver(nodeNamer)
	referenceExtractors := []extractors.ReferenceExtractorInterface{
		/**

		TODO: Implement all reference extractors

		AnnotationReferenceExtractor.NewAnnotationReferenceExtractor(),
		AnonymousClassExtractor.NewAnonymousClassExtractor(),
		ClassConstantExtractor.NewClassConstantExtractor(),
		FunctionLikeExtractor.NewFunctionLikeExtractor(),
		PropertyExtractor.NewPropertyExtractor(),
		KeywordExtractor.NewKeywordExtractor(),
		StaticExtractor.NewStaticExtractor(),
		FunctionCallResolver.NewFunctionCallResolver(),

		*/
	}
	nikicPhpParser := parsers.NewNikicPhpParser(builder.AstFileReferenceCacheInterface, typeResolver, nodeNamer, referenceExtractors)
	parserInterface := nikicPhpParser
	astLoader := ast_map2.NewAstLoader(parserInterface, eventDispatcher)

	/*
	 * Dependency
	 */
	dependencyEmitters := map[enums2.EmitterType]emitters.DependencyEmitterInterface{
		enums2.EmitterTypeClassToken:               emitters.NewClassDependencyEmitter(),
		enums2.EmitterTypeClassSuperGlobalToken:    emitters.NewClassSuperglobalDependencyEmitter(),
		enums2.EmitterTypeFileToken:                emitters.NewFileDependencyEmitter(),
		enums2.EmitterTypeFunctionToken:            emitters.NewFunctionDependencyEmitter(),
		enums2.EmitterTypeFunctionCall:             emitters.NewFunctionCallDependencyEmitter(),
		enums2.EmitterTypeFunctionSuperGlobalToken: emitters.NewFunctionSuperglobalDependencyEmitter(),
		enums2.EmitterTypeUseToken:                 emitters.NewUsesDependencyEmitter(),
	}
	inheritanceFlattener := services2.NewInheritanceFlattener()
	dependencyResolver := pkg.NewDependencyResolver(builderConfiguration.Analyser, dependencyEmitters, inheritanceFlattener, eventDispatcher)
	tokenResolver := services2.NewTokenResolver()

	astMapExtractor := ast_map2.NewAstMapExtractor(fileInputCollector, astLoader)

	layerProvider := services2.NewLayerProvider(builderConfiguration.Rulesets)
	eventHelper := dispatchers.NewEventHelper(builderConfiguration.SkipViolations, layerProvider)

	/*
	 * Events (before first possible event)
	 */
	/*
	 * Events
	 */
	event_handlers2.Map = orderedmap.NewOrderedMap[string, *orderedmap.OrderedMap[int, []event_handlers2.EventHandlerInterface]]()

	// Events
	uncoveredDependentHandler := event_handlers2.NewUncoveredDependent(builderConfiguration.IgnoreUncoveredInternalStructs)
	matchingLayersHandler := event_handlers2.NewMatchingLayers()
	allowDependencyHandler := event_handlers2.NewAllowDependency()
	consoleSubscriber := event_handlers2.NewConsole(symfonyOutput, timeStopwatch)
	progressSubscriber := event_handlers2.NewProgress(symfonyOutput)
	dependsOnDisallowedLayer := event_handlers2.NewDependsOnDisallowedLayer(eventHelper)
	dependsOnPrivateLayer := event_handlers2.NewDependsOnPrivateLayer(eventHelper)
	dependsOnInternalToken := event_handlers2.NewDependsOnInternalToken(eventHelper, builderConfiguration.Analyser)
	unmatchedSkippedViolations := event_handlers2.NewUnmatchedSkippedViolations(eventHelper)

	processEvent := &events2.ProcessEvent{}
	postProcessEvent := &events2.PostProcessEvent{}
	preCreateAstMapEvent := &events2.PreCreateAstMapEvent{}
	postCreateAstMapEvent := &events2.PostCreateAstMapEvent{}
	// Events Handlers
	// TODO: Тут надо реализовать глобальный хук на параметры deptrac чтобы сделать что-то вида "param('skip_violations')"
	event_handlers2.Reg(processEvent, allowDependencyHandler, -100)
	event_handlers2.Reg(processEvent, dependsOnPrivateLayer, -3)
	event_handlers2.Reg(processEvent, dependsOnInternalToken, -2)
	event_handlers2.Reg(processEvent, dependsOnDisallowedLayer, -1)
	event_handlers2.Reg(processEvent, matchingLayersHandler, 1)
	event_handlers2.Reg(processEvent, uncoveredDependentHandler, 2)
	event_handlers2.Reg(postProcessEvent, unmatchedSkippedViolations, event_handlers2.DefaultPriority)
	if cacheableFileSubscriber != nil {
		event_handlers2.Reg(preCreateAstMapEvent, cacheableFileSubscriber, event_handlers2.DefaultPriority)
		event_handlers2.Reg(postCreateAstMapEvent, cacheableFileSubscriber, event_handlers2.DefaultPriority)
	}

	/*
	 * OutputFormatter
	 */
	outputFormatter := map[enums2.OutputFormatterType]formatters.OutputFormatterInterface{
		enums2.Table:         formatters.NewTableOutputFormatter(),
		enums2.GithubActions: formatters.NewGithubActionsOutputFormatter(),
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
		formatter              = flag.String("formatter", string(enums2.Table), formatterUsage)
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

	analyseOptions := pkg.NewAnalyseOptions(
		nil != noProgress && *noProgress == true,
		*formatter,
		output,
		nil != reportSkipped && *reportSkipped == true,
		nil != reportUncovered && *reportUncovered == true,
		nil != failOnUncovered && *failOnUncovered == true,
	)
	event_handlers2.RegForAnalyseCommand(consoleSubscriber, progressSubscriber, !analyseOptions.NoProgress)
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
	collectorProvider := services2.NewCollectorProvider()
	collectorResolver := collectors_resolvers.NewCollectorResolver(collectorProvider)
	layerResolver := layers_resolvers.NewLayerResolver(collectorResolver, builderConfiguration.Layers)
	collectors := map[enums2.CollectorType]dependencies_collectors.CollectorInterface{
		//AttributeCollector
		enums2.CollectorTypeTypeBool:           dependencies_collectors.NewBoolCollector(collectorResolver),
		enums2.CollectorTypeTypeClass:          dependencies_collectors.NewClassCollector(),
		enums2.CollectorTypeTypeClasslike:      dependencies_collectors.NewClassLikeCollector(),
		enums2.CollectorTypeTypeClassNameRegex: dependencies_collectors.NewClassNameRegexCollector(),
		//CollectorType.TypeTagValueRegex: TagValueRegexCollector.NewTagValueRegexCollector(),
		enums2.CollectorTypeTypeDirectory: dependencies_collectors.NewDirectoryCollector(),
		//CollectorType.TypeExtends: ExtendsCollector.NewExtendsCollector(collectorResolver),
		enums2.CollectorTypeTypeFunctionName: dependencies_collectors.NewFunctionNameCollector(),
		enums2.CollectorTypeTypeGlob:         dependencies_collectors.NewGlobCollector(projectDirectory),
		//ImplementsCollector
		enums2.CollectorTypeTypeInheritance: inheritanceLevelCollector,
		enums2.CollectorTypeTypeInterface:   dependencies_collectors.NewInterfaceCollector(),
		enums2.CollectorTypeTypeInherits:    inheritsCollector,
		enums2.CollectorTypeTypeLayer:       dependencies_collectors.NewLayerCollector(layerResolver),
		enums2.CollectorTypeTypeMethod:      dependencies_collectors.NewMethodCollector(nikicPhpParser),
		enums2.CollectorTypeTypeSuperGlobal: dependencies_collectors.NewSuperglobalCollector(),
		enums2.CollectorTypeTypeTrait:       dependencies_collectors.NewTraitCollector(),
		enums2.CollectorTypeTypeUses:        usesCollector,
		//CollectorType.TypePhpInternal: PhpInternalCollector
		enums2.CollectorTypeTypeComposer: dependencies_collectors.NewComposerCollector(),
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
	analyseRunner := NewAnalyseRunner(dependencyLayersAnalyser, formatterProvider)
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

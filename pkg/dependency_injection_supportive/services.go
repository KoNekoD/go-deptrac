package dependency_injection_supportive

import (
	"flag"
	"fmt"
	"github.com/KoNekoD/go-deptrac/pkg/analyser_contract/event_helper"
	post_process_event2 "github.com/KoNekoD/go-deptrac/pkg/analyser_contract/post_process_event"
	process_event2 "github.com/KoNekoD/go-deptrac/pkg/analyser_contract/process_event"
	analyser_core2 "github.com/KoNekoD/go-deptrac/pkg/analyser_core"
	"github.com/KoNekoD/go-deptrac/pkg/analyser_core/event_handler/post_process_event"
	process_event3 "github.com/KoNekoD/go-deptrac/pkg/analyser_core/event_handler/process_event"
	"github.com/KoNekoD/go-deptrac/pkg/ast_contract"
	ast_core2 "github.com/KoNekoD/go-deptrac/pkg/ast_core"
	"github.com/KoNekoD/go-deptrac/pkg/ast_core/parser"
	Cache2 "github.com/KoNekoD/go-deptrac/pkg/ast_core/parser/cache"
	"github.com/KoNekoD/go-deptrac/pkg/ast_core/parser/extractors"
	"github.com/KoNekoD/go-deptrac/pkg/ast_core/parser/nikic_php_parser"
	"github.com/KoNekoD/go-deptrac/pkg/ast_core/parser/nikic_php_parser/node_namer"
	config_contract2 "github.com/KoNekoD/go-deptrac/pkg/config_contract"
	"github.com/KoNekoD/go-deptrac/pkg/console_supportive"
	command2 "github.com/KoNekoD/go-deptrac/pkg/console_supportive/command"
	subscriber2 "github.com/KoNekoD/go-deptrac/pkg/console_supportive/subscriber"
	symfony2 "github.com/KoNekoD/go-deptrac/pkg/console_supportive/symfony"
	dependency_core2 "github.com/KoNekoD/go-deptrac/pkg/dependency_core"
	"github.com/KoNekoD/go-deptrac/pkg/dependency_core/dependency_resolver"
	emitter2 "github.com/KoNekoD/go-deptrac/pkg/dependency_core/emitter"
	"github.com/KoNekoD/go-deptrac/pkg/dependency_injection_supportive/container_builder"
	"github.com/KoNekoD/go-deptrac/pkg/dependency_injection_supportive/event_dispatcher"
	EventSubscriberInterfaceMap2 "github.com/KoNekoD/go-deptrac/pkg/dependency_injection_supportive/event_subscriber_default_priority"
	EventSubscriberInterface2 "github.com/KoNekoD/go-deptrac/pkg/dependency_injection_supportive/event_subscriber_interface"
	"github.com/KoNekoD/go-deptrac/pkg/dependency_injection_supportive/event_subscriber_interface_map"
	"github.com/KoNekoD/go-deptrac/pkg/dependency_injection_supportive/event_subscriber_interface_map/event_subscriber_interface_map_reg"
	file_supportive2 "github.com/KoNekoD/go-deptrac/pkg/file_supportive"
	"github.com/KoNekoD/go-deptrac/pkg/input_collector_core"
	"github.com/KoNekoD/go-deptrac/pkg/layer_contract"
	"github.com/KoNekoD/go-deptrac/pkg/layer_core"
	collector2 "github.com/KoNekoD/go-deptrac/pkg/layer_core/collector"
	output_formatter_contract2 "github.com/KoNekoD/go-deptrac/pkg/output_formatter_contract"
	"github.com/KoNekoD/go-deptrac/pkg/output_formatter_supportive"
	"github.com/KoNekoD/go-deptrac/pkg/output_formatter_supportive/configuration"
	"github.com/KoNekoD/go-deptrac/pkg/time_stopwatch_supportive"
	"github.com/elliotchance/orderedmap/v2"
	"strings"
)

func getDefaultFormatter() output_formatter_contract2.OutputFormatterType {
	if console_supportive.NewEnv().GetEnv("GITHUB_ACTIONS") != "" {
		return output_formatter_supportive.NewGithubActionsOutputFormatter().GetName()
	}
	return output_formatter_supportive.NewTableOutputFormatter().GetName()
}

func Services(builder *container_builder.ContainerBuilder) error {

	cacheableFileSubscriber := builder.CacheableFileSubscriber
	builderConfiguration := builder.Configuration
	projectDirectory := builder.ProjectDirectory
	verboseBoolFlag := flag.Bool("verbose", true, "Verbose mode")
	debugBoolFlag := flag.Bool("debug", false, "Debug mode")
	style := symfony2.NewStyle(
		verboseBoolFlag != nil && *verboseBoolFlag == true,
		debugBoolFlag != nil && *debugBoolFlag == true,
	)
	symfonyOutput := symfony2.NewSymfonyOutput(style)

	timeStopwatch := time_stopwatch_supportive.NewStopwatch()

	nodeNamer := node_namer.NewNodeNamer(projectDirectory)

	/*
	 * Utilities
	 */
	eventDispatcher := event_dispatcher.NewEventDispatcher(debugBoolFlag != nil && *debugBoolFlag == true)

	fileInputCollector, err := input_collector_core.NewFileInputCollector(
		builderConfiguration.Paths,
		builderConfiguration.ExcludeFiles,
		projectDirectory,
	)
	if err != nil {
		return err
	}

	ymlFileLoader := file_supportive2.NewYmlFileLoader()
	dumper := file_supportive2.NewDumper("/deptrac_template.yaml")

	/*
	 * AST
	 */
	astFileReferenceInMemoryCache := Cache2.NewAstFileReferenceInMemoryCache()
	if builder.AstFileReferenceCacheInterface == nil {
		builder.AstFileReferenceCacheInterface = astFileReferenceInMemoryCache
	}
	typeResolver := parser.NewTypeResolver(nodeNamer)
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
	nikicPhpParser := nikic_php_parser.NewNikicPhpParser(builder.AstFileReferenceCacheInterface, typeResolver, nodeNamer, referenceExtractors)
	parserInterface := nikicPhpParser
	astLoader := ast_core2.NewAstLoader(parserInterface, eventDispatcher)

	/*
	 * Dependency
	 */
	dependencyEmitters := map[config_contract2.EmitterType]emitter2.DependencyEmitterInterface{
		config_contract2.ClassToken:               emitter2.NewClassDependencyEmitter(),
		config_contract2.ClassSuperGlobalToken:    emitter2.NewClassSuperglobalDependencyEmitter(),
		config_contract2.FileToken:                emitter2.NewFileDependencyEmitter(),
		config_contract2.FunctionToken:            emitter2.NewFunctionDependencyEmitter(),
		config_contract2.FunctionCall:             emitter2.NewFunctionCallDependencyEmitter(),
		config_contract2.FunctionSuperGlobalToken: emitter2.NewFunctionSuperglobalDependencyEmitter(),
		config_contract2.UseToken:                 emitter2.NewUsesDependencyEmitter(),
	}
	inheritanceFlattener := dependency_core2.NewInheritanceFlattener()
	dependencyResolver := dependency_resolver.NewDependencyResolver(builderConfiguration.Analyser, dependencyEmitters, inheritanceFlattener, eventDispatcher)
	tokenResolver := dependency_core2.NewTokenResolver()

	astMapExtractor := ast_core2.NewAstMapExtractor(fileInputCollector, astLoader)

	layerProvider := layer_contract.NewLayerProvider(builderConfiguration.Rulesets)
	eventHelper := event_helper.NewEventHelper(builderConfiguration.SkipViolations, layerProvider)

	/*
	 * Events (before first possible event)
	 */
	/*
	 * Events
	 */
	event_subscriber_interface_map.Map = orderedmap.NewOrderedMap[string, *orderedmap.OrderedMap[int, []EventSubscriberInterface2.EventSubscriberInterface]]()

	// Events
	uncoveredDependentHandler := process_event3.NewUncoveredDependentHandler(builderConfiguration.IgnoreUncoveredInternalStructs)
	matchingLayersHandler := process_event3.NewMatchingLayersHandler()
	allowDependencyHandler := process_event3.NewAllowDependencyHandler()
	consoleSubscriber := subscriber2.NewConsoleSubscriber(symfonyOutput, timeStopwatch)
	progressSubscriber := subscriber2.NewProgressSubscriber(symfonyOutput)
	dependsOnDisallowedLayer := process_event3.NewDependsOnDisallowedLayer(eventHelper)
	dependsOnPrivateLayer := process_event3.NewDependsOnPrivateLayer(eventHelper)
	dependsOnInternalToken := process_event3.NewDependsOnInternalToken(eventHelper, builderConfiguration.Analyser)
	unmatchedSkippedViolations := post_process_event.NewUnmatchedSkippedViolations(eventHelper)

	processEvent := &process_event2.ProcessEvent{}
	postProcessEvent := &post_process_event2.PostProcessEvent{}
	preCreateAstMapEvent := &ast_contract.PreCreateAstMapEvent{}
	postCreateAstMapEvent := &ast_contract.PostCreateAstMapEvent{}
	// Events Handlers
	// TODO: Тут надо реализовать глобальный хук на параметры deptrac чтобы сделать что-то вида "param('skip_violations')"
	event_subscriber_interface_map_reg.Reg(processEvent, allowDependencyHandler, -100)
	event_subscriber_interface_map_reg.Reg(processEvent, dependsOnPrivateLayer, -3)
	event_subscriber_interface_map_reg.Reg(processEvent, dependsOnInternalToken, -2)
	event_subscriber_interface_map_reg.Reg(processEvent, dependsOnDisallowedLayer, -1)
	event_subscriber_interface_map_reg.Reg(processEvent, matchingLayersHandler, 1)
	event_subscriber_interface_map_reg.Reg(processEvent, uncoveredDependentHandler, 2)
	event_subscriber_interface_map_reg.Reg(postProcessEvent, unmatchedSkippedViolations, EventSubscriberInterfaceMap2.DefaultPriority)
	if cacheableFileSubscriber != nil {
		event_subscriber_interface_map_reg.Reg(preCreateAstMapEvent, cacheableFileSubscriber, EventSubscriberInterfaceMap2.DefaultPriority)
		event_subscriber_interface_map_reg.Reg(postCreateAstMapEvent, cacheableFileSubscriber, EventSubscriberInterfaceMap2.DefaultPriority)
	}

	/*
	 * OutputFormatter
	 */
	outputFormatter := map[output_formatter_contract2.OutputFormatterType]output_formatter_contract2.OutputFormatterInterface{
		output_formatter_contract2.Table:         output_formatter_supportive.NewTableOutputFormatter(),
		output_formatter_contract2.GithubActions: output_formatter_supportive.NewGithubActionsOutputFormatter(),
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
	formatterProvider := output_formatter_supportive.NewFormatterProvider(outputFormatter)
	formatterConfiguration := configuration.NewFormatterConfiguration(builderConfiguration.Formatters)

	//
	knownFormattersStr := make([]string, 0)
	for _, formatterType := range formatterProvider.GetKnownFormatters() {
		knownFormattersStr = append(knownFormattersStr, fmt.Sprintf("\"%s\"", formatterType))
	}
	var (
		formatterUsagePossible = strings.Join(knownFormattersStr, ", ")
		formatterUsage         = fmt.Sprintf("Format in which to print the result_contract of the analysis. Possible: [\"%s\"]", formatterUsagePossible)
		formatter              = flag.String("formatter", string(output_formatter_contract2.Table), formatterUsage)
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

	analyseOptions := command2.NewAnalyseOptions(
		nil != noProgress && *noProgress == true,
		*formatter,
		output,
		nil != reportSkipped && *reportSkipped == true,
		nil != reportUncovered && *reportUncovered == true,
		nil != failOnUncovered && *failOnUncovered == true,
	)
	event_subscriber_interface_map_reg.RegForAnalyseCommand(consoleSubscriber, progressSubscriber, !analyseOptions.NoProgress)
	//

	/*
	 * Layer
	 */
	inheritanceLevelCollector, err := collector2.NewInheritanceLevelCollector(astMapExtractor)
	if err != nil {
		return err
	}
	inheritsCollector, err := collector2.NewInheritsCollector(astMapExtractor)
	if err != nil {
		return err
	}
	usesCollector, err := collector2.NewUsesCollector(astMapExtractor)
	if err != nil {
		return err
	}
	collectorProvider := collector2.NewCollectorProvider()
	collectorResolver := collector2.NewCollectorResolver(collectorProvider)
	layerResolver := layer_core.NewLayerResolver(collectorResolver, builderConfiguration.Layers)
	collectors := map[config_contract2.CollectorType]layer_contract.CollectorInterface{
		//AttributeCollector
		config_contract2.TypeBool:           collector2.NewBoolCollector(collectorResolver),
		config_contract2.TypeClass:          collector2.NewClassCollector(),
		config_contract2.TypeClasslike:      collector2.NewClassLikeCollector(),
		config_contract2.TypeClassNameRegex: collector2.NewClassNameRegexCollector(),
		//CollectorType.TypeTagValueRegex: TagValueRegexCollector.NewTagValueRegexCollector(),
		config_contract2.TypeDirectory: collector2.NewDirectoryCollector(),
		//CollectorType.TypeExtends: ExtendsCollector.NewExtendsCollector(collectorResolver),
		config_contract2.TypeFunctionName: collector2.NewFunctionNameCollector(),
		config_contract2.TypeGlob:         collector2.NewGlobCollector(projectDirectory),
		//ImplementsCollector
		config_contract2.TypeInheritance: inheritanceLevelCollector,
		config_contract2.TypeInterface:   collector2.NewInterfaceCollector(),
		config_contract2.TypeInherits:    inheritsCollector,
		config_contract2.TypeLayer:       collector2.NewLayerCollector(layerResolver),
		config_contract2.TypeMethod:      collector2.NewMethodCollector(nikicPhpParser),
		config_contract2.TypeSuperGlobal: collector2.NewSuperglobalCollector(),
		config_contract2.TypeTrait:       collector2.NewTraitCollector(),
		config_contract2.TypeUses:        usesCollector,
		//CollectorType.TypePhpInternal: PhpInternalCollector
		config_contract2.TypeComposer: collector2.NewComposerCollector(),
	}
	collectorProvider.Set(collectors)

	/*
	 * SetAnalyser
	 */
	dependencyLayersAnalyser := analyser_core2.NewDependencyLayersAnalyser(astMapExtractor, dependencyResolver, tokenResolver, layerResolver, eventDispatcher)
	tokenInLayerAnalyser := analyser_core2.NewTokenInLayerAnalyser(astMapExtractor, tokenResolver, layerResolver, builderConfiguration.Analyser)
	layerForTokenAnalyser := analyser_core2.NewLayerForTokenAnalyser(astMapExtractor, tokenResolver, layerResolver)
	unassignedTokenAnalyser := analyser_core2.NewUnassignedTokenAnalyser(astMapExtractor, tokenResolver, layerResolver, builderConfiguration.Analyser)
	layerDependenciesAnalyser := analyser_core2.NewLayerDependenciesAnalyser(astMapExtractor, tokenResolver, dependencyResolver, layerResolver)
	rulesetUsageAnalyser := analyser_core2.NewRulesetUsageAnalyser(layerProvider, layerResolver, astMapExtractor, dependencyResolver, tokenResolver, builderConfiguration.Layers)

	/*
	 * Console
	 */
	analyseRunner := command2.NewAnalyseRunner(dependencyLayersAnalyser, formatterProvider)
	analyseCommand := command2.NewAnalyseCommand(analyseRunner, eventDispatcher, formatterProvider, *verboseBoolFlag, *debugBoolFlag, consoleSubscriber, progressSubscriber, analyseOptions)

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

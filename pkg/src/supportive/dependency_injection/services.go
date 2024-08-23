package dependency_injection

import (
	"flag"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/analyser/event_helper"
	post_process_event2 "github.com/KoNekoD/go-deptrac/pkg/src/contract/analyser/post_process_event"
	process_event2 "github.com/KoNekoD/go-deptrac/pkg/src/contract/analyser/process_event"
	astContract "github.com/KoNekoD/go-deptrac/pkg/src/contract/ast"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/config"
	contractLayer "github.com/KoNekoD/go-deptrac/pkg/src/contract/layer"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/output_formatter"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/analyser"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/analyser/event_handler/post_process_event"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/analyser/event_handler/process_event"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/ast"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/ast/parser"
	Cache2 "github.com/KoNekoD/go-deptrac/pkg/src/core/ast/parser/cache"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/ast/parser/extractors"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/ast/parser/nikic_php_parser"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/dependency"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/dependency/dependency_resolver"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/dependency/emitter"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/input_collector"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/layer"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/layer/collector"
	"github.com/KoNekoD/go-deptrac/pkg/src/supportive/console/command"
	"github.com/KoNekoD/go-deptrac/pkg/src/supportive/console/subscriber"
	"github.com/KoNekoD/go-deptrac/pkg/src/supportive/console/symfony"
	"github.com/KoNekoD/go-deptrac/pkg/src/supportive/dependency_injection/container_builder"
	"github.com/KoNekoD/go-deptrac/pkg/src/supportive/dependency_injection/event_dispatcher"
	EventSubscriberInterfaceMap2 "github.com/KoNekoD/go-deptrac/pkg/src/supportive/dependency_injection/event_subscriber_default_priority"
	EventSubscriberInterface2 "github.com/KoNekoD/go-deptrac/pkg/src/supportive/dependency_injection/event_subscriber_interface"
	"github.com/KoNekoD/go-deptrac/pkg/src/supportive/dependency_injection/event_subscriber_interface_map"
	"github.com/KoNekoD/go-deptrac/pkg/src/supportive/dependency_injection/event_subscriber_interface_map/event_subscriber_interface_map_reg"
	"github.com/KoNekoD/go-deptrac/pkg/src/supportive/file"
	output_formatter2 "github.com/KoNekoD/go-deptrac/pkg/src/supportive/output_formatter"
	"github.com/KoNekoD/go-deptrac/pkg/src/supportive/output_formatter/configuration"
	"github.com/KoNekoD/go-deptrac/pkg/src/supportive/time_stopwatch"
	"github.com/elliotchance/orderedmap/v2"
)

func Services(builder *container_builder.ContainerBuilder) error {
	cacheableFileSubscriber := builder.CacheableFileSubscriber
	builderConfiguration := builder.Configuration
	projectDirectory := builder.ProjectDirectory
	verboseBoolFlag := flag.Bool("verbose", false, "Verbose mode")
	debugBoolFlag := flag.Bool("debug", false, "Debug mode")
	style := symfony.NewStyle(
		verboseBoolFlag != nil && *verboseBoolFlag == true,
		debugBoolFlag != nil && *debugBoolFlag == true,
	)
	symfonyOutput := symfony.NewSymfonyOutput(style)

	timeStopwatch := time_stopwatch.NewStopwatch()

	/*
	 * Utilities
	 */
	eventDispatcher := event_dispatcher.NewEventDispatcher(debugBoolFlag != nil && *debugBoolFlag == true)

	fileInputCollector, err := input_collector.NewFileInputCollector(
		builderConfiguration.Paths,
		builderConfiguration.ExcludeFiles,
		projectDirectory,
	)
	if err != nil {
		return err
	}

	ymlFileLoader := file.NewYmlFileLoader()
	dumper := file.NewDumper("/deptrac_template.yaml")

	/*
	 * AST
	 */
	astFileReferenceInMemoryCache := Cache2.NewAstFileReferenceInMemoryCache()
	if builder.AstFileReferenceCacheInterface == nil {
		builder.AstFileReferenceCacheInterface = astFileReferenceInMemoryCache
	}
	typeResolver := parser.NewTypeResolver()
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
	nikicPhpParser := nikic_php_parser.NewNikicPhpParser(builder.AstFileReferenceCacheInterface, typeResolver, referenceExtractors)
	parserInterface := nikicPhpParser
	astLoader := ast.NewAstLoader(parserInterface, eventDispatcher)

	/*
	 * Dependency
	 */
	dependencyEmitters := map[config.EmitterType]emitter.DependencyEmitterInterface{
		config.ClassToken:               emitter.NewClassDependencyEmitter(),
		config.ClassSuperGlobalToken:    emitter.NewClassSuperglobalDependencyEmitter(),
		config.FileToken:                emitter.NewFileDependencyEmitter(),
		config.FunctionToken:            emitter.NewFunctionDependencyEmitter(),
		config.FunctionCall:             emitter.NewFunctionCallDependencyEmitter(),
		config.FunctionSuperGlobalToken: emitter.NewFunctionSuperglobalDependencyEmitter(),
		config.UseToken:                 emitter.NewUsesDependencyEmitter(),
	}
	inheritanceFlattener := dependency.NewInheritanceFlattener()
	dependencyResolver := dependency_resolver.NewDependencyResolver(builderConfiguration.Analyser, dependencyEmitters, inheritanceFlattener, eventDispatcher)
	tokenResolver := dependency.NewTokenResolver()

	astMapExtractor := ast.NewAstMapExtractor(fileInputCollector, astLoader)

	layerProvider := contractLayer.NewLayerProvider(builderConfiguration.Rulesets)
	eventHelper := event_helper.NewEventHelper(builderConfiguration.SkipViolations, layerProvider)

	/*
	 * Events (before first possible event)
	 */
	/*
	 * Events
	 */
	event_subscriber_interface_map.Map = orderedmap.NewOrderedMap[string, *orderedmap.OrderedMap[int, []EventSubscriberInterface2.EventSubscriberInterface]]()

	// Events
	uncoveredDependentHandler := process_event.NewUncoveredDependentHandler(builderConfiguration.IgnoreUncoveredInternalStructs)
	matchingLayersHandler := process_event.NewMatchingLayersHandler()
	allowDependencyHandler := process_event.NewAllowDependencyHandler()
	consoleSubscriber := subscriber.NewConsoleSubscriber(symfonyOutput, timeStopwatch)
	progressSubscriber := subscriber.NewProgressSubscriber(symfonyOutput)
	dependsOnDisallowedLayer := process_event.NewDependsOnDisallowedLayer(eventHelper)
	dependsOnPrivateLayer := process_event.NewDependsOnPrivateLayer(eventHelper)
	dependsOnInternalToken := process_event.NewDependsOnInternalToken(eventHelper, builderConfiguration.Analyser)
	unmatchedSkippedViolations := post_process_event.NewUnmatchedSkippedViolations(eventHelper)

	processEvent := &process_event2.ProcessEvent{}
	postProcessEvent := &post_process_event2.PostProcessEvent{}
	preCreateAstMapEvent := &astContract.PreCreateAstMapEvent{}
	postCreateAstMapEvent := &astContract.PostCreateAstMapEvent{}
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
	 * Layer
	 */
	inheritanceLevelCollector, err := collector.NewInheritanceLevelCollector(astMapExtractor)
	if err != nil {
		return err
	}
	inheritsCollector, err := collector.NewInheritsCollector(astMapExtractor)
	if err != nil {
		return err
	}
	usesCollector, err := collector.NewUsesCollector(astMapExtractor)
	if err != nil {
		return err
	}
	collectorProvider := collector.NewCollectorProvider()
	collectorResolver := collector.NewCollectorResolver(collectorProvider)
	layerResolver := layer.NewLayerResolver(collectorResolver, builderConfiguration.Layers)
	collectors := map[config.CollectorType]contractLayer.CollectorInterface{
		//AttributeCollector
		config.TypeBool:           collector.NewBoolCollector(collectorResolver),
		config.TypeClass:          collector.NewClassCollector(),
		config.TypeClasslike:      collector.NewClassLikeCollector(),
		config.TypeClassNameRegex: collector.NewClassNameRegexCollector(),
		//CollectorType.TypeTagValueRegex: TagValueRegexCollector.NewTagValueRegexCollector(),
		config.TypeDirectory: collector.NewDirectoryCollector(),
		//CollectorType.TypeExtends: ExtendsCollector.NewExtendsCollector(collectorResolver),
		config.TypeFunctionName: collector.NewFunctionNameCollector(),
		config.TypeGlob:         collector.NewGlobCollector(projectDirectory),
		//ImplementsCollector
		config.TypeInheritance: inheritanceLevelCollector,
		config.TypeInterface:   collector.NewInterfaceCollector(),
		config.TypeInherits:    inheritsCollector,
		config.TypeLayer:       collector.NewLayerCollector(layerResolver),
		config.TypeMethod:      collector.NewMethodCollector(nikicPhpParser),
		config.TypeSuperGlobal: collector.NewSuperglobalCollector(),
		config.TypeTrait:       collector.NewTraitCollector(),
		config.TypeUses:        usesCollector,
		//CollectorType.TypePhpInternal: PhpInternalCollector
		config.TypeComposer: collector.NewComposerCollector(),
	}
	collectorProvider.Set(collectors)

	/*
	 * SetAnalyser
	 */
	dependencyLayersAnalyser := analyser.NewDependencyLayersAnalyser(astMapExtractor, dependencyResolver, tokenResolver, layerResolver, eventDispatcher)
	tokenInLayerAnalyser := analyser.NewTokenInLayerAnalyser(astMapExtractor, tokenResolver, layerResolver, builderConfiguration.Analyser)
	layerForTokenAnalyser := analyser.NewLayerForTokenAnalyser(astMapExtractor, tokenResolver, layerResolver)
	unassignedTokenAnalyser := analyser.NewUnassignedTokenAnalyser(astMapExtractor, tokenResolver, layerResolver, builderConfiguration.Analyser)
	layerDependenciesAnalyser := analyser.NewLayerDependenciesAnalyser(astMapExtractor, tokenResolver, dependencyResolver, layerResolver)
	rulesetUsageAnalyser := analyser.NewRulesetUsageAnalyser(layerProvider, layerResolver, astMapExtractor, dependencyResolver, tokenResolver, builderConfiguration.Layers)

	/*
	 * OutputFormatter
	 */
	outputFormatter := map[output_formatter.OutputFormatterType]output_formatter.OutputFormatterInterface{
		output_formatter.Table:         output_formatter2.NewTableOutputFormatter(),
		output_formatter.GithubActions: output_formatter2.NewGithubActionsOutputFormatter(),
		// TODO:
		// $services->set(ConsoleOutputFormatter::class)->tag('output_formatter');
		// $services->set(JUnitOutputFormatter::class)->tag('output_formatter');
		// $services->set(XMLOutputFormatter::class)->tag('output_formatter');
		// $services->set(BaselineOutputFormatter::class)->tag('output_formatter');
		// $services->set(JsonOutputFormatter::class)->tag('output_formatter');
		// $services->set(GraphVizOutputDisplayFormatter::class)->tag('output_formatter');
		// $services->set(GraphVizOutputImageFormatter::class)->tag('output_formatter');
		// $services->set(GraphVizOutputDotFormatter::class)->tag('output_formatter');
		// $services->set(GraphVizOutputHtmlFormatter::class)->tag('output_formatter');
		// $services->set(CodeclimateOutputFormatter::class)->tag('output_formatter');
		// $services->set(MermaidJSOutputFormatter::class)->tag('output_formatter');
	}
	formatterProvider := output_formatter2.NewFormatterProvider(outputFormatter)
	formatterConfiguration := configuration.NewFormatterConfiguration(builderConfiguration.Formatters)

	/*
	 * Console
	 */
	analyseRunner := command.NewAnalyseRunner(dependencyLayersAnalyser, formatterProvider)
	analyseCommand := command.NewAnalyseCommand(analyseRunner, eventDispatcher, formatterProvider, *verboseBoolFlag, *debugBoolFlag, consoleSubscriber, progressSubscriber)

	// TODO: other commands
	// $services->set(InitCommand::class)->autowire()->tag('console.command');
	// $services->set(ChangedFilesRunner::class)->autowire();
	// $services->set(ChangedFilesCommand::class)->autowire()->tag('console.command');
	// $services->set(DebugLayerRunner::class)->autowire()->args(['$layers' => param('layers')]);
	// $services->set(DebugLayerCommand::class)->autowire()->tag('console.command');
	// $services->set(DebugTokenRunner::class)->autowire();
	// $services->set(DebugTokenCommand::class)->autowire()->tag('console.command');
	// $services->set(DebugUnassignedRunner::class)->autowire();
	// $services->set(DebugUnassignedCommand::class)->autowire()->tag('console.command');
	// $services->set(DebugDependenciesRunner::class)->autowire();
	// $services->set(DebugDependenciesCommand::class)->autowire()->tag('console.command');
	// $services->set(DebugUnusedRunner::class)->autowire();
	// $services->set(DebugUnusedCommand::class)->autowire()->tag('console.command');

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

	return nil
}

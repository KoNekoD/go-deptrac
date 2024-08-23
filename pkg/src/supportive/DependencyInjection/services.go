package DependencyInjection

import (
	"flag"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Ast/PostCreateAstMapEvent"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Ast/PreCreateAstMapEvent"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Config/CollectorType"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Config/EmitterType"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Layer/CollectorInterface"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Layer/LayerProvider"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/OutputFormatter/OutputFormatterInterface"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/OutputFormatter/OutputFormatterInterface/OutputFormatterType"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/analyser/event_helper"
	post_process_event2 "github.com/KoNekoD/go-deptrac/pkg/src/contract/analyser/post_process_event"
	process_event2 "github.com/KoNekoD/go-deptrac/pkg/src/contract/analyser/process_event"
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
	"github.com/KoNekoD/go-deptrac/pkg/src/supportive/Console/Command/AnalyseCommand"
	"github.com/KoNekoD/go-deptrac/pkg/src/supportive/Console/Command/AnalyseRunner"
	"github.com/KoNekoD/go-deptrac/pkg/src/supportive/Console/Subscriber/ConsoleSubscriber"
	"github.com/KoNekoD/go-deptrac/pkg/src/supportive/Console/Subscriber/ProgressSubscriber"
	"github.com/KoNekoD/go-deptrac/pkg/src/supportive/Console/Symfony/Style"
	"github.com/KoNekoD/go-deptrac/pkg/src/supportive/Console/Symfony/SymfonyOutput"
	"github.com/KoNekoD/go-deptrac/pkg/src/supportive/DependencyInjection/ContainerBuilder"
	"github.com/KoNekoD/go-deptrac/pkg/src/supportive/DependencyInjection/EventDispatcher"
	EventSubscriberInterfaceMap2 "github.com/KoNekoD/go-deptrac/pkg/src/supportive/DependencyInjection/EventSubscriberDefaultPriority"
	EventSubscriberInterface2 "github.com/KoNekoD/go-deptrac/pkg/src/supportive/DependencyInjection/EventSubscriberInterface"
	"github.com/KoNekoD/go-deptrac/pkg/src/supportive/DependencyInjection/EventSubscriberInterfaceMap"
	"github.com/KoNekoD/go-deptrac/pkg/src/supportive/DependencyInjection/EventSubscriberInterfaceMap/EventSubscriberInterfaceMapReg"
	"github.com/KoNekoD/go-deptrac/pkg/src/supportive/File/Dumper"
	"github.com/KoNekoD/go-deptrac/pkg/src/supportive/File/YmlFileLoader"
	"github.com/KoNekoD/go-deptrac/pkg/src/supportive/OutputFormatter/Configuration/FormatterConfiguration"
	"github.com/KoNekoD/go-deptrac/pkg/src/supportive/OutputFormatter/FormatterProvider"
	"github.com/KoNekoD/go-deptrac/pkg/src/supportive/OutputFormatter/GithubActionsOutputFormatter"
	"github.com/KoNekoD/go-deptrac/pkg/src/supportive/OutputFormatter/TableOutputFormatter"
	"github.com/KoNekoD/go-deptrac/pkg/src/supportive/TimeStopwatch"
	"github.com/elliotchance/orderedmap/v2"
)

func Services(builder *ContainerBuilder.ContainerBuilder) error {
	cacheableFileSubscriber := builder.CacheableFileSubscriber
	configuration := builder.Configuration
	projectDirectory := builder.ProjectDirectory
	verboseBoolFlag := flag.Bool("verbose", false, "Verbose mode")
	debugBoolFlag := flag.Bool("debug", false, "Debug mode")
	style := Style.NewStyle(
		verboseBoolFlag != nil && *verboseBoolFlag == true,
		debugBoolFlag != nil && *debugBoolFlag == true,
	)
	symfonyOutput := SymfonyOutput.NewSymfonyOutput(style)

	timeStopwatch := TimeStopwatch.NewStopwatch()

	/*
	 * Utilities
	 */
	eventDispatcher := EventDispatcher.NewEventDispatcher(debugBoolFlag != nil && *debugBoolFlag == true)

	fileInputCollector, err := input_collector.NewFileInputCollector(
		configuration.Paths,
		configuration.ExcludeFiles,
		projectDirectory,
	)
	if err != nil {
		return err
	}

	ymlFileLoader := YmlFileLoader.NewYmlFileLoader()
	dumper := Dumper.NewDumper("/deptrac_template.yaml")

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
	dependencyEmitters := map[EmitterType.EmitterType]emitter.DependencyEmitterInterface{
		EmitterType.ClassToken:               emitter.NewClassDependencyEmitter(),
		EmitterType.ClassSuperGlobalToken:    emitter.NewClassSuperglobalDependencyEmitter(),
		EmitterType.FileToken:                emitter.NewFileDependencyEmitter(),
		EmitterType.FunctionToken:            emitter.NewFunctionDependencyEmitter(),
		EmitterType.FunctionCall:             emitter.NewFunctionCallDependencyEmitter(),
		EmitterType.FunctionSuperGlobalToken: emitter.NewFunctionSuperglobalDependencyEmitter(),
		EmitterType.UseToken:                 emitter.NewUsesDependencyEmitter(),
	}
	inheritanceFlattener := dependency.NewInheritanceFlattener()
	dependencyResolver := dependency_resolver.NewDependencyResolver(configuration.Analyser, dependencyEmitters, inheritanceFlattener, eventDispatcher)
	tokenResolver := dependency.NewTokenResolver()

	astMapExtractor := ast.NewAstMapExtractor(fileInputCollector, astLoader)

	layerProvider := LayerProvider.NewLayerProvider(configuration.Rulesets)
	eventHelper := event_helper.NewEventHelper(configuration.SkipViolations, layerProvider)

	/*
	 * Events (before first possible event)
	 */
	/*
	 * Events
	 */
	EventSubscriberInterfaceMap.Map = orderedmap.NewOrderedMap[string, *orderedmap.OrderedMap[int, []EventSubscriberInterface2.EventSubscriberInterface]]()

	// Events
	uncoveredDependentHandler := process_event.NewUncoveredDependentHandler(configuration.IgnoreUncoveredInternalStructs)
	matchingLayersHandler := process_event.NewMatchingLayersHandler()
	allowDependencyHandler := process_event.NewAllowDependencyHandler()
	consoleSubscriber := ConsoleSubscriber.NewConsoleSubscriber(symfonyOutput, timeStopwatch)
	progressSubscriber := ProgressSubscriber.NewProgressSubscriber(symfonyOutput)
	dependsOnDisallowedLayer := process_event.NewDependsOnDisallowedLayer(eventHelper)
	dependsOnPrivateLayer := process_event.NewDependsOnPrivateLayer(eventHelper)
	dependsOnInternalToken := process_event.NewDependsOnInternalToken(eventHelper, configuration.Analyser)
	unmatchedSkippedViolations := post_process_event.NewUnmatchedSkippedViolations(eventHelper)

	processEvent := &process_event2.ProcessEvent{}
	postProcessEvent := &post_process_event2.PostProcessEvent{}
	preCreateAstMapEvent := &PreCreateAstMapEvent.PreCreateAstMapEvent{}
	postCreateAstMapEvent := &PostCreateAstMapEvent.PostCreateAstMapEvent{}
	// Events Handlers
	// TODO: Тут надо реализовать глобальный хук на параметры deptrac чтобы сделать что-то вида "param('skip_violations')"
	EventSubscriberInterfaceMapReg.Reg(processEvent, allowDependencyHandler, -100)
	EventSubscriberInterfaceMapReg.Reg(processEvent, dependsOnPrivateLayer, -3)
	EventSubscriberInterfaceMapReg.Reg(processEvent, dependsOnInternalToken, -2)
	EventSubscriberInterfaceMapReg.Reg(processEvent, dependsOnDisallowedLayer, -1)
	EventSubscriberInterfaceMapReg.Reg(processEvent, matchingLayersHandler, 1)
	EventSubscriberInterfaceMapReg.Reg(processEvent, uncoveredDependentHandler, 2)
	EventSubscriberInterfaceMapReg.Reg(postProcessEvent, unmatchedSkippedViolations, EventSubscriberInterfaceMap2.DefaultPriority)
	if cacheableFileSubscriber != nil {
		EventSubscriberInterfaceMapReg.Reg(preCreateAstMapEvent, cacheableFileSubscriber, EventSubscriberInterfaceMap2.DefaultPriority)
		EventSubscriberInterfaceMapReg.Reg(postCreateAstMapEvent, cacheableFileSubscriber, EventSubscriberInterfaceMap2.DefaultPriority)
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
	layerResolver := layer.NewLayerResolver(collectorResolver, configuration.Layers)
	collectors := map[CollectorType.CollectorType]CollectorInterface.CollectorInterface{
		//AttributeCollector
		CollectorType.TypeBool:           collector.NewBoolCollector(collectorResolver),
		CollectorType.TypeClass:          collector.NewClassCollector(),
		CollectorType.TypeClasslike:      collector.NewClassLikeCollector(),
		CollectorType.TypeClassNameRegex: collector.NewClassNameRegexCollector(),
		//CollectorType.TypeTagValueRegex: TagValueRegexCollector.NewTagValueRegexCollector(),
		CollectorType.TypeDirectory: collector.NewDirectoryCollector(),
		//CollectorType.TypeExtends: ExtendsCollector.NewExtendsCollector(collectorResolver),
		CollectorType.TypeFunctionName: collector.NewFunctionNameCollector(),
		CollectorType.TypeGlob:         collector.NewGlobCollector(projectDirectory),
		//ImplementsCollector
		CollectorType.TypeInheritance: inheritanceLevelCollector,
		CollectorType.TypeInterface:   collector.NewInterfaceCollector(),
		CollectorType.TypeInherits:    inheritsCollector,
		CollectorType.TypeLayer:       collector.NewLayerCollector(layerResolver),
		CollectorType.TypeMethod:      collector.NewMethodCollector(nikicPhpParser),
		CollectorType.TypeSuperGlobal: collector.NewSuperglobalCollector(),
		CollectorType.TypeTrait:       collector.NewTraitCollector(),
		CollectorType.TypeUses:        usesCollector,
		//CollectorType.TypePhpInternal: PhpInternalCollector
		CollectorType.TypeComposer: collector.NewComposerCollector(),
	}
	collectorProvider.Set(collectors)

	/*
	 * SetAnalyser
	 */
	dependencyLayersAnalyser := analyser.NewDependencyLayersAnalyser(astMapExtractor, dependencyResolver, tokenResolver, layerResolver, eventDispatcher)
	tokenInLayerAnalyser := analyser.NewTokenInLayerAnalyser(astMapExtractor, tokenResolver, layerResolver, configuration.Analyser)
	layerForTokenAnalyser := analyser.NewLayerForTokenAnalyser(astMapExtractor, tokenResolver, layerResolver)
	unassignedTokenAnalyser := analyser.NewUnassignedTokenAnalyser(astMapExtractor, tokenResolver, layerResolver, configuration.Analyser)
	layerDependenciesAnalyser := analyser.NewLayerDependenciesAnalyser(astMapExtractor, tokenResolver, dependencyResolver, layerResolver)
	rulesetUsageAnalyser := analyser.NewRulesetUsageAnalyser(layerProvider, layerResolver, astMapExtractor, dependencyResolver, tokenResolver, configuration.Layers)

	/*
	 * OutputFormatter
	 */
	outputFormatter := map[OutputFormatterType.OutputFormatterType]OutputFormatterInterface.OutputFormatterInterface{
		OutputFormatterType.Table:         TableOutputFormatter.NewTableOutputFormatter(),
		OutputFormatterType.GithubActions: GithubActionsOutputFormatter.NewGithubActionsOutputFormatter(),
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
	formatterProvider := FormatterProvider.NewFormatterProvider(outputFormatter)
	formatterConfiguration := FormatterConfiguration.NewFormatterConfiguration(configuration.Formatters)

	/*
	 * Console
	 */
	analyseRunner := AnalyseRunner.NewAnalyseRunner(dependencyLayersAnalyser, formatterProvider)
	analyseCommand := AnalyseCommand.NewAnalyseCommand(analyseRunner, eventDispatcher, formatterProvider, *verboseBoolFlag, *debugBoolFlag, consoleSubscriber, progressSubscriber)

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

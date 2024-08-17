package DependencyInjection

import (
	"flag"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Analyser/EventHelper"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Analyser/PostProcessEvent"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Analyser/ProcessEvent"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Ast/PostCreateAstMapEvent"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Ast/PreCreateAstMapEvent"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Config/CollectorType"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Config/EmitterType"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Layer/CollectorInterface"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Layer/LayerProvider"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/OutputFormatter/OutputFormatterInterface"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/OutputFormatter/OutputFormatterInterface/OutputFormatterType"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/Analyser/DependencyLayersAnalyser"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/Analyser/EventHandler/PostProcessEvent/UnmatchedSkippedViolations"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/Analyser/EventHandler/ProcessEvent/AllowDependencyHandler"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/Analyser/EventHandler/ProcessEvent/DependsOnDisallowedLayer"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/Analyser/EventHandler/ProcessEvent/DependsOnInternalToken"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/Analyser/EventHandler/ProcessEvent/DependsOnPrivateLayer"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/Analyser/EventHandler/ProcessEvent/MatchingLayersHandler"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/Analyser/EventHandler/ProcessEvent/UncoveredDependentHandler"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/Analyser/LayerDependenciesAnalyser"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/Analyser/LayerForTokenAnalyser"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/Analyser/RulesetUsageAnalyser"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/Analyser/TokenInLayerAnalyser"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/Analyser/UnassignedTokenAnalyser"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/Ast/AstLoader"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/Ast/AstMapExtractor"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/Ast/Parser/Cache/AstFileReferenceInMemoryCache"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/Ast/Parser/Extractors/ReferenceExtractorInterface"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/Ast/Parser/NikicPhpParser/NikicPhpParser"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/Ast/Parser/TypeResolver"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/Dependency/DependencyResolver"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/Dependency/Emitter/ClassDependencyEmitter"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/Dependency/Emitter/ClassSuperglobalDependencyEmitter"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/Dependency/Emitter/DependencyEmitterInterface"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/Dependency/Emitter/FileDependencyEmitter"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/Dependency/Emitter/FunctionCallDependencyEmitter"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/Dependency/Emitter/FunctionDependencyEmitter"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/Dependency/Emitter/FunctionSuperglobalDependencyEmitter"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/Dependency/Emitter/UsesDependencyEmitter"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/Dependency/InheritanceFlattener"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/Dependency/TokenResolver"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/InputCollector/FileInputCollector"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/Layer/Collector/BoolCollector"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/Layer/Collector/ClassCollector"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/Layer/Collector/ClassLikeCollector"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/Layer/Collector/ClassNameRegexCollector"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/Layer/Collector/CollectorProvider"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/Layer/Collector/CollectorResolver"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/Layer/Collector/ComposerCollector"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/Layer/Collector/DirectoryCollector"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/Layer/Collector/FunctionNameCollector"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/Layer/Collector/GlobCollector"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/Layer/Collector/InheritanceLevelCollector"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/Layer/Collector/InheritsCollector"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/Layer/Collector/InterfaceCollector"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/Layer/Collector/LayerCollector"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/Layer/Collector/MethodCollector"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/Layer/Collector/SuperglobalCollector"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/Layer/Collector/TraitCollector"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/Layer/Collector/UsesCollector"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/Layer/LayerResolver"
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

	fileInputCollector, err := FileInputCollector.NewFileInputCollector(
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
	astFileReferenceInMemoryCache := AstFileReferenceInMemoryCache.NewAstFileReferenceInMemoryCache()
	if builder.AstFileReferenceCacheInterface == nil {
		builder.AstFileReferenceCacheInterface = astFileReferenceInMemoryCache
	}
	typeResolver := TypeResolver.NewTypeResolver()
	referenceExtractors := []ReferenceExtractorInterface.ReferenceExtractorInterface{
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
	nikicPhpParser := NikicPhpParser.NewNikicPhpParser(builder.AstFileReferenceCacheInterface, typeResolver, referenceExtractors)
	parserInterface := nikicPhpParser
	astLoader := AstLoader.NewAstLoader(parserInterface, eventDispatcher)

	/*
	 * Dependency
	 */
	dependencyEmitters := map[EmitterType.EmitterType]DependencyEmitterInterface.DependencyEmitterInterface{
		EmitterType.ClassToken:               ClassDependencyEmitter.NewClassDependencyEmitter(),
		EmitterType.ClassSuperGlobalToken:    ClassSuperglobalDependencyEmitter.NewClassSuperglobalDependencyEmitter(),
		EmitterType.FileToken:                FileDependencyEmitter.NewFileDependencyEmitter(),
		EmitterType.FunctionToken:            FunctionDependencyEmitter.NewFunctionDependencyEmitter(),
		EmitterType.FunctionCall:             FunctionCallDependencyEmitter.NewFunctionCallDependencyEmitter(),
		EmitterType.FunctionSuperGlobalToken: FunctionSuperglobalDependencyEmitter.NewFunctionSuperglobalDependencyEmitter(),
		EmitterType.UseToken:                 UsesDependencyEmitter.NewUsesDependencyEmitter(),
	}
	inheritanceFlattener := InheritanceFlattener.NewInheritanceFlattener()
	dependencyResolver := DependencyResolver.NewDependencyResolver(configuration.Analyser, dependencyEmitters, inheritanceFlattener, eventDispatcher)
	tokenResolver := TokenResolver.NewTokenResolver()

	astMapExtractor := AstMapExtractor.NewAstMapExtractor(fileInputCollector, astLoader)

	layerProvider := LayerProvider.NewLayerProvider(configuration.Rulesets)
	eventHelper := EventHelper.NewEventHelper(configuration.SkipViolations, layerProvider)

	/*
	 * Events (before first possible event)
	 */
	/*
	 * Events
	 */
	EventSubscriberInterfaceMap.Map = orderedmap.NewOrderedMap[string, *orderedmap.OrderedMap[int, []EventSubscriberInterface2.EventSubscriberInterface]]()

	// Events
	uncoveredDependentHandler := UncoveredDependentHandler.NewUncoveredDependentHandler(configuration.IgnoreUncoveredInternalStructs)
	matchingLayersHandler := MatchingLayersHandler.NewMatchingLayersHandler()
	allowDependencyHandler := AllowDependencyHandler.NewAllowDependencyHandler()
	consoleSubscriber := ConsoleSubscriber.NewConsoleSubscriber(symfonyOutput, timeStopwatch)
	progressSubscriber := ProgressSubscriber.NewProgressSubscriber(symfonyOutput)
	dependsOnDisallowedLayer := DependsOnDisallowedLayer.NewDependsOnDisallowedLayer(eventHelper)
	dependsOnPrivateLayer := DependsOnPrivateLayer.NewDependsOnPrivateLayer(eventHelper)
	dependsOnInternalToken := DependsOnInternalToken.NewDependsOnInternalToken(eventHelper, configuration.Analyser)
	unmatchedSkippedViolations := UnmatchedSkippedViolations.NewUnmatchedSkippedViolations(eventHelper)

	processEvent := &ProcessEvent.ProcessEvent{}
	postProcessEvent := &PostProcessEvent.PostProcessEvent{}
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
	inheritanceLevelCollector, err := InheritanceLevelCollector.NewInheritanceLevelCollector(astMapExtractor)
	if err != nil {
		return err
	}
	inheritsCollector, err := InheritsCollector.NewInheritsCollector(astMapExtractor)
	if err != nil {
		return err
	}
	usesCollector, err := UsesCollector.NewUsesCollector(astMapExtractor)
	if err != nil {
		return err
	}
	collectorProvider := CollectorProvider.NewCollectorProvider()
	collectorResolver := CollectorResolver.NewCollectorResolver(collectorProvider)
	layerResolver := LayerResolver.NewLayerResolver(collectorResolver, configuration.Layers)
	collectors := map[CollectorType.CollectorType]CollectorInterface.CollectorInterface{
		//AttributeCollector
		CollectorType.TypeBool:           BoolCollector.NewBoolCollector(collectorResolver),
		CollectorType.TypeClass:          ClassCollector.NewClassCollector(),
		CollectorType.TypeClasslike:      ClassLikeCollector.NewClassLikeCollector(),
		CollectorType.TypeClassNameRegex: ClassNameRegexCollector.NewClassNameRegexCollector(),
		//CollectorType.TypeTagValueRegex: TagValueRegexCollector.NewTagValueRegexCollector(),
		CollectorType.TypeDirectory: DirectoryCollector.NewDirectoryCollector(),
		//CollectorType.TypeExtends: ExtendsCollector.NewExtendsCollector(collectorResolver),
		CollectorType.TypeFunctionName: FunctionNameCollector.NewFunctionNameCollector(),
		CollectorType.TypeGlob:         GlobCollector.NewGlobCollector(projectDirectory),
		//ImplementsCollector
		CollectorType.TypeInheritance: inheritanceLevelCollector,
		CollectorType.TypeInterface:   InterfaceCollector.NewInterfaceCollector(),
		CollectorType.TypeInherits:    inheritsCollector,
		CollectorType.TypeLayer:       LayerCollector.NewLayerCollector(layerResolver),
		CollectorType.TypeMethod:      MethodCollector.NewMethodCollector(nikicPhpParser),
		CollectorType.TypeSuperGlobal: SuperglobalCollector.NewSuperglobalCollector(),
		CollectorType.TypeTrait:       TraitCollector.NewTraitCollector(),
		CollectorType.TypeUses:        usesCollector,
		//CollectorType.TypePhpInternal: PhpInternalCollector
		CollectorType.TypeComposer: ComposerCollector.NewComposerCollector(),
	}
	collectorProvider.Set(collectors)

	/*
	 * SetAnalyser
	 */
	dependencyLayersAnalyser := DependencyLayersAnalyser.NewDependencyLayersAnalyser(astMapExtractor, dependencyResolver, tokenResolver, layerResolver, eventDispatcher)
	tokenInLayerAnalyser := TokenInLayerAnalyser.NewTokenInLayerAnalyser(astMapExtractor, tokenResolver, layerResolver, configuration.Analyser)
	layerForTokenAnalyser := LayerForTokenAnalyser.NewLayerForTokenAnalyser(astMapExtractor, tokenResolver, layerResolver)
	unassignedTokenAnalyser := UnassignedTokenAnalyser.NewUnassignedTokenAnalyser(astMapExtractor, tokenResolver, layerResolver, configuration.Analyser)
	layerDependenciesAnalyser := LayerDependenciesAnalyser.NewLayerDependenciesAnalyser(astMapExtractor, tokenResolver, dependencyResolver, layerResolver)
	rulesetUsageAnalyser := RulesetUsageAnalyser.NewRulesetUsageAnalyser(layerProvider, layerResolver, astMapExtractor, dependencyResolver, tokenResolver, configuration.Layers)

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

package app

import (
	"flag"
	"fmt"
	"github.com/KoNekoD/go-deptrac/pkg/ast_map"
	"github.com/KoNekoD/go-deptrac/pkg/collectors"
	"github.com/KoNekoD/go-deptrac/pkg/commands"
	"github.com/KoNekoD/go-deptrac/pkg/dependencies"
	"github.com/KoNekoD/go-deptrac/pkg/domain/stopwatch"
	"github.com/KoNekoD/go-deptrac/pkg/domain/utils"
	"github.com/KoNekoD/go-deptrac/pkg/emitters"
	"github.com/KoNekoD/go-deptrac/pkg/events"
	"github.com/KoNekoD/go-deptrac/pkg/flatteners"
	"github.com/KoNekoD/go-deptrac/pkg/formatters"
	"github.com/KoNekoD/go-deptrac/pkg/hooks"
	"github.com/KoNekoD/go-deptrac/pkg/layers"
	"github.com/KoNekoD/go-deptrac/pkg/nodes"
	"github.com/KoNekoD/go-deptrac/pkg/parsers"
	"github.com/KoNekoD/go-deptrac/pkg/references"
	"github.com/KoNekoD/go-deptrac/pkg/results"
	"github.com/KoNekoD/go-deptrac/pkg/rules"
	"github.com/KoNekoD/go-deptrac/pkg/runners"
	"github.com/KoNekoD/go-deptrac/pkg/subscribers"
	"github.com/KoNekoD/go-deptrac/pkg/tokens"
	"github.com/KoNekoD/go-deptrac/pkg/types"
	"github.com/KoNekoD/go-deptrac/pkg/violations"
	"github.com/elliotchance/orderedmap/v2"
	"os"
	"strings"
)

func getDefaultFormatter() formatters.OutputFormatterType {
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
	symfonyOutput := results.NewSymfonyOutput(style)

	timeStopwatch := stopwatch.NewStopwatch()

	nodeNamer := nodes.NewNodeNamer(projectDirectory)

	/*
	 * Utilities
	 */
	eventDispatcher := events.NewEventDispatcher(debugBoolFlag != nil && *debugBoolFlag == true)

	fileInputCollector, err := collectors.NewFileInputCollector(
		builderConfiguration.Paths,
		builderConfiguration.ExcludeFiles,
		projectDirectory,
	)
	if err != nil {
		return err
	}

	ymlFileLoader := hooks.NewYmlFileLoader()
	dumper := utils.NewDumper("/deptrac_template.yaml")

	/*
	 * AST
	 */
	astFileReferenceInMemoryCache := ast_map.NewAstFileReferenceInMemoryCache()
	if builder.AstFileReferenceCacheInterface == nil {
		builder.AstFileReferenceCacheInterface = astFileReferenceInMemoryCache
	}
	typeResolver := types.NewTypeResolver(nodeNamer)
	referenceExtractors := []references.ReferenceExtractorInterface{
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
	astLoader := ast_map.NewAstLoader(parserInterface, eventDispatcher)

	/*
	 * Dependency
	 */
	dependencyEmitters := map[emitters.EmitterType]emitters.DependencyEmitterInterface{
		emitters.EmitterTypeClassToken:               emitters.NewClassDependencyEmitter(),
		emitters.EmitterTypeClassSuperGlobalToken:    emitters.NewClassSuperglobalDependencyEmitter(),
		emitters.EmitterTypeFileToken:                emitters.NewFileDependencyEmitter(),
		emitters.EmitterTypeFunctionToken:            emitters.NewFunctionDependencyEmitter(),
		emitters.EmitterTypeFunctionCall:             emitters.NewFunctionCallDependencyEmitter(),
		emitters.EmitterTypeFunctionSuperGlobalToken: emitters.NewFunctionSuperglobalDependencyEmitter(),
		emitters.EmitterTypeUseToken:                 emitters.NewUsesDependencyEmitter(),
	}
	inheritanceFlattener := flatteners.NewInheritanceFlattener()
	dependencyResolver := dependencies.NewDependencyResolver(builderConfiguration.Analyser, dependencyEmitters, inheritanceFlattener, eventDispatcher)
	tokenResolver := tokens.NewTokenResolver()

	astMapExtractor := ast_map.NewAstMapExtractor(fileInputCollector, astLoader)

	layerProvider := layers.NewLayerProvider(builderConfiguration.Rulesets)
	eventHelper := events.NewEventHelper(builderConfiguration.SkipViolations, layerProvider)

	/*
	 * Events (before first possible event)
	 */
	/*
	 * Events
	 */
	subscribers.Map = orderedmap.NewOrderedMap[string, *orderedmap.OrderedMap[int, []subscribers.EventSubscriberInterface]]()

	// Events
	uncoveredDependentHandler := dependencies.NewUncoveredDependentHandler(builderConfiguration.IgnoreUncoveredInternalStructs)
	matchingLayersHandler := layers.NewMatchingLayersHandler()
	allowDependencyHandler := dependencies.NewAllowDependencyHandler()
	consoleSubscriber := subscribers.NewConsoleSubscriber(symfonyOutput, timeStopwatch)
	progressSubscriber := subscribers.NewProgressSubscriber(symfonyOutput)
	dependsOnDisallowedLayer := subscribers.NewDependsOnDisallowedLayer(eventHelper)
	dependsOnPrivateLayer := subscribers.NewDependsOnPrivateLayer(eventHelper)
	dependsOnInternalToken := subscribers.NewDependsOnInternalToken(eventHelper, builderConfiguration.Analyser)
	unmatchedSkippedViolations := violations.NewUnmatchedSkippedViolations(eventHelper)

	processEvent := &events.ProcessEvent{}
	postProcessEvent := &events.PostProcessEvent{}
	preCreateAstMapEvent := &ast_map.PreCreateAstMapEvent{}
	postCreateAstMapEvent := &ast_map.PostCreateAstMapEvent{}
	// Events Handlers
	// TODO: Тут надо реализовать глобальный хук на параметры deptrac чтобы сделать что-то вида "param('skip_violations')"
	subscribers.Reg(processEvent, allowDependencyHandler, -100)
	subscribers.Reg(processEvent, dependsOnPrivateLayer, -3)
	subscribers.Reg(processEvent, dependsOnInternalToken, -2)
	subscribers.Reg(processEvent, dependsOnDisallowedLayer, -1)
	subscribers.Reg(processEvent, matchingLayersHandler, 1)
	subscribers.Reg(processEvent, uncoveredDependentHandler, 2)
	subscribers.Reg(postProcessEvent, unmatchedSkippedViolations, subscribers.DefaultPriority)
	if cacheableFileSubscriber != nil {
		subscribers.Reg(preCreateAstMapEvent, cacheableFileSubscriber, subscribers.DefaultPriority)
		subscribers.Reg(postCreateAstMapEvent, cacheableFileSubscriber, subscribers.DefaultPriority)
	}

	/*
	 * OutputFormatter
	 */
	outputFormatter := map[formatters.OutputFormatterType]formatters.OutputFormatterInterface{
		formatters.Table:         formatters.NewTableOutputFormatter(),
		formatters.GithubActions: formatters.NewGithubActionsOutputFormatter(),
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
		formatter              = flag.String("formatter", string(formatters.Table), formatterUsage)
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

	analyseOptions := rules.NewAnalyseOptions(
		nil != noProgress && *noProgress == true,
		*formatter,
		output,
		nil != reportSkipped && *reportSkipped == true,
		nil != reportUncovered && *reportUncovered == true,
		nil != failOnUncovered && *failOnUncovered == true,
	)
	subscribers.RegForAnalyseCommand(consoleSubscriber, progressSubscriber, !analyseOptions.NoProgress)
	//

	/*
	 * Layer
	 */
	inheritanceLevelCollector, err := collectors.NewInheritanceLevelCollector(astMapExtractor)
	if err != nil {
		return err
	}
	inheritsCollector, err := collectors.NewInheritsCollector(astMapExtractor)
	if err != nil {
		return err
	}
	usesCollector, err := collectors.NewUsesCollector(astMapExtractor)
	if err != nil {
		return err
	}
	collectorProvider := collectors.NewCollectorProvider()
	collectorResolver := collectors.NewCollectorResolver(collectorProvider)
	layerResolver := layers.NewLayerResolver(collectorResolver, builderConfiguration.Layers)
	collectors := map[collectors.CollectorType]collectors.CollectorInterface{
		//AttributeCollector
		collectors.CollectorTypeTypeBool:           collectors.NewBoolCollector(collectorResolver),
		collectors.CollectorTypeTypeClass:          collectors.NewClassCollector(),
		collectors.CollectorTypeTypeClasslike:      collectors.NewClassLikeCollector(),
		collectors.CollectorTypeTypeClassNameRegex: collectors.NewClassNameRegexCollector(),
		//CollectorType.TypeTagValueRegex: TagValueRegexCollector.NewTagValueRegexCollector(),
		collectors.CollectorTypeTypeDirectory: collectors.NewDirectoryCollector(),
		//CollectorType.TypeExtends: ExtendsCollector.NewExtendsCollector(collectorResolver),
		collectors.CollectorTypeTypeFunctionName: collectors.NewFunctionNameCollector(),
		collectors.CollectorTypeTypeGlob:         collectors.NewGlobCollector(projectDirectory),
		//ImplementsCollector
		collectors.CollectorTypeTypeInheritance: inheritanceLevelCollector,
		collectors.CollectorTypeTypeInterface:   collectors.NewInterfaceCollector(),
		collectors.CollectorTypeTypeInherits:    inheritsCollector,
		collectors.CollectorTypeTypeLayer:       collectors.NewLayerCollector(layerResolver),
		collectors.CollectorTypeTypeMethod:      collectors.NewMethodCollector(nikicPhpParser),
		collectors.CollectorTypeTypeSuperGlobal: collectors.NewSuperglobalCollector(),
		collectors.CollectorTypeTypeTrait:       collectors.NewTraitCollector(),
		collectors.CollectorTypeTypeUses:        usesCollector,
		//CollectorType.TypePhpInternal: PhpInternalCollector
		collectors.CollectorTypeTypeComposer: collectors.NewComposerCollector(),
	}
	collectorProvider.Set(collectors)

	/*
	 * SetAnalyser
	 */
	dependencyLayersAnalyser := dependencies.NewDependencyLayersAnalyser(astMapExtractor, dependencyResolver, tokenResolver, layerResolver, eventDispatcher)
	tokenInLayerAnalyser := tokens.NewTokenInLayerAnalyser(astMapExtractor, tokenResolver, layerResolver, builderConfiguration.Analyser)
	layerForTokenAnalyser := tokens.NewLayerForTokenAnalyser(astMapExtractor, tokenResolver, layerResolver)
	unassignedTokenAnalyser := tokens.NewUnassignedTokenAnalyser(astMapExtractor, tokenResolver, layerResolver, builderConfiguration.Analyser)
	layerDependenciesAnalyser := dependencies.NewLayerDependenciesAnalyser(astMapExtractor, tokenResolver, dependencyResolver, layerResolver)
	rulesetUsageAnalyser := rules.NewRulesetUsageAnalyser(layerProvider, layerResolver, astMapExtractor, dependencyResolver, tokenResolver, builderConfiguration.Layers)

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

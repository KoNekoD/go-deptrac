package di

import (
	"github.com/KoNekoD/go-deptrac/pkg/application/event_dispatchers"
	"github.com/KoNekoD/go-deptrac/pkg/application/event_handlers"
	applicationServices "github.com/KoNekoD/go-deptrac/pkg/application/services"
	"github.com/KoNekoD/go-deptrac/pkg/application/services/analysers"
	"github.com/KoNekoD/go-deptrac/pkg/application/services/ast_file_reference_cache"
	"github.com/KoNekoD/go-deptrac/pkg/application/services/ast_map"
	"github.com/KoNekoD/go-deptrac/pkg/application/services/collectors_resolvers"
	"github.com/KoNekoD/go-deptrac/pkg/application/services/input_collectors"
	"github.com/KoNekoD/go-deptrac/pkg/application/services/layers_resolvers"
	"github.com/KoNekoD/go-deptrac/pkg/application/services/parsers"
	"github.com/KoNekoD/go-deptrac/pkg/application/services/references_extractors"
	"github.com/KoNekoD/go-deptrac/pkg/application/services/types"
	"github.com/KoNekoD/go-deptrac/pkg/domain/dtos/commands_options"
	"github.com/KoNekoD/go-deptrac/pkg/domain/dtos/configs"
	domainServices "github.com/KoNekoD/go-deptrac/pkg/domain/services"
	"github.com/KoNekoD/go-deptrac/pkg/domain/stopwatch"
	"github.com/KoNekoD/go-deptrac/pkg/infrastructure/commands"
	"github.com/KoNekoD/go-deptrac/pkg/infrastructure/services"
	"github.com/KoNekoD/go-deptrac/pkg/infrastructure/services/formatters"
	"github.com/KoNekoD/go-deptrac/pkg/infrastructure/services/runners"
)

type ContainerBuilder struct {
	ProjectDirectory                       string
	CacheFile                              *string
	Configuration                          *configs.DeptracConfig
	EventDispatcher                        event_dispatchers.EventDispatcherInterface
	FileInputCollector                     input_collectors.InputCollector
	YmlFileLoader                          *services.YmlFileLoader
	Dumper                                 *domainServices.Dumper
	AstLoader                              *ast_map.AstLoader
	AstFileReferenceFileCache              *ast_file_reference_cache.AstFileReferenceFileCache
	AstFileReferenceDeferredCacheInterface ast_file_reference_cache.AstFileReferenceDeferredCacheInterface
	AstFileReferenceCacheInterface         ast_file_reference_cache.AstFileReferenceCacheInterface
	CacheableFileSubscriber                *event_handlers.CacheableFile
	AstFileReferenceInMemoryCache          *ast_file_reference_cache.AstFileReferenceInMemoryCache
	TypeResolver                           *types.TypeResolver
	ReferenceExtractors                    []references_extractors.ReferenceExtractorInterface
	ParserInterface                        parsers.ParserInterface
	LayerProvider                          *applicationServices.LayerProvider
	EventHelper                            *applicationServices.EventHelper
	AllowDependencyHandler                 *event_handlers.AllowDependency
	DependsOnPrivateLayer                  *event_handlers.DependsOnPrivateLayer
	DependsOnInternalToken                 *event_handlers.DependsOnInternalToken
	DependsOnDisallowedLayer               *event_handlers.DependsOnDisallowedLayer
	MatchingLayersHandler                  *event_handlers.MatchingLayers
	UncoveredDependentHandler              *event_handlers.UncoveredDependent
	UnmatchedSkippedViolations             *event_handlers.UnmatchedSkippedViolations
	ConsoleSubscriber                      *event_handlers.Console
	ProgressSubscriber                     *event_handlers.Progress
	VerboseBoolFlag                        *bool
	DebugBoolFlag                          *bool
	Style                                  *formatters.Style
	SymfonyOutput                          *services.SymfonyOutput
	TimeStopwatch                          *stopwatch.Stopwatch
	AstMapExtractor                        *ast_map.AstMapExtractor
	InheritanceFlattener                   *applicationServices.InheritanceFlattener
	DependencyResolver                     *applicationServices.DependencyResolver
	TokenResolver                          *applicationServices.TokenResolver
	CollectorResolver                      *collectors_resolvers.CollectorResolver
	LayerResolver                          layers_resolvers.LayerResolverInterface
	NikicPhpParser                         *parsers.NikicPhpParser
	CollectorProvider                      *applicationServices.CollectorProvider
	DependencyLayersAnalyser               *analysers.DependencyLayersAnalyser
	TokenInLayerAnalyser                   *analysers.TokenInLayerAnalyser
	LayerForTokenAnalyser                  *analysers.LayerForTokenAnalyser
	UnassignedTokenAnalyser                *analysers.UnassignedTokenAnalyser
	LayerDependenciesAnalyser              *analysers.LayerDependenciesAnalyser
	RulesetUsageAnalyser                   *analysers.RulesetUsageAnalyser
	FormatterProvider                      *formatters.FormatterProvider
	FormatterConfiguration                 *formatters.FormatterConfiguration
	AnalyseRunner                          *runners.AnalyseRunner
	NodeNamer                              *domainServices.NodeNamer
	AnalyseOptions                         *commands_options.AnalyseOptions
	AnalyseCommand                         *commands.AnalyseCommand
	ChangedFilesCommand                    *commands.ChangedFilesCommand
	DebugDependenciesCommand               *commands.DebugDependenciesCommand
	DebugLayerCommand                      *commands.DebugLayerCommand
	DebugTokenCommand                      *commands.DebugTokenCommand
	DebugUnassignedCommand                 *commands.DebugUnassignedCommand
	DebugUnusedCommand                     *commands.DebugUnusedCommand
	InitCommand                            *commands.InitCommand
}

func NewContainerBuilder(workingDirectory string) *ContainerBuilder {
	return &ContainerBuilder{}
}

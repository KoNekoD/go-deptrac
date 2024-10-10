package app

import (
	event_handlers2 "github.com/KoNekoD/go-deptrac/pkg/application/event_handlers"
	services2 "github.com/KoNekoD/go-deptrac/pkg/application/services"
	analysers2 "github.com/KoNekoD/go-deptrac/pkg/application/services/analysers"
	"github.com/KoNekoD/go-deptrac/pkg/application/services/ast_file_reference_cache"
	ast_map2 "github.com/KoNekoD/go-deptrac/pkg/application/services/ast_map"
	"github.com/KoNekoD/go-deptrac/pkg/application/services/collectors_resolvers"
	"github.com/KoNekoD/go-deptrac/pkg/application/services/extractors"
	"github.com/KoNekoD/go-deptrac/pkg/application/services/input_collectors"
	"github.com/KoNekoD/go-deptrac/pkg/application/services/layers_resolvers"
	parsers2 "github.com/KoNekoD/go-deptrac/pkg/application/services/parsers"
	"github.com/KoNekoD/go-deptrac/pkg/application/services/types"
	"github.com/KoNekoD/go-deptrac/pkg/domain/dtos/configs"
	"github.com/KoNekoD/go-deptrac/pkg/domain/dtos/options"
	"github.com/KoNekoD/go-deptrac/pkg/domain/services"
	"github.com/KoNekoD/go-deptrac/pkg/domain/stopwatch"
	"github.com/KoNekoD/go-deptrac/pkg/domain/utils"
	"github.com/KoNekoD/go-deptrac/pkg/infrastructure/commands"
	services3 "github.com/KoNekoD/go-deptrac/pkg/infrastructure/services"
	dispatchers2 "github.com/KoNekoD/go-deptrac/pkg/infrastructure/services/dispatchers"
	formatters2 "github.com/KoNekoD/go-deptrac/pkg/infrastructure/services/formatters"
)

type ContainerBuilder struct {
	ProjectDirectory                       string
	CacheFile                              *string
	Configuration                          *configs.DeptracConfig
	EventDispatcher                        dispatchers2.EventDispatcherInterface
	FileInputCollector                     input_collectors.InputCollector
	YmlFileLoader                          *services3.YmlFileLoader
	Dumper                                 *utils.Dumper
	AstLoader                              *ast_map2.AstLoader
	AstFileReferenceFileCache              *ast_file_reference_cache.AstFileReferenceFileCache
	AstFileReferenceDeferredCacheInterface ast_file_reference_cache.AstFileReferenceDeferredCacheInterface
	AstFileReferenceCacheInterface         ast_file_reference_cache.AstFileReferenceCacheInterface
	CacheableFileSubscriber                *event_handlers2.CacheableFile
	AstFileReferenceInMemoryCache          *ast_file_reference_cache.AstFileReferenceInMemoryCache
	TypeResolver                           *types.TypeResolver
	ReferenceExtractors                    []extractors.ReferenceExtractorInterface
	ParserInterface                        parsers2.ParserInterface
	LayerProvider                          *services2.LayerProvider
	EventHelper                            *dispatchers2.EventHelper
	AllowDependencyHandler                 *event_handlers2.AllowDependency
	DependsOnPrivateLayer                  *event_handlers2.DependsOnPrivateLayer
	DependsOnInternalToken                 *event_handlers2.DependsOnInternalToken
	DependsOnDisallowedLayer               *event_handlers2.DependsOnDisallowedLayer
	MatchingLayersHandler                  *event_handlers2.MatchingLayers
	UncoveredDependentHandler              *event_handlers2.UncoveredDependent
	UnmatchedSkippedViolations             *event_handlers2.UnmatchedSkippedViolations
	ConsoleSubscriber                      *event_handlers2.Console
	ProgressSubscriber                     *event_handlers2.Progress
	VerboseBoolFlag                        *bool
	DebugBoolFlag                          *bool
	Style                                  *formatters2.Style
	SymfonyOutput                          *services3.SymfonyOutput
	TimeStopwatch                          *stopwatch.Stopwatch
	AstMapExtractor                        *ast_map2.AstMapExtractor
	InheritanceFlattener                   *services2.InheritanceFlattener
	DependencyResolver                     *services2.DependencyResolver
	TokenResolver                          *services2.TokenResolver
	CollectorResolver                      *collectors_resolvers.CollectorResolver
	LayerResolver                          layers_resolvers.LayerResolverInterface
	NikicPhpParser                         *parsers2.NikicPhpParser
	CollectorProvider                      *services2.CollectorProvider
	DependencyLayersAnalyser               *analysers2.DependencyLayersAnalyser
	TokenInLayerAnalyser                   *analysers2.TokenInLayerAnalyser
	LayerForTokenAnalyser                  *analysers2.LayerForTokenAnalyser
	UnassignedTokenAnalyser                *analysers2.UnassignedTokenAnalyser
	LayerDependenciesAnalyser              *analysers2.LayerDependenciesAnalyser
	RulesetUsageAnalyser                   *analysers2.RulesetUsageAnalyser
	FormatterProvider                      *formatters2.FormatterProvider
	FormatterConfiguration                 *formatters2.FormatterConfiguration
	AnalyseRunner                          *AnalyseRunner
	AnalyseCommand                         *commands.AnalyseCommand
	NodeNamer                              *services.NodeNamer
	AnalyseOptions                         *options.AnalyseOptions
}

func NewContainerBuilder(workingDirectory string) *ContainerBuilder {
	return &ContainerBuilder{}
}

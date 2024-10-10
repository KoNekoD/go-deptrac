package app

import (
	"github.com/KoNekoD/go-deptrac/pkg"
	"github.com/KoNekoD/go-deptrac/pkg/analysers"
	event_handlers2 "github.com/KoNekoD/go-deptrac/pkg/application/event_handlers"
	services2 "github.com/KoNekoD/go-deptrac/pkg/application/services"
	"github.com/KoNekoD/go-deptrac/pkg/application/services/ast_file_reference_cache"
	ast_map2 "github.com/KoNekoD/go-deptrac/pkg/application/services/ast_map"
	"github.com/KoNekoD/go-deptrac/pkg/application/services/collectors_resolvers"
	"github.com/KoNekoD/go-deptrac/pkg/application/services/input_collectors"
	"github.com/KoNekoD/go-deptrac/pkg/application/services/layers_resolvers"
	parsers2 "github.com/KoNekoD/go-deptrac/pkg/application/services/parsers"
	"github.com/KoNekoD/go-deptrac/pkg/application/services/types"
	"github.com/KoNekoD/go-deptrac/pkg/commands"
	"github.com/KoNekoD/go-deptrac/pkg/dispatchers"
	"github.com/KoNekoD/go-deptrac/pkg/domain/dtos/configs"
	"github.com/KoNekoD/go-deptrac/pkg/domain/services"
	"github.com/KoNekoD/go-deptrac/pkg/domain/stopwatch"
	"github.com/KoNekoD/go-deptrac/pkg/domain/utils"
	"github.com/KoNekoD/go-deptrac/pkg/extractors"
	"github.com/KoNekoD/go-deptrac/pkg/formatters"
)

type ContainerBuilder struct {
	ProjectDirectory                       string
	CacheFile                              *string
	Configuration                          *configs.DeptracConfig
	EventDispatcher                        dispatchers.EventDispatcherInterface
	FileInputCollector                     input_collectors.InputCollector
	YmlFileLoader                          *pkg.YmlFileLoader
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
	EventHelper                            *dispatchers.EventHelper
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
	Style                                  *formatters.Style
	SymfonyOutput                          *pkg.SymfonyOutput
	TimeStopwatch                          *stopwatch.Stopwatch
	AstMapExtractor                        *ast_map2.AstMapExtractor
	InheritanceFlattener                   *services2.InheritanceFlattener
	DependencyResolver                     *pkg.DependencyResolver
	TokenResolver                          *services2.TokenResolver
	CollectorResolver                      *collectors_resolvers.CollectorResolver
	LayerResolver                          layers_resolvers.LayerResolverInterface
	NikicPhpParser                         *parsers2.NikicPhpParser
	CollectorProvider                      *services2.CollectorProvider
	DependencyLayersAnalyser               *analysers.DependencyLayersAnalyser
	TokenInLayerAnalyser                   *analysers.TokenInLayerAnalyser
	LayerForTokenAnalyser                  *analysers.LayerForTokenAnalyser
	UnassignedTokenAnalyser                *analysers.UnassignedTokenAnalyser
	LayerDependenciesAnalyser              *analysers.LayerDependenciesAnalyser
	RulesetUsageAnalyser                   *analysers.RulesetUsageAnalyser
	FormatterProvider                      *formatters.FormatterProvider
	FormatterConfiguration                 *formatters.FormatterConfiguration
	AnalyseRunner                          *AnalyseRunner
	AnalyseCommand                         *commands.AnalyseCommand
	NodeNamer                              *services.NodeNamer
	AnalyseOptions                         *pkg.AnalyseOptions
}

func NewContainerBuilder(workingDirectory string) *ContainerBuilder {
	return &ContainerBuilder{}
}

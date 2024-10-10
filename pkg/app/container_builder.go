package app

import (
	"github.com/KoNekoD/go-deptrac/pkg"
	"github.com/KoNekoD/go-deptrac/pkg/analysers"
	event_handlers2 "github.com/KoNekoD/go-deptrac/pkg/application/event_handlers"
	services2 "github.com/KoNekoD/go-deptrac/pkg/application/services"
	"github.com/KoNekoD/go-deptrac/pkg/application/services/collectors_resolvers"
	"github.com/KoNekoD/go-deptrac/pkg/application/services/input_collectors"
	parsers2 "github.com/KoNekoD/go-deptrac/pkg/application/services/parsers"
	"github.com/KoNekoD/go-deptrac/pkg/application/services/types"
	"github.com/KoNekoD/go-deptrac/pkg/ast_map"
	"github.com/KoNekoD/go-deptrac/pkg/commands"
	"github.com/KoNekoD/go-deptrac/pkg/dispatchers"
	"github.com/KoNekoD/go-deptrac/pkg/domain/dtos/configs"
	"github.com/KoNekoD/go-deptrac/pkg/domain/services"
	"github.com/KoNekoD/go-deptrac/pkg/domain/stopwatch"
	"github.com/KoNekoD/go-deptrac/pkg/domain/utils"
	"github.com/KoNekoD/go-deptrac/pkg/extractors"
	"github.com/KoNekoD/go-deptrac/pkg/flatteners"
	"github.com/KoNekoD/go-deptrac/pkg/formatters"
	"github.com/KoNekoD/go-deptrac/pkg/hooks"
	"github.com/KoNekoD/go-deptrac/pkg/layers"
)

type ContainerBuilder struct {
	ProjectDirectory                       string
	CacheFile                              *string
	Configuration                          *configs.DeptracConfig
	EventDispatcher                        dispatchers.EventDispatcherInterface
	FileInputCollector                     input_collectors.InputCollector
	YmlFileLoader                          *hooks.YmlFileLoader
	Dumper                                 *utils.Dumper
	AstLoader                              *ast_map.AstLoader
	AstFileReferenceFileCache              *ast_map.AstFileReferenceFileCache
	AstFileReferenceDeferredCacheInterface ast_map.AstFileReferenceDeferredCacheInterface
	AstFileReferenceCacheInterface         ast_map.AstFileReferenceCacheInterface
	CacheableFileSubscriber                *event_handlers2.CacheableFile
	AstFileReferenceInMemoryCache          *ast_map.AstFileReferenceInMemoryCache
	TypeResolver                           *types.TypeResolver
	ReferenceExtractors                    []extractors.ReferenceExtractorInterface
	ParserInterface                        parsers2.ParserInterface
	LayerProvider                          *layers.LayerProvider
	EventHelper                            *dispatchers.EventHelper
	AllowDependencyHandler                 *event_handlers2.AllowDependency
	DependsOnPrivateLayer                  *event_handlers2.DependsOnPrivateLayer
	DependsOnInternalToken                 *event_handlers2.DependsOnInternalToken
	DependsOnDisallowedLayer               *event_handlers2.DependsOnDisallowedLayer
	MatchingLayersHandler                  *layers.MatchingLayersHandler
	UncoveredDependentHandler              *event_handlers2.UncoveredDependent
	UnmatchedSkippedViolations             *event_handlers2.UnmatchedSkippedViolations
	ConsoleSubscriber                      *event_handlers2.Console
	ProgressSubscriber                     *event_handlers2.Progress
	VerboseBoolFlag                        *bool
	DebugBoolFlag                          *bool
	Style                                  *formatters.Style
	SymfonyOutput                          *pkg.SymfonyOutput
	TimeStopwatch                          *stopwatch.Stopwatch
	AstMapExtractor                        *ast_map.AstMapExtractor
	InheritanceFlattener                   *flatteners.InheritanceFlattener
	DependencyResolver                     *pkg.DependencyResolver
	TokenResolver                          *services2.TokenResolver
	CollectorResolver                      *collectors_resolvers.CollectorResolver
	LayerResolver                          layers.LayerResolverInterface
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

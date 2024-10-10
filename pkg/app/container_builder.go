package app

import (
	"github.com/KoNekoD/go-deptrac/pkg/ast_map"
	"github.com/KoNekoD/go-deptrac/pkg/collectors_shared"
	"github.com/KoNekoD/go-deptrac/pkg/commands"
	"github.com/KoNekoD/go-deptrac/pkg/configs"
	"github.com/KoNekoD/go-deptrac/pkg/dependencies"
	"github.com/KoNekoD/go-deptrac/pkg/dispatchers"
	"github.com/KoNekoD/go-deptrac/pkg/domain/services"
	"github.com/KoNekoD/go-deptrac/pkg/domain/stopwatch"
	"github.com/KoNekoD/go-deptrac/pkg/domain/utils"
	"github.com/KoNekoD/go-deptrac/pkg/flatteners"
	"github.com/KoNekoD/go-deptrac/pkg/formatters"
	"github.com/KoNekoD/go-deptrac/pkg/hooks"
	"github.com/KoNekoD/go-deptrac/pkg/layers"
	"github.com/KoNekoD/go-deptrac/pkg/parsers"
	"github.com/KoNekoD/go-deptrac/pkg/references"
	"github.com/KoNekoD/go-deptrac/pkg/results"
	"github.com/KoNekoD/go-deptrac/pkg/rules"
	"github.com/KoNekoD/go-deptrac/pkg/subscribers"
	"github.com/KoNekoD/go-deptrac/pkg/tokens"
	"github.com/KoNekoD/go-deptrac/pkg/types"
	"github.com/KoNekoD/go-deptrac/pkg/violations"
)

type ContainerBuilder struct {
	ProjectDirectory                       string
	CacheFile                              *string
	Configuration                          *configs.DeptracConfig
	EventDispatcher                        dispatchers.EventDispatcherInterface
	FileInputCollector                     collectors_shared.InputCollectorInterface
	YmlFileLoader                          *hooks.YmlFileLoader
	Dumper                                 *utils.Dumper
	AstLoader                              *ast_map.AstLoader
	AstFileReferenceFileCache              *ast_map.AstFileReferenceFileCache
	AstFileReferenceDeferredCacheInterface ast_map.AstFileReferenceDeferredCacheInterface
	AstFileReferenceCacheInterface         ast_map.AstFileReferenceCacheInterface
	CacheableFileSubscriber                *subscribers.CacheableFileSubscriber
	AstFileReferenceInMemoryCache          *ast_map.AstFileReferenceInMemoryCache
	TypeResolver                           *types.TypeResolver
	ReferenceExtractors                    []references.ReferenceExtractorInterface
	ParserInterface                        parsers.ParserInterface
	LayerProvider                          *layers.LayerProvider
	EventHelper                            *dispatchers.EventHelper
	AllowDependencyHandler                 *dependencies.AllowDependencyHandler
	DependsOnPrivateLayer                  *subscribers.DependsOnPrivateLayer
	DependsOnInternalToken                 *subscribers.DependsOnInternalToken
	DependsOnDisallowedLayer               *subscribers.DependsOnDisallowedLayer
	MatchingLayersHandler                  *layers.MatchingLayersHandler
	UncoveredDependentHandler              *dependencies.UncoveredDependentHandler
	UnmatchedSkippedViolations             *violations.UnmatchedSkippedViolations
	ConsoleSubscriber                      *subscribers.ConsoleSubscriber
	ProgressSubscriber                     *subscribers.ProgressSubscriber
	VerboseBoolFlag                        *bool
	DebugBoolFlag                          *bool
	Style                                  *formatters.Style
	SymfonyOutput                          *results.SymfonyOutput
	TimeStopwatch                          *stopwatch.Stopwatch
	AstMapExtractor                        *ast_map.AstMapExtractor
	InheritanceFlattener                   *flatteners.InheritanceFlattener
	DependencyResolver                     *dependencies.DependencyResolver
	TokenResolver                          *tokens.TokenResolver
	CollectorResolver                      *collectors_shared.CollectorResolver
	LayerResolver                          layers.LayerResolverInterface
	NikicPhpParser                         *parsers.NikicPhpParser
	CollectorProvider                      *collectors_shared.CollectorProvider
	DependencyLayersAnalyser               *dependencies.DependencyLayersAnalyser
	TokenInLayerAnalyser                   *tokens.TokenInLayerAnalyser
	LayerForTokenAnalyser                  *tokens.LayerForTokenAnalyser
	UnassignedTokenAnalyser                *tokens.UnassignedTokenAnalyser
	LayerDependenciesAnalyser              *dependencies.LayerDependenciesAnalyser
	RulesetUsageAnalyser                   *rules.RulesetUsageAnalyser
	FormatterProvider                      *formatters.FormatterProvider
	FormatterConfiguration                 *formatters.FormatterConfiguration
	AnalyseRunner                          *AnalyseRunner
	AnalyseCommand                         *commands.AnalyseCommand
	NodeNamer                              *services.NodeNamer
	AnalyseOptions                         *rules.AnalyseOptions
}

func NewContainerBuilder(workingDirectory string) *ContainerBuilder {
	return &ContainerBuilder{}
}

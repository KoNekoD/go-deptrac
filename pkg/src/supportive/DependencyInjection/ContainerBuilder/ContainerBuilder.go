package ContainerBuilder

import (
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Config/DeptracConfig"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Layer/LayerProvider"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/analyser/event_helper"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/analyser"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/analyser/event_handler/post_process_event"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/analyser/event_handler/process_event"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/ast"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/ast/parser"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/ast/parser/cache"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/ast/parser/extractors"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/ast/parser/nikic_php_parser"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/dependency"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/dependency/dependency_resolver"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/input_collector"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/layer/collector"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/layer/layer_resolver_interface"
	"github.com/KoNekoD/go-deptrac/pkg/src/supportive/Console/Command/AnalyseCommand"
	"github.com/KoNekoD/go-deptrac/pkg/src/supportive/Console/Command/AnalyseRunner"
	"github.com/KoNekoD/go-deptrac/pkg/src/supportive/Console/Subscriber/ConsoleSubscriber"
	"github.com/KoNekoD/go-deptrac/pkg/src/supportive/Console/Subscriber/ProgressSubscriber"
	"github.com/KoNekoD/go-deptrac/pkg/src/supportive/Console/Symfony/Style"
	"github.com/KoNekoD/go-deptrac/pkg/src/supportive/Console/Symfony/SymfonyOutput"
	"github.com/KoNekoD/go-deptrac/pkg/src/supportive/DependencyInjection/EventDispatcher/EventDispatcherInterface"
	"github.com/KoNekoD/go-deptrac/pkg/src/supportive/File/Dumper"
	"github.com/KoNekoD/go-deptrac/pkg/src/supportive/File/YmlFileLoader"
	"github.com/KoNekoD/go-deptrac/pkg/src/supportive/OutputFormatter/Configuration/FormatterConfiguration"
	"github.com/KoNekoD/go-deptrac/pkg/src/supportive/OutputFormatter/FormatterProvider"
	"github.com/KoNekoD/go-deptrac/pkg/src/supportive/TimeStopwatch"
)

type ContainerBuilder struct {
	ProjectDirectory                       string
	CacheFile                              *string
	Configuration                          *DeptracConfig.DeptracConfig
	EventDispatcher                        util.EventDispatcherInterface
	FileInputCollector                     input_collector.InputCollectorInterface
	YmlFileLoader                          *YmlFileLoader.YmlFileLoader
	Dumper                                 *Dumper.Dumper
	AstLoader                              *ast.AstLoader
	AstFileReferenceFileCache              *cache.AstFileReferenceFileCache
	AstFileReferenceDeferredCacheInterface cache.AstFileReferenceDeferredCacheInterface
	AstFileReferenceCacheInterface         cache.AstFileReferenceCacheInterface
	CacheableFileSubscriber                *cache.CacheableFileSubscriber
	AstFileReferenceInMemoryCache          *cache.AstFileReferenceInMemoryCache
	TypeResolver                           *parser.TypeResolver
	ReferenceExtractors                    []extractors.ReferenceExtractorInterface
	ParserInterface                        parser.ParserInterface
	LayerProvider                          *LayerProvider.LayerProvider
	EventHelper                            *event_helper.EventHelper
	AllowDependencyHandler                 *process_event.AllowDependencyHandler
	DependsOnPrivateLayer                  *process_event.DependsOnPrivateLayer
	DependsOnInternalToken                 *process_event.DependsOnInternalToken
	DependsOnDisallowedLayer               *process_event.DependsOnDisallowedLayer
	MatchingLayersHandler                  *process_event.MatchingLayersHandler
	UncoveredDependentHandler              *process_event.UncoveredDependentHandler
	UnmatchedSkippedViolations             *post_process_event.UnmatchedSkippedViolations
	ConsoleSubscriber                      *ConsoleSubscriber.ConsoleSubscriber
	ProgressSubscriber                     *ProgressSubscriber.ProgressSubscriber
	VerboseBoolFlag                        *bool
	DebugBoolFlag                          *bool
	Style                                  *Style.Style
	SymfonyOutput                          *SymfonyOutput.SymfonyOutput
	TimeStopwatch                          *TimeStopwatch.Stopwatch
	AstMapExtractor                        *ast.AstMapExtractor
	InheritanceFlattener                   *dependency.InheritanceFlattener
	DependencyResolver                     *dependency_resolver.DependencyResolver
	TokenResolver                          *dependency.TokenResolver
	CollectorResolver                      *collector.CollectorResolver
	LayerResolver                          layer_resolver_interface.LayerResolverInterface
	NikicPhpParser                         *nikic_php_parser.NikicPhpParser
	CollectorProvider                      *collector.CollectorProvider
	DependencyLayersAnalyser               *analyser.DependencyLayersAnalyser
	TokenInLayerAnalyser                   *analyser.TokenInLayerAnalyser
	LayerForTokenAnalyser                  *analyser.LayerForTokenAnalyser
	UnassignedTokenAnalyser                *analyser.UnassignedTokenAnalyser
	LayerDependenciesAnalyser              *analyser.LayerDependenciesAnalyser
	RulesetUsageAnalyser                   *analyser.RulesetUsageAnalyser
	FormatterProvider                      *FormatterProvider.FormatterProvider
	FormatterConfiguration                 *FormatterConfiguration.FormatterConfiguration
	AnalyseRunner                          *AnalyseRunner.AnalyseRunner
	AnalyseCommand                         *AnalyseCommand.AnalyseCommand
}

func NewContainerBuilder(workingDirectory string) *ContainerBuilder {
	return &ContainerBuilder{}
}

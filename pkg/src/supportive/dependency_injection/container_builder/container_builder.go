package container_builder

import (
	analyser2 "github.com/KoNekoD/go-deptrac/pkg/src/contract/analyser/event_helper"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/config/deptrac_config"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/layer"
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
	"github.com/KoNekoD/go-deptrac/pkg/src/supportive/console/command"
	"github.com/KoNekoD/go-deptrac/pkg/src/supportive/console/subscriber"
	"github.com/KoNekoD/go-deptrac/pkg/src/supportive/console/symfony"
	"github.com/KoNekoD/go-deptrac/pkg/src/supportive/dependency_injection/event_dispatcher/event_dispatcher_interface"
	"github.com/KoNekoD/go-deptrac/pkg/src/supportive/file"
	"github.com/KoNekoD/go-deptrac/pkg/src/supportive/output_formatter"
	"github.com/KoNekoD/go-deptrac/pkg/src/supportive/output_formatter/configuration"
	"github.com/KoNekoD/go-deptrac/pkg/src/supportive/time_stopwatch"
)

type ContainerBuilder struct {
	ProjectDirectory                       string
	CacheFile                              *string
	Configuration                          *deptrac_config.DeptracConfig
	EventDispatcher                        util.EventDispatcherInterface
	FileInputCollector                     input_collector.InputCollectorInterface
	YmlFileLoader                          *file.YmlFileLoader
	Dumper                                 *file.Dumper
	AstLoader                              *ast.AstLoader
	AstFileReferenceFileCache              *cache.AstFileReferenceFileCache
	AstFileReferenceDeferredCacheInterface cache.AstFileReferenceDeferredCacheInterface
	AstFileReferenceCacheInterface         cache.AstFileReferenceCacheInterface
	CacheableFileSubscriber                *cache.CacheableFileSubscriber
	AstFileReferenceInMemoryCache          *cache.AstFileReferenceInMemoryCache
	TypeResolver                           *parser.TypeResolver
	ReferenceExtractors                    []extractors.ReferenceExtractorInterface
	ParserInterface                        parser.ParserInterface
	LayerProvider                          *layer.LayerProvider
	EventHelper                            *analyser2.EventHelper
	AllowDependencyHandler                 *process_event.AllowDependencyHandler
	DependsOnPrivateLayer                  *process_event.DependsOnPrivateLayer
	DependsOnInternalToken                 *process_event.DependsOnInternalToken
	DependsOnDisallowedLayer               *process_event.DependsOnDisallowedLayer
	MatchingLayersHandler                  *process_event.MatchingLayersHandler
	UncoveredDependentHandler              *process_event.UncoveredDependentHandler
	UnmatchedSkippedViolations             *post_process_event.UnmatchedSkippedViolations
	ConsoleSubscriber                      *subscriber.ConsoleSubscriber
	ProgressSubscriber                     *subscriber.ProgressSubscriber
	VerboseBoolFlag                        *bool
	DebugBoolFlag                          *bool
	Style                                  *symfony.Style
	SymfonyOutput                          *symfony.SymfonyOutput
	TimeStopwatch                          *time_stopwatch.Stopwatch
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
	FormatterProvider                      *output_formatter.FormatterProvider
	FormatterConfiguration                 *configuration.FormatterConfiguration
	AnalyseRunner                          *command.AnalyseRunner
	AnalyseCommand                         *command.AnalyseCommand
}

func NewContainerBuilder(workingDirectory string) *ContainerBuilder {
	return &ContainerBuilder{}
}

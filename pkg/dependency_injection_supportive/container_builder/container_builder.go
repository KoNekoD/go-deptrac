package container_builder

import (
	analyser2 "github.com/KoNekoD/go-deptrac/pkg/analyser_contract/event_helper"
	analyser_core2 "github.com/KoNekoD/go-deptrac/pkg/analyser_core"
	"github.com/KoNekoD/go-deptrac/pkg/analyser_core/event_handler/post_process_event"
	process_event2 "github.com/KoNekoD/go-deptrac/pkg/analyser_core/event_handler/process_event"
	ast_core2 "github.com/KoNekoD/go-deptrac/pkg/ast_core"
	parser2 "github.com/KoNekoD/go-deptrac/pkg/ast_core/parser"
	cache2 "github.com/KoNekoD/go-deptrac/pkg/ast_core/parser/cache"
	"github.com/KoNekoD/go-deptrac/pkg/ast_core/parser/extractors"
	"github.com/KoNekoD/go-deptrac/pkg/ast_core/parser/nikic_php_parser"
	"github.com/KoNekoD/go-deptrac/pkg/ast_core/parser/nikic_php_parser/node_namer"
	"github.com/KoNekoD/go-deptrac/pkg/config_contract/deptrac_config"
	command2 "github.com/KoNekoD/go-deptrac/pkg/console_supportive/command"
	subscriber2 "github.com/KoNekoD/go-deptrac/pkg/console_supportive/subscriber"
	symfony2 "github.com/KoNekoD/go-deptrac/pkg/console_supportive/symfony"
	dependency_core2 "github.com/KoNekoD/go-deptrac/pkg/dependency_core"
	"github.com/KoNekoD/go-deptrac/pkg/dependency_core/dependency_resolver"
	"github.com/KoNekoD/go-deptrac/pkg/dependency_injection_supportive/event_dispatcher/event_dispatcher_interface"
	file_supportive2 "github.com/KoNekoD/go-deptrac/pkg/file_supportive"
	"github.com/KoNekoD/go-deptrac/pkg/input_collector_core"
	"github.com/KoNekoD/go-deptrac/pkg/layer_contract"
	collector2 "github.com/KoNekoD/go-deptrac/pkg/layer_core/collector"
	"github.com/KoNekoD/go-deptrac/pkg/layer_core/layer_resolver_interface"
	"github.com/KoNekoD/go-deptrac/pkg/output_formatter_supportive"
	"github.com/KoNekoD/go-deptrac/pkg/output_formatter_supportive/configuration"
	"github.com/KoNekoD/go-deptrac/pkg/time_stopwatch_supportive"
)

type ContainerBuilder struct {
	ProjectDirectory                       string
	CacheFile                              *string
	Configuration                          *deptrac_config.DeptracConfig
	EventDispatcher                        event_dispatcher_interface.EventDispatcherInterface
	FileInputCollector                     input_collector_core.InputCollectorInterface
	YmlFileLoader                          *file_supportive2.YmlFileLoader
	Dumper                                 *file_supportive2.Dumper
	AstLoader                              *ast_core2.AstLoader
	AstFileReferenceFileCache              *cache2.AstFileReferenceFileCache
	AstFileReferenceDeferredCacheInterface cache2.AstFileReferenceDeferredCacheInterface
	AstFileReferenceCacheInterface         cache2.AstFileReferenceCacheInterface
	CacheableFileSubscriber                *cache2.CacheableFileSubscriber
	AstFileReferenceInMemoryCache          *cache2.AstFileReferenceInMemoryCache
	TypeResolver                           *parser2.TypeResolver
	ReferenceExtractors                    []extractors.ReferenceExtractorInterface
	ParserInterface                        parser2.ParserInterface
	LayerProvider                          *layer_contract.LayerProvider
	EventHelper                            *analyser2.EventHelper
	AllowDependencyHandler                 *process_event2.AllowDependencyHandler
	DependsOnPrivateLayer                  *process_event2.DependsOnPrivateLayer
	DependsOnInternalToken                 *process_event2.DependsOnInternalToken
	DependsOnDisallowedLayer               *process_event2.DependsOnDisallowedLayer
	MatchingLayersHandler                  *process_event2.MatchingLayersHandler
	UncoveredDependentHandler              *process_event2.UncoveredDependentHandler
	UnmatchedSkippedViolations             *post_process_event.UnmatchedSkippedViolations
	ConsoleSubscriber                      *subscriber2.ConsoleSubscriber
	ProgressSubscriber                     *subscriber2.ProgressSubscriber
	VerboseBoolFlag                        *bool
	DebugBoolFlag                          *bool
	Style                                  *symfony2.Style
	SymfonyOutput                          *symfony2.SymfonyOutput
	TimeStopwatch                          *time_stopwatch_supportive.Stopwatch
	AstMapExtractor                        *ast_core2.AstMapExtractor
	InheritanceFlattener                   *dependency_core2.InheritanceFlattener
	DependencyResolver                     *dependency_resolver.DependencyResolver
	TokenResolver                          *dependency_core2.TokenResolver
	CollectorResolver                      *collector2.CollectorResolver
	LayerResolver                          layer_resolver_interface.LayerResolverInterface
	NikicPhpParser                         *nikic_php_parser.NikicPhpParser
	CollectorProvider                      *collector2.CollectorProvider
	DependencyLayersAnalyser               *analyser_core2.DependencyLayersAnalyser
	TokenInLayerAnalyser                   *analyser_core2.TokenInLayerAnalyser
	LayerForTokenAnalyser                  *analyser_core2.LayerForTokenAnalyser
	UnassignedTokenAnalyser                *analyser_core2.UnassignedTokenAnalyser
	LayerDependenciesAnalyser              *analyser_core2.LayerDependenciesAnalyser
	RulesetUsageAnalyser                   *analyser_core2.RulesetUsageAnalyser
	FormatterProvider                      *output_formatter_supportive.FormatterProvider
	FormatterConfiguration                 *configuration.FormatterConfiguration
	AnalyseRunner                          *command2.AnalyseRunner
	AnalyseCommand                         *command2.AnalyseCommand
	NodeNamer                              *node_namer.NodeNamer
	AnalyseOptions                         *command2.AnalyseOptions
}

func NewContainerBuilder(workingDirectory string) *ContainerBuilder {
	return &ContainerBuilder{}
}

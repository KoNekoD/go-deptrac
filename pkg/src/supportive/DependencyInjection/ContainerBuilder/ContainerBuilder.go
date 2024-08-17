package ContainerBuilder

import (
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Analyser/EventHelper"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Config/DeptracConfig"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Layer/LayerProvider"
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
	"github.com/KoNekoD/go-deptrac/pkg/src/core/Ast/Parser/Cache/AstFileReferenceCacheInterface"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/Ast/Parser/Cache/AstFileReferenceDeferredCacheInterface"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/Ast/Parser/Cache/AstFileReferenceFileCache"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/Ast/Parser/Cache/AstFileReferenceInMemoryCache"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/Ast/Parser/Cache/CacheableFileSubscriber"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/Ast/Parser/Extractors/ReferenceExtractorInterface"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/Ast/Parser/NikicPhpParser/NikicPhpParser"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/Ast/Parser/ParserInterface"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/Ast/Parser/TypeResolver"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/Dependency/DependencyResolver"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/Dependency/InheritanceFlattener"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/Dependency/TokenResolver"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/InputCollector/InputCollectorInterface"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/Layer/Collector/CollectorProvider"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/Layer/Collector/CollectorResolver"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/Layer/LayerResolverInterface"
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
	FileInputCollector                     InputCollectorInterface.InputCollectorInterface
	YmlFileLoader                          *YmlFileLoader.YmlFileLoader
	Dumper                                 *Dumper.Dumper
	AstLoader                              *AstLoader.AstLoader
	AstFileReferenceFileCache              *AstFileReferenceFileCache.AstFileReferenceFileCache
	AstFileReferenceDeferredCacheInterface AstFileReferenceDeferredCacheInterface.AstFileReferenceDeferredCacheInterface
	AstFileReferenceCacheInterface         AstFileReferenceCacheInterface.AstFileReferenceCacheInterface
	CacheableFileSubscriber                *CacheableFileSubscriber.CacheableFileSubscriber
	AstFileReferenceInMemoryCache          *AstFileReferenceInMemoryCache.AstFileReferenceInMemoryCache
	TypeResolver                           *TypeResolver.TypeResolver
	ReferenceExtractors                    []ReferenceExtractorInterface.ReferenceExtractorInterface
	ParserInterface                        ParserInterface.ParserInterface
	LayerProvider                          *LayerProvider.LayerProvider
	EventHelper                            *EventHelper.EventHelper
	AllowDependencyHandler                 *AllowDependencyHandler.AllowDependencyHandler
	DependsOnPrivateLayer                  *DependsOnPrivateLayer.DependsOnPrivateLayer
	DependsOnInternalToken                 *DependsOnInternalToken.DependsOnInternalToken
	DependsOnDisallowedLayer               *DependsOnDisallowedLayer.DependsOnDisallowedLayer
	MatchingLayersHandler                  *MatchingLayersHandler.MatchingLayersHandler
	UncoveredDependentHandler              *UncoveredDependentHandler.UncoveredDependentHandler
	UnmatchedSkippedViolations             *UnmatchedSkippedViolations.UnmatchedSkippedViolations
	ConsoleSubscriber                      *ConsoleSubscriber.ConsoleSubscriber
	ProgressSubscriber                     *ProgressSubscriber.ProgressSubscriber
	VerboseBoolFlag                        *bool
	DebugBoolFlag                          *bool
	Style                                  *Style.Style
	SymfonyOutput                          *SymfonyOutput.SymfonyOutput
	TimeStopwatch                          *TimeStopwatch.Stopwatch
	AstMapExtractor                        *AstMapExtractor.AstMapExtractor
	InheritanceFlattener                   *InheritanceFlattener.InheritanceFlattener
	DependencyResolver                     *DependencyResolver.DependencyResolver
	TokenResolver                          *TokenResolver.TokenResolver
	CollectorResolver                      *CollectorResolver.CollectorResolver
	LayerResolver                          LayerResolverInterface.LayerResolverInterface
	NikicPhpParser                         *NikicPhpParser.NikicPhpParser
	CollectorProvider                      *CollectorProvider.CollectorProvider
	DependencyLayersAnalyser               *DependencyLayersAnalyser.DependencyLayersAnalyser
	TokenInLayerAnalyser                   *TokenInLayerAnalyser.TokenInLayerAnalyser
	LayerForTokenAnalyser                  *LayerForTokenAnalyser.LayerForTokenAnalyser
	UnassignedTokenAnalyser                *UnassignedTokenAnalyser.UnassignedTokenAnalyser
	LayerDependenciesAnalyser              *LayerDependenciesAnalyser.LayerDependenciesAnalyser
	RulesetUsageAnalyser                   *RulesetUsageAnalyser.RulesetUsageAnalyser
	FormatterProvider                      *FormatterProvider.FormatterProvider
	FormatterConfiguration                 *FormatterConfiguration.FormatterConfiguration
	AnalyseRunner                          *AnalyseRunner.AnalyseRunner
	AnalyseCommand                         *AnalyseCommand.AnalyseCommand
}

func NewContainerBuilder(workingDirectory string) *ContainerBuilder {
	return &ContainerBuilder{}
}

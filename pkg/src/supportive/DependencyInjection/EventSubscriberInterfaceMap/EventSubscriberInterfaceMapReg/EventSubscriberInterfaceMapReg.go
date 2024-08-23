package EventSubscriberInterfaceMapReg

import (
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Ast/AstFileAnalysedEvent"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Ast/AstFileSyntaxErrorEvent"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Ast/PostCreateAstMapEvent"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Ast/PreCreateAstMapEvent"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Dependency/PostEmitEvent"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Dependency/PostFlattenEvent"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Dependency/PreEmitEvent"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Dependency/PreFlattenEvent"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/analyser/post_process_event"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/analyser/process_event"
	"github.com/KoNekoD/go-deptrac/pkg/src/supportive/Console/Subscriber/ConsoleSubscriber"
	"github.com/KoNekoD/go-deptrac/pkg/src/supportive/Console/Subscriber/ProgressSubscriber"
	"github.com/KoNekoD/go-deptrac/pkg/src/supportive/DependencyInjection/EventSubscriberDefaultPriority"
	"github.com/KoNekoD/go-deptrac/pkg/src/supportive/DependencyInjection/EventSubscriberInterface"
	EventSubscriberInterfaceMap2 "github.com/KoNekoD/go-deptrac/pkg/src/supportive/DependencyInjection/EventSubscriberInterfaceMap"
	"github.com/elliotchance/orderedmap/v2"
	"reflect"
)

func RegForAnalyseCommand(consoleSubscriber *ConsoleSubscriber.ConsoleSubscriber, progressSubscriber *ProgressSubscriber.ProgressSubscriber, withProgress bool) {
	processEvent := &process_event.ProcessEvent{}
	postProcessEvent := &post_process_event.PostProcessEvent{}
	preCreateAstMapEvent := &PreCreateAstMapEvent.PreCreateAstMapEvent{}
	postCreateAstMapEvent := &PostCreateAstMapEvent.PostCreateAstMapEvent{}
	astFileAnalysedEvent := &AstFileAnalysedEvent.AstFileAnalysedEvent{}
	astFileSyntaxErrorEvent := &AstFileSyntaxErrorEvent.AstFileSyntaxErrorEvent{}
	preEmitEvent := &PreEmitEvent.PreEmitEvent{}
	postEmitEvent := &PostEmitEvent.PostEmitEvent{}
	preFlattenEvent := &PreFlattenEvent.PreFlattenEvent{}
	postFlattenEvent := &PostFlattenEvent.PostFlattenEvent{}

	Reg(preCreateAstMapEvent, consoleSubscriber, EventSubscriberInterfaceMap.DefaultPriority)
	Reg(postCreateAstMapEvent, consoleSubscriber, EventSubscriberInterfaceMap.DefaultPriority)
	Reg(processEvent, consoleSubscriber, EventSubscriberInterfaceMap.DefaultPriority)
	Reg(postProcessEvent, consoleSubscriber, EventSubscriberInterfaceMap.DefaultPriority)
	Reg(astFileAnalysedEvent, consoleSubscriber, EventSubscriberInterfaceMap.DefaultPriority)
	Reg(astFileSyntaxErrorEvent, consoleSubscriber, EventSubscriberInterfaceMap.DefaultPriority)
	Reg(preEmitEvent, consoleSubscriber, EventSubscriberInterfaceMap.DefaultPriority)
	Reg(postEmitEvent, consoleSubscriber, EventSubscriberInterfaceMap.DefaultPriority)
	Reg(preFlattenEvent, consoleSubscriber, EventSubscriberInterfaceMap.DefaultPriority)
	Reg(postFlattenEvent, consoleSubscriber, EventSubscriberInterfaceMap.DefaultPriority)

	if withProgress {
		Reg(preCreateAstMapEvent, progressSubscriber, EventSubscriberInterfaceMap.DefaultPriority)
		Reg(postCreateAstMapEvent, progressSubscriber, 1)
		Reg(astFileAnalysedEvent, progressSubscriber, EventSubscriberInterfaceMap.DefaultPriority)
	}
}

func Reg(event interface{}, sub EventSubscriberInterface.EventSubscriberInterface, priority int) {
	eventTypeof := reflect.TypeOf(event)
	eventType := eventTypeof.String()

	// Get or create event type row
	e, ok := EventSubscriberInterfaceMap2.Map.Get(eventType)
	if !ok {
		e = orderedmap.NewOrderedMap[int, []EventSubscriberInterface.EventSubscriberInterface]()
		EventSubscriberInterfaceMap2.Map.Set(eventType, e)
	}

	// Get or create priority column
	subs, ok := e.Get(priority)
	if !ok {
		subs = []EventSubscriberInterface.EventSubscriberInterface{}
	}

	subs = append(subs, sub)

	e.Set(priority, subs)
}

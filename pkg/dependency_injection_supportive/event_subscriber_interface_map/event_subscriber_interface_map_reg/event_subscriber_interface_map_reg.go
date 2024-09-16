package event_subscriber_interface_map_reg

import (
	"github.com/KoNekoD/go-deptrac/pkg/analyser_contract/post_process_event"
	"github.com/KoNekoD/go-deptrac/pkg/analyser_contract/process_event"
	ast_contract2 "github.com/KoNekoD/go-deptrac/pkg/ast_contract"
	subscriber2 "github.com/KoNekoD/go-deptrac/pkg/console_supportive/subscriber"
	dependency_contract2 "github.com/KoNekoD/go-deptrac/pkg/dependency_contract"
	"github.com/KoNekoD/go-deptrac/pkg/dependency_injection_supportive/event_subscriber_default_priority"
	"github.com/KoNekoD/go-deptrac/pkg/dependency_injection_supportive/event_subscriber_interface"
	EventSubscriberInterfaceMap2 "github.com/KoNekoD/go-deptrac/pkg/dependency_injection_supportive/event_subscriber_interface_map"
	"github.com/elliotchance/orderedmap/v2"
	"reflect"
)

func RegForAnalyseCommand(consoleSubscriber *subscriber2.ConsoleSubscriber, progressSubscriber *subscriber2.ProgressSubscriber, withProgress bool) {
	processEvent := &process_event.ProcessEvent{}
	postProcessEvent := &post_process_event.PostProcessEvent{}
	preCreateAstMapEvent := &ast_contract2.PreCreateAstMapEvent{}
	postCreateAstMapEvent := &ast_contract2.PostCreateAstMapEvent{}
	astFileAnalysedEvent := &ast_contract2.AstFileAnalysedEvent{}
	astFileSyntaxErrorEvent := &ast_contract2.AstFileSyntaxErrorEvent{}
	preEmitEvent := &dependency_contract2.PreEmitEvent{}
	postEmitEvent := &dependency_contract2.PostEmitEvent{}
	preFlattenEvent := &dependency_contract2.PreFlattenEvent{}
	postFlattenEvent := &dependency_contract2.PostFlattenEvent{}

	Reg(preCreateAstMapEvent, consoleSubscriber, EventSubscriberInterfaceMap.EventSubscriberInterfaceMap.DefaultPriority)
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

func Reg(event interface{}, sub event_subscriber_interface.EventSubscriberInterface, priority int) {
	eventTypeof := reflect.TypeOf(event)
	eventType := eventTypeof.String()

	// Get or create event type row
	e, ok := EventSubscriberInterfaceMap2.Map.Get(eventType)
	if !ok {
		e = orderedmap.NewOrderedMap[int, []event_subscriber_interface.EventSubscriberInterface]()
		EventSubscriberInterfaceMap2.Map.Set(eventType, e)
	}

	// Get or create priority column
	subs, ok := e.Get(priority)
	if !ok {
		subs = []event_subscriber_interface.EventSubscriberInterface{}
	}

	subs = append(subs, sub)

	e.Set(priority, subs)
}

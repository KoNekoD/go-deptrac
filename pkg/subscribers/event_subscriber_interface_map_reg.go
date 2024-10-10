package subscribers

import (
	"github.com/KoNekoD/go-deptrac/pkg/ast_map"
	"github.com/KoNekoD/go-deptrac/pkg/emitters"
	"github.com/KoNekoD/go-deptrac/pkg/events"
	"github.com/elliotchance/orderedmap/v2"
	"reflect"
)

func RegForAnalyseCommand(consoleSubscriber *ConsoleSubscriber, progressSubscriber *ProgressSubscriber, withProgress bool) {
	processEvent := &events.ProcessEvent{}
	postProcessEvent := &events.PostProcessEvent{}
	preCreateAstMapEvent := &ast_map.PreCreateAstMapEvent{}
	postCreateAstMapEvent := &ast_map.PostCreateAstMapEvent{}
	astFileAnalysedEvent := &events.AstFileAnalysedEvent{}
	astFileSyntaxErrorEvent := &ast_map.AstFileSyntaxErrorEvent{}
	preEmitEvent := &emitters.PreEmitEvent{}
	postEmitEvent := &emitters.PostEmitEvent{}
	preFlattenEvent := &events.PreFlattenEvent{}
	postFlattenEvent := &events.PostFlattenEvent{}

	Reg(preCreateAstMapEvent, consoleSubscriber, DefaultPriority)
	Reg(postCreateAstMapEvent, consoleSubscriber, DefaultPriority)
	Reg(processEvent, consoleSubscriber, DefaultPriority)
	Reg(postProcessEvent, consoleSubscriber, DefaultPriority)
	Reg(astFileAnalysedEvent, consoleSubscriber, DefaultPriority)
	Reg(astFileSyntaxErrorEvent, consoleSubscriber, DefaultPriority)
	Reg(preEmitEvent, consoleSubscriber, DefaultPriority)
	Reg(postEmitEvent, consoleSubscriber, DefaultPriority)
	Reg(preFlattenEvent, consoleSubscriber, DefaultPriority)
	Reg(postFlattenEvent, consoleSubscriber, DefaultPriority)

	if withProgress {
		Reg(preCreateAstMapEvent, progressSubscriber, DefaultPriority)
		Reg(postCreateAstMapEvent, progressSubscriber, 1)
		Reg(astFileAnalysedEvent, progressSubscriber, DefaultPriority)
	}
}

func Reg(event interface{}, sub EventSubscriberInterface, priority int) {
	eventTypeof := reflect.TypeOf(event)
	eventType := eventTypeof.String()

	// Get or create event type row
	e, ok := Map.Get(eventType)
	if !ok {
		e = orderedmap.NewOrderedMap[int, []EventSubscriberInterface]()
		Map.Set(eventType, e)
	}

	// Get or create priority column
	subs, ok := e.Get(priority)
	if !ok {
		subs = []EventSubscriberInterface{}
	}

	subs = append(subs, sub)

	e.Set(priority, subs)
}

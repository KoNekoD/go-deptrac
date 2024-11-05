package event_dispatchers

import (
	"fmt"
	"github.com/elliotchance/orderedmap/v2"
	"reflect"
	"slices"
)

type EventDispatcherInterface interface {
	DispatchEvent(event interface{}) error
}

type EventDispatcher struct {
	isDebug bool
}

func NewEventDispatcher(isDebug bool) EventDispatcherInterface {
	return &EventDispatcher{isDebug: isDebug}
}

func (d *EventDispatcher) DispatchEvent(event interface{}) error {
	typeName := reflect.TypeOf(event).String()

	handlers, ok := Map.Get(typeName)

	if !ok {
		return nil // No subscribers registered for this event
	}

	stop := false
	stopPropagation := func() {
		stop = true
	}

	subscribersPriorities := handlers.Keys()

	// Sort high to low priority
	slices.Sort(subscribersPriorities)
	slices.Reverse(subscribersPriorities)

	for _, priority := range subscribersPriorities {
		subscribersByPriority, okGet := handlers.Get(priority)
		if !okGet {
			continue
		}

		for _, subscriber := range subscribersByPriority {
			if stop {
				break
			}

			subscriberName := reflect.TypeOf(subscriber).String()

			if d.isDebug {
				fmt.Println("calling:", typeName, priority, subscriberName)
			}

			err := subscriber.HandleEvent(event, stopPropagation)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

type EventHandlerInterface interface {
	HandleEvent(rawEvent interface{}, stopPropagation func()) error
}

var Map *orderedmap.OrderedMap[string, *orderedmap.OrderedMap[int, []EventHandlerInterface]]

const DefaultPriority = 0

func Reg(event interface{}, sub EventHandlerInterface, priority int) {
	eventTypeof := reflect.TypeOf(event)
	eventType := eventTypeof.String()

	// Get or create event type row
	e, ok := Map.Get(eventType)
	if !ok {
		e = orderedmap.NewOrderedMap[int, []EventHandlerInterface]()
		Map.Set(eventType, e)
	}

	// Get or create priority column
	subs, ok := e.Get(priority)
	if !ok {
		subs = []EventHandlerInterface{}
	}

	subs = append(subs, sub)

	e.Set(priority, subs)
}

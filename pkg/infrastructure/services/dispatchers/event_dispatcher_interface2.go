package dispatchers

import (
	"fmt"
	"github.com/KoNekoD/go-deptrac/pkg/application/event_handlers"
	"reflect"
	"slices"
)

type EventDispatcher struct {
	isDebug bool
}

func NewEventDispatcher(isDebug bool) EventDispatcherInterface {
	return &EventDispatcher{
		isDebug: isDebug,
	}
}

func (d *EventDispatcher) DispatchEvent(event interface{}) error {
	typeName := reflect.TypeOf(event).String()

	handlers, ok := event_handlers.Map.Get(typeName)

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
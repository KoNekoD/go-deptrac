package EventDispatcher

import (
	"fmt"
	util "github.com/KoNekoD/go-deptrac/pkg/src/supportive/DependencyInjection/EventDispatcher/EventDispatcherInterface"
	"github.com/KoNekoD/go-deptrac/pkg/src/supportive/DependencyInjection/EventSubscriberInterfaceMap"
	"reflect"
	"slices"
)

type EventDispatcher struct {
	isDebug bool
}

func NewEventDispatcher(isDebug bool) util.EventDispatcherInterface {
	return &EventDispatcher{
		isDebug: isDebug,
	}
}

func (d *EventDispatcher) DispatchEvent(event interface{}) error {
	typeName := reflect.TypeOf(event).String()

	subscribers, ok := EventSubscriberInterfaceMap.Map.Get(typeName)

	if !ok {
		return nil // No subscribers registered for this event
	}

	stop := false
	stopPropagation := func() {
		stop = true
	}

	subscribersPriorities := subscribers.Keys()

	// Sort high to low priority
	slices.Sort(subscribersPriorities)
	slices.Reverse(subscribersPriorities)

	for _, priority := range subscribersPriorities {
		subscribersByPriority, okGet := subscribers.Get(priority)
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

			err := subscriber.InvokeEventSubscriber(event, stopPropagation)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

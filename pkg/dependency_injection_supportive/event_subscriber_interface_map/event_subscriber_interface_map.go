package event_subscriber_interface_map

import (
	"github.com/KoNekoD/go-deptrac/pkg/dependency_injection_supportive/event_subscriber_interface"
	"github.com/elliotchance/orderedmap/v2"
)

var Map *orderedmap.OrderedMap[string, *orderedmap.OrderedMap[int, []event_subscriber_interface.EventSubscriberInterface]]

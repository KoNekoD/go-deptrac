package event_subscriber_interface_map

import (
	"github.com/KoNekoD/go-deptrac/pkg/src/supportive/dependency_injection/event_subscriber_interface"
	"github.com/elliotchance/orderedmap/v2"
)

var Map *orderedmap.OrderedMap[string, *orderedmap.OrderedMap[int, []event_subscriber_interface.EventSubscriberInterface]]

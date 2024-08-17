package EventSubscriberInterfaceMap

import (
	"github.com/KoNekoD/go-deptrac/pkg/src/supportive/DependencyInjection/EventSubscriberInterface"
	"github.com/elliotchance/orderedmap/v2"
)

var Map *orderedmap.OrderedMap[string, *orderedmap.OrderedMap[int, []EventSubscriberInterface.EventSubscriberInterface]]

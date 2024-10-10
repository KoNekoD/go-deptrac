package subscribers

import (
	"github.com/elliotchance/orderedmap/v2"
)

var Map *orderedmap.OrderedMap[string, *orderedmap.OrderedMap[int, []EventSubscriberInterface]]

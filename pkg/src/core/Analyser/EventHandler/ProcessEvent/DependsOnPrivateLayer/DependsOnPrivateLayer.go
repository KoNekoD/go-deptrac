package DependsOnPrivateLayer

import (
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Analyser/EventHelper"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Analyser/ProcessEvent"
)

type DependsOnPrivateLayer struct {
	eventHelper *EventHelper.EventHelper
}

func NewDependsOnPrivateLayer(eventHelper *EventHelper.EventHelper) *DependsOnPrivateLayer {
	return &DependsOnPrivateLayer{eventHelper: eventHelper}
}

func (d *DependsOnPrivateLayer) InvokeEventSubscriber(rawEvent interface{}, stopPropagation func()) error {
	event := rawEvent.(*ProcessEvent.ProcessEvent)
	ruleset := event.GetResult()
	for dependentLayer, isPublic := range event.DependentLayers {
		if event.DependerLayer == dependentLayer && !isPublic {
			d.eventHelper.AddSkippableViolation(event, ruleset, dependentLayer, d)
			stopPropagation()
		}
	}
	return nil
}

func (d *DependsOnPrivateLayer) RuleName() string {
	return "DependsOnPrivateLayer"
}

func (d *DependsOnPrivateLayer) RuleDescription() string {
	return "You are depending on a part of a layer that was defined as private to that layer and you are not part of that layer."
}

package process_event

import (
	"github.com/KoNekoD/go-deptrac/pkg/contract/analyser/event_helper"
	"github.com/KoNekoD/go-deptrac/pkg/contract/analyser/process_event"
)

type DependsOnPrivateLayer struct {
	eventHelper *event_helper.EventHelper
}

func NewDependsOnPrivateLayer(eventHelper *event_helper.EventHelper) *DependsOnPrivateLayer {
	return &DependsOnPrivateLayer{eventHelper: eventHelper}
}

func (d *DependsOnPrivateLayer) InvokeEventSubscriber(rawEvent interface{}, stopPropagation func()) error {
	event := rawEvent.(*process_event.ProcessEvent)
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

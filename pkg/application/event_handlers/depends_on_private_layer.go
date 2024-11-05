package event_handlers

import (
	"github.com/KoNekoD/go-deptrac/pkg/application/services"
	"github.com/KoNekoD/go-deptrac/pkg/domain/events"
)

type DependsOnPrivateLayer struct {
	eventHelper *services.EventHelper
}

func NewDependsOnPrivateLayer(eventHelper *services.EventHelper) *DependsOnPrivateLayer {
	return &DependsOnPrivateLayer{eventHelper: eventHelper}
}

func (d *DependsOnPrivateLayer) HandleEvent(rawEvent interface{}, stopPropagation func()) error {
	event := rawEvent.(*events.ProcessEvent)
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
	return "You are depending on a part of a layer_contract that was defined as private to that layer_contract and you are not part of that layer_contract."
}

package DependsOnDisallowedLayer

import (
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Analyser/EventHelper"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Analyser/ProcessEvent"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Result/Error"
	"github.com/KoNekoD/go-deptrac/pkg/util"
)

type DependsOnDisallowedLayer struct {
	eventHelper *EventHelper.EventHelper
}

func NewDependsOnDisallowedLayer(eventHelper *EventHelper.EventHelper) *DependsOnDisallowedLayer {
	return &DependsOnDisallowedLayer{eventHelper: eventHelper}
}

func (d *DependsOnDisallowedLayer) InvokeEventSubscriber(rawEvent interface{}, stopPropagation func()) error {
	event := rawEvent.(*ProcessEvent.ProcessEvent)

	ruleset := event.GetResult()
	allowedLayers, err := d.eventHelper.LayerProvider.GetAllowedLayers(event.DependerLayer)
	if err != nil {
		ruleset.AddError(Error.NewError(err.Error()))
		stopPropagation()
		return nil
	}
	for dependentLayer := range event.DependentLayers {
		if !util.InArray(dependentLayer, allowedLayers) {
			d.eventHelper.AddSkippableViolation(event, ruleset, dependentLayer, d)
			stopPropagation()
		}
	}

	return nil
}

func (d *DependsOnDisallowedLayer) RuleName() string {
	return "DependsOnDisallowedLayer"
}

func (d *DependsOnDisallowedLayer) RuleDescription() string {
	return "You are depending on token that is a part of a layer that you are not allowed to depend on."
}

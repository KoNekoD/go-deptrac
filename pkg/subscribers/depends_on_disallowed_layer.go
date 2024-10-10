package subscribers

import (
	"github.com/KoNekoD/go-deptrac/pkg/domain/apperrors"
	"github.com/KoNekoD/go-deptrac/pkg/domain/utils"
	"github.com/KoNekoD/go-deptrac/pkg/events"
)

type DependsOnDisallowedLayer struct {
	eventHelper *events.EventHelper
}

func NewDependsOnDisallowedLayer(eventHelper *events.EventHelper) *DependsOnDisallowedLayer {
	return &DependsOnDisallowedLayer{eventHelper: eventHelper}
}

func (d *DependsOnDisallowedLayer) InvokeEventSubscriber(rawEvent interface{}, stopPropagation func()) error {
	event := rawEvent.(*events.ProcessEvent)

	ruleset := event.GetResult()
	allowedLayers, err := d.eventHelper.LayerProvider.GetAllowedLayers(event.DependerLayer)
	if err != nil {
		ruleset.AddError(apperrors.NewError(err.Error()))
		stopPropagation()
		return nil
	}
	for dependentLayer := range event.DependentLayers {
		if !utils.InArray(dependentLayer, allowedLayers) {
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
	return "You are depending on token that is a part of a layer_contract that you are not allowed to depend on."
}
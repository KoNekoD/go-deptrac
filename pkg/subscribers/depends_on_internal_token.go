package subscribers

import (
	"github.com/KoNekoD/go-deptrac/pkg/dispatchers"
	"github.com/KoNekoD/go-deptrac/pkg/domain/dtos"
	"github.com/KoNekoD/go-deptrac/pkg/domain/dtos/tokens_references"
	"github.com/KoNekoD/go-deptrac/pkg/events"
)

type DependsOnInternalToken struct {
	eventHelper *dispatchers.EventHelper
	internalTag *string
}

func NewDependsOnInternalToken(eventHelper *dispatchers.EventHelper, analyser *dtos.AnalyserConfig) *DependsOnInternalToken {
	return &DependsOnInternalToken{eventHelper: eventHelper, internalTag: analyser.InternalTag}
}

func (d *DependsOnInternalToken) InvokeEventSubscriber(rawEvent interface{}, stopPropagation func()) error {
	event := rawEvent.(*events.ProcessEvent)

	ruleset := event.GetResult()
	for dependentLayer := range event.DependentLayers {
		if event.DependerLayer != dependentLayer {
			if ref, ok := event.DependentReference.(*tokens_references.ClassLikeReference); ok {
				isInternal := ref.HasTag("@deptrac-internal")
				if !isInternal && nil != d.internalTag {
					isInternal = ref.HasTag(*d.internalTag)
				}

				if isInternal {
					d.eventHelper.AddSkippableViolation(event, ruleset, dependentLayer, d)
					stopPropagation()
				}
			}
		}
	}
	return nil
}

func (d *DependsOnInternalToken) RuleName() string {
	return "DependsOnInternalToken"
}

func (d *DependsOnInternalToken) RuleDescription() string {
	return "You are depending on a token that is internal to the layer_contract and you are not part of that layer_contract."
}

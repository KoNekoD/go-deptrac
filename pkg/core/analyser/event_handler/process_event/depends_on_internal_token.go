package process_event

import (
	"github.com/KoNekoD/go-deptrac/pkg/contract/analyser/event_helper"
	"github.com/KoNekoD/go-deptrac/pkg/contract/analyser/process_event"
	"github.com/KoNekoD/go-deptrac/pkg/contract/config"
	"github.com/KoNekoD/go-deptrac/pkg/core/ast/ast_map"
)

type DependsOnInternalToken struct {
	eventHelper *event_helper.EventHelper
	internalTag *string
}

func NewDependsOnInternalToken(eventHelper *event_helper.EventHelper, analyser *config.AnalyserConfig) *DependsOnInternalToken {
	return &DependsOnInternalToken{eventHelper: eventHelper, internalTag: analyser.InternalTag}
}

func (d *DependsOnInternalToken) InvokeEventSubscriber(rawEvent interface{}, stopPropagation func()) error {
	event := rawEvent.(*process_event.ProcessEvent)

	ruleset := event.GetResult()
	for dependentLayer := range event.DependentLayers {
		if event.DependerLayer != dependentLayer {
			if ref, ok := event.DependentReference.(*ast_map.ClassLikeReference); ok {
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
	return "You are depending on a token that is internal to the layer and you are not part of that layer."
}
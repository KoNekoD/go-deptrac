package DependsOnInternalToken

import (
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Analyser/EventHelper"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Analyser/ProcessEvent"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Config/AnalyserConfig"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/Ast/AstMap"
)

type DependsOnInternalToken struct {
	eventHelper *EventHelper.EventHelper
	internalTag *string
}

func NewDependsOnInternalToken(eventHelper *EventHelper.EventHelper, analyser *AnalyserConfig.AnalyserConfig) *DependsOnInternalToken {
	return &DependsOnInternalToken{eventHelper: eventHelper, internalTag: analyser.InternalTag}
}

func (d *DependsOnInternalToken) InvokeEventSubscriber(rawEvent interface{}, stopPropagation func()) error {
	event := rawEvent.(*ProcessEvent.ProcessEvent)

	ruleset := event.GetResult()
	for dependentLayer := range event.DependentLayers {
		if event.DependerLayer != dependentLayer {
			if ref, ok := event.DependentReference.(*AstMap.ClassLikeReference); ok {
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

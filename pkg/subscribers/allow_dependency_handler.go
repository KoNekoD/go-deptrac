package subscribers

import (
	"github.com/KoNekoD/go-deptrac/pkg/domain/dtos/analysis_results/violations_rules"
	"github.com/KoNekoD/go-deptrac/pkg/events"
)

type AllowDependencyHandler struct{}

func NewAllowDependencyHandler() *AllowDependencyHandler {
	return &AllowDependencyHandler{}
}

func (AllowDependencyHandler) InvokeEventSubscriber(rawEvent interface{}, stopPropagation func()) error {
	event := rawEvent.(*events.ProcessEvent)

	ruleset := event.GetResult()
	for dependentLayer := range event.DependentLayers {
		ruleset.AddRule(violations_rules.NewAllowed(event.Dependency, event.DependerLayer, dependentLayer))
		stopPropagation()
	}

	return nil
}

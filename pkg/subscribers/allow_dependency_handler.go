package subscribers

import (
	"github.com/KoNekoD/go-deptrac/pkg/events"
	"github.com/KoNekoD/go-deptrac/pkg/rules"
)

type AllowDependencyHandler struct{}

func NewAllowDependencyHandler() *AllowDependencyHandler {
	return &AllowDependencyHandler{}
}

func (AllowDependencyHandler) InvokeEventSubscriber(rawEvent interface{}, stopPropagation func()) error {
	event := rawEvent.(*events.ProcessEvent)

	ruleset := event.GetResult()
	for dependentLayer := range event.DependentLayers {
		ruleset.AddRule(rules.NewAllowed(event.Dependency, event.DependerLayer, dependentLayer))
		stopPropagation()
	}

	return nil
}

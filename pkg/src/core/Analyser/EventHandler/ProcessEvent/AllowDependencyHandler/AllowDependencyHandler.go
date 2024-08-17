package AllowDependencyHandler

import (
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Analyser/ProcessEvent"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Result/Allowed"
)

type AllowDependencyHandler struct{}

func NewAllowDependencyHandler() *AllowDependencyHandler {
	return &AllowDependencyHandler{}
}

func (AllowDependencyHandler) InvokeEventSubscriber(rawEvent interface{}, stopPropagation func()) error {
	event := rawEvent.(*ProcessEvent.ProcessEvent)

	ruleset := event.GetResult()
	for dependentLayer := range event.DependentLayers {
		ruleset.AddRule(Allowed.NewAllowed(event.Dependency, event.DependerLayer, dependentLayer))
		stopPropagation()
	}

	return nil
}

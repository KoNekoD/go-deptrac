package process_event

import (
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/analyser/process_event"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/result"
)

type AllowDependencyHandler struct{}

func NewAllowDependencyHandler() *AllowDependencyHandler {
	return &AllowDependencyHandler{}
}

func (AllowDependencyHandler) InvokeEventSubscriber(rawEvent interface{}, stopPropagation func()) error {
	event := rawEvent.(*process_event.ProcessEvent)

	ruleset := event.GetResult()
	for dependentLayer := range event.DependentLayers {
		ruleset.AddRule(result.NewAllowed(event.Dependency, event.DependerLayer, dependentLayer))
		stopPropagation()
	}

	return nil
}

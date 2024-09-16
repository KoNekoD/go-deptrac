package process_event

import (
	"github.com/KoNekoD/go-deptrac/pkg/analyser_contract/process_event"
	"github.com/KoNekoD/go-deptrac/pkg/result_contract"
)

type AllowDependencyHandler struct{}

func NewAllowDependencyHandler() *AllowDependencyHandler {
	return &AllowDependencyHandler{}
}

func (AllowDependencyHandler) InvokeEventSubscriber(rawEvent interface{}, stopPropagation func()) error {
	event := rawEvent.(*process_event.ProcessEvent)

	ruleset := event.GetResult()
	for dependentLayer := range event.DependentLayers {
		ruleset.AddRule(result_contract.NewAllowed(event.Dependency, event.DependerLayer, dependentLayer))
		stopPropagation()
	}

	return nil
}

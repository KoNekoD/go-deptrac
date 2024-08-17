package MatchingLayersHandler

import "github.com/KoNekoD/go-deptrac/pkg/src/contract/Analyser/ProcessEvent"

type MatchingLayersHandler struct{}

func NewMatchingLayersHandler() *MatchingLayersHandler {
	return &MatchingLayersHandler{}
}

func (m *MatchingLayersHandler) InvokeEventSubscriber(rawEvent interface{}, stopPropagation func()) error {
	event := rawEvent.(*ProcessEvent.ProcessEvent)
	for dependeeLayer := range event.DependentLayers {
		if event.DependerLayer != dependeeLayer {
			return nil
		}

		// For empty dependee layers see UncoveredDependeeHandler
		stopPropagation()
	}
	return nil
}

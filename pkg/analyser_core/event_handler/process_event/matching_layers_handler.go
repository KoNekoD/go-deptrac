package process_event

import (
	"github.com/KoNekoD/go-deptrac/pkg/analyser_contract/process_event"
)

type MatchingLayersHandler struct{}

func NewMatchingLayersHandler() *MatchingLayersHandler {
	return &MatchingLayersHandler{}
}

func (m *MatchingLayersHandler) InvokeEventSubscriber(rawEvent interface{}, stopPropagation func()) error {
	event := rawEvent.(*process_event.ProcessEvent)
	for dependeeLayer := range event.DependentLayers {
		if event.DependerLayer != dependeeLayer {
			return nil
		}

		// For empty dependee layers see UncoveredDependeeHandler
		stopPropagation()
	}
	return nil
}

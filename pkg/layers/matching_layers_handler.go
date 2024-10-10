package layers

import (
	"github.com/KoNekoD/go-deptrac/pkg/events"
)

type MatchingLayersHandler struct{}

func NewMatchingLayersHandler() *MatchingLayersHandler {
	return &MatchingLayersHandler{}
}

func (m *MatchingLayersHandler) InvokeEventSubscriber(rawEvent interface{}, stopPropagation func()) error {
	event := rawEvent.(*events.ProcessEvent)
	for dependeeLayer := range event.DependentLayers {
		if event.DependerLayer != dependeeLayer {
			return nil
		}

		// For empty dependee layers see UncoveredDependeeHandler
		stopPropagation()
	}
	return nil
}
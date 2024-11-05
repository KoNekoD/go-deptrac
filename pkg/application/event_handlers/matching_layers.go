package event_handlers

import (
	"github.com/KoNekoD/go-deptrac/pkg/domain/events"
)

type MatchingLayers struct{}

func NewMatchingLayers() *MatchingLayers {
	return &MatchingLayers{}
}

func (m *MatchingLayers) HandleEvent(rawEvent interface{}, stopPropagation func()) error {
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

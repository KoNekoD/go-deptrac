package UnmatchedSkippedViolations

import (
	"fmt"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Analyser/EventHelper"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Analyser/PostProcessEvent"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Result/Error"
)

type UnmatchedSkippedViolations struct {
	eventHelper *EventHelper.EventHelper
}

func NewUnmatchedSkippedViolations(eventHelper *EventHelper.EventHelper) *UnmatchedSkippedViolations {
	return &UnmatchedSkippedViolations{eventHelper: eventHelper}
}

func (u *UnmatchedSkippedViolations) InvokeEventSubscriber(rawEvent interface{}, stopPropagation func()) error {
	event := rawEvent.(*PostProcessEvent.PostProcessEvent)

	ruleset := event.GetResult()
	for tokenA, tokensB := range u.eventHelper.UnmatchedSkippedViolations() {
		for _, tokenB := range tokensB {
			ruleset.AddError(Error.NewError(fmt.Sprintf("Skipped violation \"%s\" for \"%s\" was not matched.", tokenB, tokenA)))
		}
	}
	stopPropagation()
	return nil
}

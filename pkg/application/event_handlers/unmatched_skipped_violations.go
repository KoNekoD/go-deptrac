package event_handlers

import (
	"fmt"
	"github.com/KoNekoD/go-deptrac/pkg/dispatchers"
	"github.com/KoNekoD/go-deptrac/pkg/domain/dtos/analysis_results/issues"
	"github.com/KoNekoD/go-deptrac/pkg/domain/events"
)

type UnmatchedSkippedViolations struct {
	eventHelper *dispatchers.EventHelper
}

func NewUnmatchedSkippedViolations(eventHelper *dispatchers.EventHelper) *UnmatchedSkippedViolations {
	return &UnmatchedSkippedViolations{eventHelper: eventHelper}
}

func (u *UnmatchedSkippedViolations) HandleEvent(rawEvent interface{}, stopPropagation func()) error {
	event := rawEvent.(*events.PostProcessEvent)

	ruleset := event.GetResult()
	for tokenA, tokensB := range u.eventHelper.UnmatchedSkippedViolations() {
		for _, tokenB := range tokensB {
			ruleset.AddError(issues.NewError(fmt.Sprintf("Skipped violation \"%s\" for \"%s\" was not matched.", tokenB, tokenA)))
		}
	}
	stopPropagation()
	return nil
}

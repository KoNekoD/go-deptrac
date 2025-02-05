package event_handlers

import (
	"fmt"
	"github.com/KoNekoD/go-deptrac/pkg/application/services"
	"github.com/KoNekoD/go-deptrac/pkg/domain/dtos/results/issues"
	"github.com/KoNekoD/go-deptrac/pkg/domain/events"
)

type UnmatchedSkippedViolations struct {
	eventHelper *services.EventHelper
}

func NewUnmatchedSkippedViolations(eventHelper *services.EventHelper) *UnmatchedSkippedViolations {
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

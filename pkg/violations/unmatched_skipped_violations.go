package violations

import (
	"fmt"
	"github.com/KoNekoD/go-deptrac/pkg/dispatchers"
	"github.com/KoNekoD/go-deptrac/pkg/domain/apperrors"
	"github.com/KoNekoD/go-deptrac/pkg/events"
)

type UnmatchedSkippedViolations struct {
	eventHelper *dispatchers.EventHelper
}

func NewUnmatchedSkippedViolations(eventHelper *dispatchers.EventHelper) *UnmatchedSkippedViolations {
	return &UnmatchedSkippedViolations{eventHelper: eventHelper}
}

func (u *UnmatchedSkippedViolations) InvokeEventSubscriber(rawEvent interface{}, stopPropagation func()) error {
	event := rawEvent.(*events.PostProcessEvent)

	ruleset := event.GetResult()
	for tokenA, tokensB := range u.eventHelper.UnmatchedSkippedViolations() {
		for _, tokenB := range tokensB {
			ruleset.AddError(apperrors.NewError(fmt.Sprintf("Skipped violation \"%s\" for \"%s\" was not matched.", tokenB, tokenA)))
		}
	}
	stopPropagation()
	return nil
}

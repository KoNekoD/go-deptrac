package post_process_event

import (
	"fmt"
	"github.com/KoNekoD/go-deptrac/pkg/analyser_contract/event_helper"
	"github.com/KoNekoD/go-deptrac/pkg/analyser_contract/post_process_event"
	"github.com/KoNekoD/go-deptrac/pkg/result_contract"
)

type UnmatchedSkippedViolations struct {
	eventHelper *event_helper.EventHelper
}

func NewUnmatchedSkippedViolations(eventHelper *event_helper.EventHelper) *UnmatchedSkippedViolations {
	return &UnmatchedSkippedViolations{eventHelper: eventHelper}
}

func (u *UnmatchedSkippedViolations) InvokeEventSubscriber(rawEvent interface{}, stopPropagation func()) error {
	event := rawEvent.(*post_process_event.PostProcessEvent)

	ruleset := event.GetResult()
	for tokenA, tokensB := range u.eventHelper.UnmatchedSkippedViolations() {
		for _, tokenB := range tokensB {
			ruleset.AddError(result_contract.NewError(fmt.Sprintf("Skipped violation \"%s\" for \"%s\" was not matched.", tokenB, tokenA)))
		}
	}
	stopPropagation()
	return nil
}

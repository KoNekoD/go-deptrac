package event_handlers

import (
	"github.com/KoNekoD/go-deptrac/pkg/domain/dtos/results/violations_rules"
	"github.com/KoNekoD/go-deptrac/pkg/domain/dtos/tokens"
	"github.com/KoNekoD/go-deptrac/pkg/domain/events"
)

type UncoveredDependent struct {
	ignoreUncoveredInternalClasses bool
}

func NewUncoveredDependent(ignoreUncoveredInternalClasses bool) *UncoveredDependent {
	return &UncoveredDependent{ignoreUncoveredInternalClasses: ignoreUncoveredInternalClasses}
}

func (h *UncoveredDependent) HandleEvent(rawEvent interface{}, stopPropagation func()) error {
	event := rawEvent.(*events.ProcessEvent)

	dependent := event.Dependency.GetDependent()

	ruleset := event.GetResult()

	if len(event.DependentLayers) != 0 {
		return nil
	}

	if dependentClassLike, ok := dependent.(*tokens.ClassLikeToken); ok {
		if !h.isIgnoreUncoveredInternalClasses(dependentClassLike) {
			ruleset.AddRule(violations_rules.NewUncovered(event.Dependency, event.DependerLayer))
		}
	}

	stopPropagation()

	return nil
}

func (h *UncoveredDependent) isIgnoreUncoveredInternalClasses(token *tokens.ClassLikeToken) bool {
	if !h.ignoreUncoveredInternalClasses {
		return false
	}

	tokenString := token.ToString()

	PhpStormStubsMap := map[string]bool{"ReturnTypeWillChange": true} // TODO: Add more stubs

	_, ok := PhpStormStubsMap[tokenString]
	return ok || "ReturnTypeWillChange" == tokenString
}

package dependencies

import (
	"github.com/KoNekoD/go-deptrac/pkg/events"
	"github.com/KoNekoD/go-deptrac/pkg/rules"
	"github.com/KoNekoD/go-deptrac/pkg/tokens"
)

type UncoveredDependentHandler struct {
	ignoreUncoveredInternalClasses bool
}

func NewUncoveredDependentHandler(ignoreUncoveredInternalClasses bool) *UncoveredDependentHandler {
	return &UncoveredDependentHandler{ignoreUncoveredInternalClasses: ignoreUncoveredInternalClasses}
}

func (h *UncoveredDependentHandler) InvokeEventSubscriber(rawEvent interface{}, stopPropagation func()) error {
	event := rawEvent.(*events.ProcessEvent)

	dependent := event.Dependency.GetDependent()

	ruleset := event.GetResult()

	if len(event.DependentLayers) != 0 {
		return nil
	}

	if dependentClassLike, ok := dependent.(*tokens.ClassLikeToken); ok {
		if !h.isIgnoreUncoveredInternalClasses(dependentClassLike) {
			ruleset.AddRule(rules.NewUncovered(event.Dependency, event.DependerLayer))
		}
	}

	stopPropagation()

	return nil
}

func (h *UncoveredDependentHandler) isIgnoreUncoveredInternalClasses(token *tokens.ClassLikeToken) bool {
	if !h.ignoreUncoveredInternalClasses {
		return false
	}

	tokenString := token.ToString()

	PhpStormStubsMap := map[string]bool{"ReturnTypeWillChange": true} // TODO: Add more stubs

	_, ok := PhpStormStubsMap[tokenString]
	return ok || "ReturnTypeWillChange" == tokenString
}
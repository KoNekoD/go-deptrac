package UncoveredDependentHandler

import (
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Analyser/ProcessEvent"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Result/Uncovered"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/Ast/AstMap"
)

type UncoveredDependentHandler struct {
	ignoreUncoveredInternalClasses bool
}

func NewUncoveredDependentHandler(ignoreUncoveredInternalClasses bool) *UncoveredDependentHandler {
	return &UncoveredDependentHandler{ignoreUncoveredInternalClasses: ignoreUncoveredInternalClasses}
}

func (h *UncoveredDependentHandler) InvokeEventSubscriber(rawEvent interface{}, stopPropagation func()) error {
	event := rawEvent.(*ProcessEvent.ProcessEvent)

	dependent := event.Dependency.GetDependent()

	ruleset := event.GetResult()

	if len(event.DependentLayers) != 0 {
		return nil
	}

	if dependentClassLike, ok := dependent.(*AstMap.ClassLikeToken); ok {
		if !h.isIgnoreUncoveredInternalClasses(dependentClassLike) {
			ruleset.AddRule(Uncovered.NewUncovered(event.Dependency, event.DependerLayer))
		}
	}

	stopPropagation()

	return nil
}

func (h *UncoveredDependentHandler) isIgnoreUncoveredInternalClasses(token *AstMap.ClassLikeToken) bool {
	if !h.ignoreUncoveredInternalClasses {
		return false
	}

	tokenString := token.ToString()

	PhpStormStubsMap := map[string]bool{"ReturnTypeWillChange": true} // TODO: Add more stubs

	_, ok := PhpStormStubsMap[tokenString]
	return ok || "ReturnTypeWillChange" == tokenString
}

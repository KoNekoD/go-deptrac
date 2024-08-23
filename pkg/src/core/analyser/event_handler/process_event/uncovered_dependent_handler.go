package process_event

import (
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Result/Uncovered"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/analyser/process_event"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/ast/ast_map"
)

type UncoveredDependentHandler struct {
	ignoreUncoveredInternalClasses bool
}

func NewUncoveredDependentHandler(ignoreUncoveredInternalClasses bool) *UncoveredDependentHandler {
	return &UncoveredDependentHandler{ignoreUncoveredInternalClasses: ignoreUncoveredInternalClasses}
}

func (h *UncoveredDependentHandler) InvokeEventSubscriber(rawEvent interface{}, stopPropagation func()) error {
	event := rawEvent.(*process_event.ProcessEvent)

	dependent := event.Dependency.GetDependent()

	ruleset := event.GetResult()

	if len(event.DependentLayers) != 0 {
		return nil
	}

	if dependentClassLike, ok := dependent.(*ast_map.ClassLikeToken); ok {
		if !h.isIgnoreUncoveredInternalClasses(dependentClassLike) {
			ruleset.AddRule(Uncovered.NewUncovered(event.Dependency, event.DependerLayer))
		}
	}

	stopPropagation()

	return nil
}

func (h *UncoveredDependentHandler) isIgnoreUncoveredInternalClasses(token *ast_map.ClassLikeToken) bool {
	if !h.ignoreUncoveredInternalClasses {
		return false
	}

	tokenString := token.ToString()

	PhpStormStubsMap := map[string]bool{"ReturnTypeWillChange": true} // TODO: Add more stubs

	_, ok := PhpStormStubsMap[tokenString]
	return ok || "ReturnTypeWillChange" == tokenString
}

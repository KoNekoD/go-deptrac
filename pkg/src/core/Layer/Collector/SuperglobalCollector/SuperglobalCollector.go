package SuperglobalCollector

import (
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Ast/TokenReferenceInterface"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Layer/InvalidCollectorDefinitionException"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/Ast/AstMap"
	"github.com/KoNekoD/go-deptrac/pkg/util"
)

type SuperglobalCollector struct{}

func NewSuperglobalCollector() *SuperglobalCollector {
	return &SuperglobalCollector{}
}

func (c SuperglobalCollector) Satisfy(config map[string]interface{}, reference TokenReferenceInterface.TokenReferenceInterface) (bool, error) {
	if _, ok := reference.(*AstMap.VariableReference); !ok {
		return false, nil
	}

	names, err := c.getNames(config)
	if err != nil {
		return false, err
	}

	for _, name := range names {
		if reference.GetToken().ToString() == name {
			return true, nil
		}
	}

	return false, nil
}

func (c SuperglobalCollector) getNames(config map[string]interface{}) ([]string, error) {
	if !util.MapKeyExists(config, "value") || !util.MapKeyIsArrayOfStrings(config, "value") {
		return nil, InvalidCollectorDefinitionException.NewInvalidCollectorDefinitionExceptionInvalidCollectorConfiguration("SuperglobalCollector needs the names configuration.")
	}

	return config["value"].([]string), nil
}

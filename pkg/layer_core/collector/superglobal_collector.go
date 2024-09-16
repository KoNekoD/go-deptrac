package collector

import (
	"github.com/KoNekoD/go-deptrac/pkg/ast_contract"
	"github.com/KoNekoD/go-deptrac/pkg/ast_core/ast_map"
	"github.com/KoNekoD/go-deptrac/pkg/layer_contract"
	"github.com/KoNekoD/go-deptrac/pkg/util"
)

type SuperglobalCollector struct{}

func NewSuperglobalCollector() *SuperglobalCollector {
	return &SuperglobalCollector{}
}

func (c SuperglobalCollector) Satisfy(config map[string]interface{}, reference ast_contract.TokenReferenceInterface) (bool, error) {
	if _, ok := reference.(*ast_map.VariableReference); !ok {
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
		return nil, layer_contract.NewInvalidCollectorDefinitionExceptionInvalidCollectorConfiguration("SuperglobalCollector needs the names configuration.")
	}

	return config["value"].([]string), nil
}

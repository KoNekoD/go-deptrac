package collector

import (
	"github.com/KoNekoD/go-deptrac/pkg/contract/ast"
	"github.com/KoNekoD/go-deptrac/pkg/contract/layer"
	"github.com/KoNekoD/go-deptrac/pkg/core/ast/ast_map"
	"github.com/KoNekoD/go-deptrac/pkg/util"
)

type SuperglobalCollector struct{}

func NewSuperglobalCollector() *SuperglobalCollector {
	return &SuperglobalCollector{}
}

func (c SuperglobalCollector) Satisfy(config map[string]interface{}, reference ast.TokenReferenceInterface) (bool, error) {
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
		return nil, layer.NewInvalidCollectorDefinitionExceptionInvalidCollectorConfiguration("SuperglobalCollector needs the names configuration.")
	}

	return config["value"].([]string), nil
}

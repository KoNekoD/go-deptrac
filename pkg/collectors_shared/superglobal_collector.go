package collectors_shared

import (
	"github.com/KoNekoD/go-deptrac/pkg/domain/apperrors"
	"github.com/KoNekoD/go-deptrac/pkg/domain/utils"
	"github.com/KoNekoD/go-deptrac/pkg/references"
	"github.com/KoNekoD/go-deptrac/pkg/tokens"
)

type SuperglobalCollector struct{}

func NewSuperglobalCollector() *SuperglobalCollector {
	return &SuperglobalCollector{}
}

func (c SuperglobalCollector) Satisfy(config map[string]interface{}, reference tokens.TokenReferenceInterface) (bool, error) {
	if _, ok := reference.(*references.VariableReference); !ok {
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
	if !utils.MapKeyExists(config, "value") || !utils.MapKeyIsArrayOfStrings(config, "value") {
		return nil, apperrors.NewInvalidCollectorDefinitionInvalidCollectorConfiguration("SuperglobalCollector needs the names configuration.")
	}

	return config["value"].([]string), nil
}

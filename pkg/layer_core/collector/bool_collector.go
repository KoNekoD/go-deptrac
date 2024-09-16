package collector

import (
	"fmt"
	"github.com/KoNekoD/go-deptrac/pkg/ast_contract"
	"github.com/KoNekoD/go-deptrac/pkg/config_contract"
	"github.com/KoNekoD/go-deptrac/pkg/layer_contract"
)

type BoolCollector struct {
	collectorResolver CollectorResolverInterface
}

func NewBoolCollector(collectorResolver CollectorResolverInterface) *BoolCollector {
	return &BoolCollector{
		collectorResolver: collectorResolver,
	}
}

func (b *BoolCollector) Satisfy(config map[string]interface{}, reference ast_contract.TokenReferenceInterface) (bool, error) {
	configuration, err := b.normalizeConfiguration(config)
	if err != nil {
		return false, err
	}

	for _, v := range configuration["must"].([]interface{}) {
		collectable, resolveErr := b.collectorResolver.Resolve(v.(map[string]interface{}))
		if resolveErr != nil {
			return false, resolveErr
		}
		satisfied, err := collectable.Collector.Satisfy(collectable.Attributes, reference)
		if err != nil {
			return false, err
		}
		if !satisfied {
			return false, nil
		}
	}
	for _, v := range configuration["must_not"].([]interface{}) {
		collectable, resolveErr := b.collectorResolver.Resolve(v.(map[string]interface{}))
		if resolveErr != nil {
			return false, resolveErr
		}
		satisfied, err := collectable.Collector.Satisfy(collectable.Attributes, reference)
		if err != nil {
			return false, err
		}
		if satisfied {
			return false, nil
		}
	}

	return true, nil
}

func (b *BoolCollector) normalizeConfiguration(configuration map[string]interface{}) (map[string]interface{}, error) {
	if _, ok := configuration["must"]; !ok {
		configuration["must"] = make([]*config_contract.CollectorConfig, 0)
	}

	if _, ok := configuration["must_not"]; !ok {
		configuration["must_not"] = make([]*config_contract.CollectorConfig, 0)
	}

	if len(configuration["must"].([]interface{})) == 0 && len(configuration["must_not"].([]interface{})) == 0 {
		return nil, layer_contract.NewInvalidCollectorDefinitionExceptionInvalidCollectorConfiguration(fmt.Sprintf("\"bool\" collector must have a \"must\" or a \"must_not\" attribute."))
	}

	return configuration, nil
}

package CollectorConfig

import (
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Config/CollectorType"
)

// CollectorConfig - Abstract
type CollectorConfig struct {
	CollectorType CollectorType.CollectorType
	Payload       map[string]interface{}
	private       bool
}

func NewCollectorConfig(collectorType CollectorType.CollectorType, payload map[string]interface{}, private bool) *CollectorConfig {
	return &CollectorConfig{
		CollectorType: collectorType,
		Payload:       payload,
		private:       private,
	}
}

func (c *CollectorConfig) ToArray() map[string]interface{} {
	data := map[string]interface{}{
		"private": c.private,
		"type":    string(c.CollectorType),
	}

	for key, value := range c.Payload {
		data[key] = value
	}

	return data
}

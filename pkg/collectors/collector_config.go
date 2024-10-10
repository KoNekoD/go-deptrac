package collectors

import (
	"github.com/KoNekoD/go-deptrac/pkg/domain/enums"
)

// CollectorConfig - Abstract
type CollectorConfig struct {
	CollectorType enums.CollectorType
	Payload       map[string]interface{}
	private       bool
}

func NewCollectorConfig(collectorType enums.CollectorType, payload map[string]interface{}, private bool) *CollectorConfig {
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

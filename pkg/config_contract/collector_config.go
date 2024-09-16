package config_contract

// CollectorConfig - Abstract
type CollectorConfig struct {
	CollectorType CollectorType
	Payload       map[string]interface{}
	private       bool
}

func NewCollectorConfig(collectorType CollectorType, payload map[string]interface{}, private bool) *CollectorConfig {
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

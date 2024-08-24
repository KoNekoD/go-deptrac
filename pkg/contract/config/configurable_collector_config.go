package config

import (
	"strings"
)

var escapees = []string{
	"\\", "\\\\", "\\\"", "\"", "\x00", "\x01", "\x02", "\x03", "\x04", "\x05", "\x06", "\x07", "\x08", "\t", "\n", "\v", "\f", "\r", "\x0e", "\x0f", "\x10", "\x11", "\x12", "\x13", "\x14", "\x15", "\x16", "\x17", "\x18", "\x19", "\x1a", "\x1b", "\x1c", "\x1d", "\x1e", "\x1f", "", "", " ", " ", " ",
}

var escaped = []string{
	"\\\\", "\\\"", "\\\\", "\\\"", "\\0", "\\x01", "\\x02", "\\x03", "\\x04", "\\x05", "\\x06", "\\a", "\\b", "\\t", "\\n", "\\v", "\\f", "\\r", "\\x0e", "\\x0f", "\\x10", "\\x11", "\\x12", "\\x13", "\\x14", "\\x15", "\\x16", "\\x17", "\\x18", "\\x19", "\\x1a", "\\e", "\\x1c", "\\x1d", "\\x1e", "\\x1f", "\\x7f", "\\N", "\\_", "\\L", "\\P",
}

type ConfigurableCollectorConfig struct {
	*CollectorConfig

	config string
}

func newConfigurableCollectorConfig(config string) *ConfigurableCollectorConfig {
	return &ConfigurableCollectorConfig{
		CollectorConfig: &CollectorConfig{},
		config:          config,
	}
}

func CreateConfigurableCollectorConfig(config string) *ConfigurableCollectorConfig {
	return newConfigurableCollectorConfig(regex(config))
}

func (c *ConfigurableCollectorConfig) ToArray() map[string]interface{} {
	data := map[string]interface{}{}

	for key, value := range c.CollectorConfig.ToArray() {
		data[key] = value
	}

	data["value"] = c.config

	return data
}

func regex(regex string) string {
	merged := make([]string, 0)

	if len(escaped) != len(escapees) {
		panic("length mismatch escaped and escapees")
	}

	for i := 0; i < len(escapees); i++ {
		merged = append(merged, escapees[i])
		merged = append(merged, escaped[i])
	}

	replaced := strings.NewReplacer(merged...).Replace(regex)

	return replaced
}

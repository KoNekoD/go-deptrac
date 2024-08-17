package ConfigurationCodeclimate

type severityType string

const (
	Failure   severityType = "failure"
	Skipped   severityType = "skipped"
	Uncovered severityType = "uncovered"
)

type ConfigurationCodeclimate struct {
	severityMap map[severityType]string
}

func NewConfigurationCodeclimateFromArray(array map[severityType]interface{}) *ConfigurationCodeclimate {
	severityUntyped, ok := array["severityType"]

	severity := make(map[severityType]string)

	if !ok {
		return newConfigurationCodeclimate(severity)
	}

	return newConfigurationCodeclimate(severityUntyped.(map[severityType]string))
}

func newConfigurationCodeclimate(severityMap map[severityType]string) *ConfigurationCodeclimate {
	return &ConfigurationCodeclimate{severityMap: severityMap}
}

func (c *ConfigurationCodeclimate) getSeverity(key severityType) *string {
	v, ok := c.severityMap[key]
	if !ok {
		return nil
	}

	return &v
}

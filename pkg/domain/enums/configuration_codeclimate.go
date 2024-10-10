package enums

type SeverityType string

const (
	SeverityTypeFailure   SeverityType = "failure"
	SeverityTypeSkipped   SeverityType = "skipped"
	SeverityTypeUncovered SeverityType = "uncovered"
)

type ConfigurationCodeclimate struct {
	severityMap map[SeverityType]string
}

func NewConfigurationCodeclimateFromArray(array map[SeverityType]interface{}) *ConfigurationCodeclimate {
	severityUntyped, ok := array["severityType"]

	severity := make(map[SeverityType]string)

	if !ok {
		return newConfigurationCodeclimate(severity)
	}

	return newConfigurationCodeclimate(severityUntyped.(map[SeverityType]string))
}

func newConfigurationCodeclimate(severityMap map[SeverityType]string) *ConfigurationCodeclimate {
	return &ConfigurationCodeclimate{severityMap: severityMap}
}

func (c *ConfigurationCodeclimate) GetSeverity(key SeverityType) *string {
	v, ok := c.severityMap[key]
	if !ok {
		return nil
	}

	return &v
}

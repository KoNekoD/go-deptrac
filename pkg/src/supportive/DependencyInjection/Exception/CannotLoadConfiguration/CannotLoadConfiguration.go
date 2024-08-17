package CannotLoadConfiguration

import "fmt"

type CannotLoadConfiguration struct {
	Message string
}

func (e *CannotLoadConfiguration) Error() string {
	return e.Message
}

func newCannotLoadConfiguration(message string) *CannotLoadConfiguration {
	return &CannotLoadConfiguration{Message: message}
}

func NewCannotLoadConfigurationFromConfig(filename string, message string) *CannotLoadConfiguration {
	return newCannotLoadConfiguration(fmt.Sprintf("Could not load %s. Reason: %s", filename, message))
}

func NewCannotLoadConfigurationFromServices(filename string, message string) *CannotLoadConfiguration {
	return newCannotLoadConfiguration(fmt.Sprintf("Could not load %s. Reason: %s", filename, message))
}

func NewCannotLoadConfigurationFromCache(filename string, message string) *CannotLoadConfiguration {
	return newCannotLoadConfiguration(fmt.Sprintf("Could not load %s. Reason: %s", filename, message))
}

package Warning

import (
	"fmt"
	"strings"
)

// Warning - Represents a situation that, while valid, is not recommended. This can be used to guide the end-user to a more proper solution.
type Warning struct {
	Message string
}

func NewWarning(message string) *Warning {
	return &Warning{Message: message}
}

func NewWarningTokenIsInMoreThanOneLayer(tokenName string, layerNames []string) *Warning {
	return &Warning{
		Message: fmt.Sprintf("%s is in more than one layer [%s]. It is recommended that one token should only be in one layer.", tokenName, strings.Join(layerNames, ", ")),
	}
}

func (w *Warning) ToString() string {
	return w.Message
}

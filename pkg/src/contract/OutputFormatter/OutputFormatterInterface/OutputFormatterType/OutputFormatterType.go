package OutputFormatterType

import "errors"

type OutputFormatterType string

const (
	GithubActions OutputFormatterType = "github-actions"
	Table         OutputFormatterType = "table"
)

var availableOutputFormatters = []OutputFormatterType{
	GithubActions,
	Table,
}

func NewOutputFormatterTypeFromString(input string) (OutputFormatterType, error) {
	for _, formatter := range availableOutputFormatters {
		if formatter == OutputFormatterType(input) {
			return formatter, nil
		}
	}

	return "", errors.New("invalid OutputFormatterType")
}

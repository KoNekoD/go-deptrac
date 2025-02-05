package enums

import "github.com/pkg/errors"

type OutputFormatterType string

const (
	GithubActions OutputFormatterType = "github-actions"
	Table         OutputFormatterType = "table"
	Baseline      OutputFormatterType = "baseline"
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

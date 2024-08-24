package output_formatter

import "strings"

const ProgressAdvanceDefault = 1

// OutputStyleInterface - Wrapper around Symfony OutputStyleInterface.
type OutputStyleInterface interface {
	Title(message string)
	Section(message string)
	Success(message StringOrArrayOfStrings)
	Error(message StringOrArrayOfStrings)
	Warning(message StringOrArrayOfStrings)
	Note(message StringOrArrayOfStrings)
	Caution(message StringOrArrayOfStrings)
	DefinitionList(list []StringOrArrayOfStringsOrTableSeparator)
	Table(headers []string, rows [][]string)
	// NewLine - Writes a new line, default 1
	NewLine(count int)
	// ProgressStart - default 0
	ProgressStart(max int)
	// ProgressAdvance - default 1
	ProgressAdvance(step int) error
	ProgressFinish() error
	IsVerbose() bool
	IsDebug() bool
}

type StringOrArrayOfStrings struct {
	Strings []string
	String  string
}

func (s StringOrArrayOfStrings) ToString() string {
	if s.String != "" {
		s.Strings = append([]string{s.String}, s.Strings...)
	}

	return strings.Join(s.Strings, "\n")
}

type StringOrArrayOfStringsOrTableSeparator struct {
	StringsMap     map[string]string
	String         string
	TableSeparator bool
}

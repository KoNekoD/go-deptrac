package results

// OutputInterface - Wrapper around Symfony OutputInterface.
type OutputInterface interface {
	WriteFormatted(message string)
	WriteLineFormatted(message StringOrArrayOfStrings)
	WriteRaw(message string)
	GetStyle() OutputStyleInterface
	IsVerbose() bool
	IsDebug() bool
}

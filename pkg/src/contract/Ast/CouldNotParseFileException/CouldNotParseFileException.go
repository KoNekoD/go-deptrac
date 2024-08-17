package CouldNotParseFileException

// CouldNotParseFileException - Exception thrown in a collector when it cannot parse a file.
type CouldNotParseFileException struct {
	Reason   string
	Previous error
}

func (c *CouldNotParseFileException) Error() string {
	return c.Reason
}

func NewCouldNotParseFileException(reason string, previous error) *CouldNotParseFileException {
	return &CouldNotParseFileException{Reason: reason, Previous: previous}
}

package violations

// FileOccurrence - Where in the file_supportive has the dependency_contract occurred.
type FileOccurrence struct {
	FilePath string
	Line     int
}

func NewFileOccurrence(filepath string, line int) *FileOccurrence {
	return &FileOccurrence{FilePath: filepath, Line: line}
}

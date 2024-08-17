package FileOccurrence

// FileOccurrence - Where in the file has the dependency occurred.
type FileOccurrence struct {
	Filepath string
	Line     int
}

func NewFileOccurrence(filepath string, line int) *FileOccurrence {
	return &FileOccurrence{Filepath: filepath, Line: line}
}

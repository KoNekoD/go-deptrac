package ast

// FileOccurrence - Where in the file has the dependency occurred.
type FileOccurrence struct {
	FilePath string
	Line     int
}

func NewFileOccurrence(filepath string, line int) *FileOccurrence {
	return &FileOccurrence{FilePath: filepath, Line: line}
}

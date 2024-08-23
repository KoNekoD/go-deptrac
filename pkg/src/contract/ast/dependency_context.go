package ast

// DependencyContext - Context of the dependency. Any additional info about where the dependency occurred.
type DependencyContext struct {
	FileOccurrence *FileOccurrence
	DependencyType DependencyType
}

func NewDependencyContext(fileOccurrence *FileOccurrence, dependencyType DependencyType) *DependencyContext {
	return &DependencyContext{FileOccurrence: fileOccurrence, DependencyType: dependencyType}
}

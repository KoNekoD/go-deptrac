package ast_contract

// DependencyContext - Context of the dependency_contract. Any additional info about where the dependency_contract occurred.
type DependencyContext struct {
	FileOccurrence *FileOccurrence
	DependencyType DependencyType
}

func NewDependencyContext(fileOccurrence *FileOccurrence, dependencyType DependencyType) *DependencyContext {
	return &DependencyContext{FileOccurrence: fileOccurrence, DependencyType: dependencyType}
}

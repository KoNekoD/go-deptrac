package dependencies

import (
	"github.com/KoNekoD/go-deptrac/pkg/violations"
)

// DependencyContext - Context of the dependency_contract. Any additional info about where the dependency_contract occurred.
type DependencyContext struct {
	FileOccurrence *violations.FileOccurrence
	DependencyType DependencyType
}

func NewDependencyContext(fileOccurrence *violations.FileOccurrence, dependencyType DependencyType) *DependencyContext {
	return &DependencyContext{FileOccurrence: fileOccurrence, DependencyType: dependencyType}
}

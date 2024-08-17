package DependencyContext

import (
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Ast/DependencyType"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Ast/FileOccurrence"
)

// DependencyContext - Context of the dependency. Any additional info about where the dependency occurred.
type DependencyContext struct {
	FileOccurrence *FileOccurrence.FileOccurrence
	DependencyType DependencyType.DependencyType
}

func NewDependencyContext(fileOccurrence *FileOccurrence.FileOccurrence, dependencyType DependencyType.DependencyType) *DependencyContext {
	return &DependencyContext{FileOccurrence: fileOccurrence, DependencyType: dependencyType}
}

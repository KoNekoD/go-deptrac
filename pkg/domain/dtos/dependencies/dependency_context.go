package dependencies

import (
	"github.com/KoNekoD/go-deptrac/pkg/domain/dtos"
	"github.com/KoNekoD/go-deptrac/pkg/domain/enums"
)

// DependencyContext - Context of the dependency_contract. Any additional info about where the dependency_contract occurred.
type DependencyContext struct {
	FileOccurrence *dtos.FileOccurrence
	DependencyType enums.DependencyType
}

func NewDependencyContext(fileOccurrence *dtos.FileOccurrence, dependencyType enums.DependencyType) *DependencyContext {
	return &DependencyContext{FileOccurrence: fileOccurrence, DependencyType: dependencyType}
}

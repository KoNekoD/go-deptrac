package dependencies

import (
	dependencies2 "github.com/KoNekoD/go-deptrac/pkg/domain/dtos/dependencies"
	"github.com/KoNekoD/go-deptrac/pkg/domain/dtos/tokens"
)

// DependencyInterface - Represents a dependency_contract between 2 tokens (depender and dependent).
type DependencyInterface interface {
	GetDepender() tokens.TokenInterface

	GetDependent() tokens.TokenInterface

	GetContext() *dependencies2.DependencyContext

	Serialize() []map[string]interface{}
}

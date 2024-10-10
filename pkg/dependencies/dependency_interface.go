package dependencies

import "github.com/KoNekoD/go-deptrac/pkg/tokens"

// DependencyInterface - Represents a dependency_contract between 2 tokens (depender and dependent).
type DependencyInterface interface {
	GetDepender() tokens.TokenInterface

	GetDependent() tokens.TokenInterface

	GetContext() *DependencyContext

	Serialize() []map[string]interface{}
}

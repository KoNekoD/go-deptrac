package dependency

import (
	"github.com/KoNekoD/go-deptrac/pkg/contract/ast"
)

// DependencyInterface - Represents a dependency between 2 tokens (depender and dependent).
type DependencyInterface interface {
	GetDepender() ast.TokenInterface

	GetDependent() ast.TokenInterface

	GetContext() *ast.DependencyContext

	Serialize() []map[string]interface{}
}

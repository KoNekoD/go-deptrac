package DependencyInterface

import (
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Ast/DependencyContext"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Ast/TokenInterface"
)

// DependencyInterface - Represents a dependency between 2 tokens (depender and dependent).
type DependencyInterface interface {
	GetDepender() TokenInterface.TokenInterface

	GetDependent() TokenInterface.TokenInterface

	GetContext() *DependencyContext.DependencyContext

	Serialize() []map[string]interface{}
}

package dependency_contract

import (
	ast_contract2 "github.com/KoNekoD/go-deptrac/pkg/ast_contract"
)

// DependencyInterface - Represents a dependency_contract between 2 tokens (depender and dependent).
type DependencyInterface interface {
	GetDepender() ast_contract2.TokenInterface

	GetDependent() ast_contract2.TokenInterface

	GetContext() *ast_contract2.DependencyContext

	Serialize() []map[string]interface{}
}

package layer

import (
	"github.com/KoNekoD/go-deptrac/pkg/contract/ast"
)

// CollectorInterface - A collector is responsible to tell whether an AST node (e.g. a specific class) is part of a layer.
type CollectorInterface interface {
	Satisfy(config map[string]interface{}, reference ast.TokenReferenceInterface) (bool, error)
}

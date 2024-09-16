package layer_contract

import (
	"github.com/KoNekoD/go-deptrac/pkg/ast_contract"
)

// CollectorInterface - A collector is responsible to tell whether an AST node (e.g. a specific class) is part of a layer_contract.
type CollectorInterface interface {
	Satisfy(config map[string]interface{}, reference ast_contract.TokenReferenceInterface) (bool, error)
}

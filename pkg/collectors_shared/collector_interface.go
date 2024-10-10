package collectors_shared

import (
	"github.com/KoNekoD/go-deptrac/pkg/tokens"
)

// CollectorInterface - A collector is responsible to tell whether an AST node (e.g. a specific class) is part of a layer_contract.
type CollectorInterface interface {
	Satisfy(config map[string]interface{}, reference tokens.TokenReferenceInterface) (bool, error)
}

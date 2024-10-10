package dependencies_collectors

import (
	"github.com/KoNekoD/go-deptrac/pkg/domain/dtos/tokens_references"
)

// CollectorInterface - A collector is responsible to tell whether an AST node (e.g. a specific class) is part of a layer_contract.
type CollectorInterface interface {
	Satisfy(config map[string]interface{}, reference tokens_references.TokenReferenceInterface) (bool, error)
}

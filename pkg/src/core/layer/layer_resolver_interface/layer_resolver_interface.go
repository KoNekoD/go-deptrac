package layer_resolver_interface

import (
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/ast"
)

type LayerResolverInterface interface {
	// GetLayersForReference - Returns a layer name and whether the dependency is public(true) or private(false)
	GetLayersForReference(reference ast.TokenReferenceInterface) (map[string]bool, error)

	IsReferenceInLayer(layer string, reference ast.TokenReferenceInterface) (bool, error)

	Has(layer string) (bool, error)
}

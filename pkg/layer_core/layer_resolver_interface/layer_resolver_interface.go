package layer_resolver_interface

import (
	"github.com/KoNekoD/go-deptrac/pkg/ast_contract"
)

type LayerResolverInterface interface {
	// GetLayersForReference - Returns a layer_contract name and whether the dependency_contract is public(true) or private(false)
	GetLayersForReference(reference ast_contract.TokenReferenceInterface) (map[string]bool, error)

	IsReferenceInLayer(layer string, reference ast_contract.TokenReferenceInterface) (bool, error)

	Has(layer string) (bool, error)
}

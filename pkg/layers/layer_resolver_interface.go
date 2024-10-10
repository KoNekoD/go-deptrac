package layers

import (
	"github.com/KoNekoD/go-deptrac/pkg/domain/dtos/tokens_references"
)

type LayerResolverInterface interface {
	// GetLayersForReference - Returns a layer_contract name and whether the dependency_contract is public(true) or private(false)
	GetLayersForReference(reference tokens_references.TokenReferenceInterface) (map[string]bool, error)

	IsReferenceInLayer(layer string, reference tokens_references.TokenReferenceInterface) (bool, error)

	Has(layer string) (bool, error)
}

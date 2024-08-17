package LayerResolverInterface

import "github.com/KoNekoD/go-deptrac/pkg/src/contract/Ast/TokenReferenceInterface"

type LayerResolverInterface interface {
	// GetLayersForReference - Returns a layer name and whether the dependency is public(true) or private(false)
	GetLayersForReference(reference TokenReferenceInterface.TokenReferenceInterface) (map[string]bool, error)

	IsReferenceInLayer(layer string, reference TokenReferenceInterface.TokenReferenceInterface) (bool, error)

	Has(layer string) (bool, error)
}

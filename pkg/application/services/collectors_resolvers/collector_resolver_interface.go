package collectors_resolvers

import (
	"github.com/KoNekoD/go-deptrac/pkg/application/dtos"
)

type CollectorResolverInterface interface {
	Resolve(config map[string]interface{}) (*dtos.Collectable, error)
}

package collectors_shared

import (
	"github.com/KoNekoD/go-deptrac/pkg/violations"
)

type CollectorResolverInterface interface {
	Resolve(config map[string]interface{}) (*violations.Collectable, error)
}

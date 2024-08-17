package CollectorResolverInterface

import (
	"github.com/KoNekoD/go-deptrac/pkg/src/core/Layer/Collector/Collectable"
)

type CollectorResolverInterface interface {
	Resolve(config map[string]interface{}) (*Collectable.Collectable, error)
}

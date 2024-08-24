package collector

type CollectorResolverInterface interface {
	Resolve(config map[string]interface{}) (*Collectable, error)
}

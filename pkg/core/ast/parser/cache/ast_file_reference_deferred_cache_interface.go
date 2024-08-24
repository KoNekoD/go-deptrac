package cache

type AstFileReferenceDeferredCacheInterface interface {
	AstFileReferenceCacheInterface
	Load() error
	Write() error
}

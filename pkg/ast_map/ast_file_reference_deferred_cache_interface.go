package ast_map

type AstFileReferenceDeferredCacheInterface interface {
	AstFileReferenceCacheInterface
	Load() error
	Write() error
}

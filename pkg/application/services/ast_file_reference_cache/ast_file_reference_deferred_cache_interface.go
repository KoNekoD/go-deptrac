package ast_file_reference_cache

type AstFileReferenceDeferredCacheInterface interface {
	AstFileReferenceCacheInterface
	Load() error
	Write() error
}

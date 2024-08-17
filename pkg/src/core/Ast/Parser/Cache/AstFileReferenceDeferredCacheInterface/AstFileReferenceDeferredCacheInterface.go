package AstFileReferenceDeferredCacheInterface

import "github.com/KoNekoD/go-deptrac/pkg/src/core/Ast/Parser/Cache/AstFileReferenceCacheInterface"

type AstFileReferenceDeferredCacheInterface interface {
	AstFileReferenceCacheInterface.AstFileReferenceCacheInterface
	Load() error
	Write() error
}

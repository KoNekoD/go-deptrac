package app

import (
	"github.com/KoNekoD/go-deptrac/pkg/ast_map"
	"github.com/KoNekoD/go-deptrac/pkg/subscribers"
)

func Cache(builder *ContainerBuilder) {
	builder.AstFileReferenceFileCache = ast_map.NewAstFileReferenceFileCache(*builder.CacheFile, Version)
	builder.AstFileReferenceDeferredCacheInterface = builder.AstFileReferenceFileCache
	builder.AstFileReferenceCacheInterface = builder.AstFileReferenceFileCache
	builder.CacheableFileSubscriber = subscribers.NewCacheableFileSubscriber(builder.AstFileReferenceDeferredCacheInterface)
}

package app

import (
	"github.com/KoNekoD/go-deptrac/pkg/application/event_handlers"
	"github.com/KoNekoD/go-deptrac/pkg/ast_map"
)

func Cache(builder *ContainerBuilder) {
	builder.AstFileReferenceFileCache = ast_map.NewAstFileReferenceFileCache(*builder.CacheFile, Version)
	builder.AstFileReferenceDeferredCacheInterface = builder.AstFileReferenceFileCache
	builder.AstFileReferenceCacheInterface = builder.AstFileReferenceFileCache
	builder.CacheableFileSubscriber = event_handlers.NewCacheableFile(builder.AstFileReferenceDeferredCacheInterface)
}

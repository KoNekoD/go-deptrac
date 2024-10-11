package di

import (
	"github.com/KoNekoD/go-deptrac/pkg/application/event_handlers"
	"github.com/KoNekoD/go-deptrac/pkg/application/services/ast_file_reference_cache"
	"github.com/KoNekoD/go-deptrac/pkg/infrastructure/app"
)

func Cache(builder *ContainerBuilder) {
	builder.AstFileReferenceFileCache = ast_file_reference_cache.NewAstFileReferenceFileCache(*builder.CacheFile, app.Version)
	builder.AstFileReferenceDeferredCacheInterface = builder.AstFileReferenceFileCache
	builder.AstFileReferenceCacheInterface = builder.AstFileReferenceFileCache
	builder.CacheableFileSubscriber = event_handlers.NewCacheableFile(builder.AstFileReferenceDeferredCacheInterface)
}

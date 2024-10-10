package violations

import (
	"github.com/KoNekoD/go-deptrac/pkg/app"
	"github.com/KoNekoD/go-deptrac/pkg/ast_map"
	"github.com/KoNekoD/go-deptrac/pkg/subscribers"
)

func Cache(builder *app.ContainerBuilder) {
	builder.AstFileReferenceFileCache = ast_map.NewAstFileReferenceFileCache(*builder.CacheFile, app.Version)
	builder.AstFileReferenceDeferredCacheInterface = builder.AstFileReferenceFileCache
	builder.AstFileReferenceCacheInterface = builder.AstFileReferenceFileCache
	builder.CacheableFileSubscriber = subscribers.NewCacheableFileSubscriber(builder.AstFileReferenceDeferredCacheInterface)
}

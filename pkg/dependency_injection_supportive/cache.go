package dependency_injection_supportive

import (
	"github.com/KoNekoD/go-deptrac/pkg/ast_core/parser/cache"
	"github.com/KoNekoD/go-deptrac/pkg/console_supportive/application/application_version"
	"github.com/KoNekoD/go-deptrac/pkg/dependency_injection_supportive/container_builder"
)

func Cache(builder *container_builder.ContainerBuilder) {
	builder.AstFileReferenceFileCache = cache.NewAstFileReferenceFileCache(*builder.CacheFile, application_version.Version)
	builder.AstFileReferenceDeferredCacheInterface = builder.AstFileReferenceFileCache
	builder.AstFileReferenceCacheInterface = builder.AstFileReferenceFileCache
	builder.CacheableFileSubscriber = cache.NewCacheableFileSubscriber(builder.AstFileReferenceDeferredCacheInterface)
}

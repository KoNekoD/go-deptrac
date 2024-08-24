package dependency_injection

import (
	Cache2 "github.com/KoNekoD/go-deptrac/pkg/core/ast/parser/cache"
	"github.com/KoNekoD/go-deptrac/pkg/supportive/console/application/application_version"
	"github.com/KoNekoD/go-deptrac/pkg/supportive/dependency_injection/container_builder"
)

func Cache(builder *container_builder.ContainerBuilder) {
	builder.AstFileReferenceFileCache = Cache2.NewAstFileReferenceFileCache(*builder.CacheFile, application_version.Version)
	builder.AstFileReferenceDeferredCacheInterface = builder.AstFileReferenceFileCache
	builder.AstFileReferenceCacheInterface = builder.AstFileReferenceFileCache
	builder.CacheableFileSubscriber = Cache2.NewCacheableFileSubscriber(builder.AstFileReferenceDeferredCacheInterface)
}

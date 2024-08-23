package DependencyInjection

import (
	Cache2 "github.com/KoNekoD/go-deptrac/pkg/src/core/ast/parser/cache"
	"github.com/KoNekoD/go-deptrac/pkg/src/supportive/Console/Application/ApplicationVersion"
	"github.com/KoNekoD/go-deptrac/pkg/src/supportive/DependencyInjection/ContainerBuilder"
)

func Cache(builder *ContainerBuilder.ContainerBuilder) {
	builder.AstFileReferenceFileCache = Cache2.NewAstFileReferenceFileCache(*builder.CacheFile, ApplicationVersion.Version)
	builder.AstFileReferenceDeferredCacheInterface = builder.AstFileReferenceFileCache
	builder.AstFileReferenceCacheInterface = builder.AstFileReferenceFileCache
	builder.CacheableFileSubscriber = Cache2.NewCacheableFileSubscriber(builder.AstFileReferenceDeferredCacheInterface)
}

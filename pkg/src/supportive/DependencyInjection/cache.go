package DependencyInjection

import (
	"github.com/KoNekoD/go-deptrac/pkg/src/core/Ast/Parser/Cache/AstFileReferenceFileCache"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/Ast/Parser/Cache/CacheableFileSubscriber"
	"github.com/KoNekoD/go-deptrac/pkg/src/supportive/Console/Application/ApplicationVersion"
	"github.com/KoNekoD/go-deptrac/pkg/src/supportive/DependencyInjection/ContainerBuilder"
)

func Cache(builder *ContainerBuilder.ContainerBuilder) {
	builder.AstFileReferenceFileCache = AstFileReferenceFileCache.NewAstFileReferenceFileCache(*builder.CacheFile, ApplicationVersion.Version)
	builder.AstFileReferenceDeferredCacheInterface = builder.AstFileReferenceFileCache
	builder.AstFileReferenceCacheInterface = builder.AstFileReferenceFileCache
	builder.CacheableFileSubscriber = CacheableFileSubscriber.NewCacheableFileSubscriber(builder.AstFileReferenceDeferredCacheInterface)
}

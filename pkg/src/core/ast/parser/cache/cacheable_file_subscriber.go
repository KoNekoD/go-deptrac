package cache

import (
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/ast"
)

type CacheableFileSubscriber struct {
	deferredCache AstFileReferenceDeferredCacheInterface
}

func NewCacheableFileSubscriber(deferredCache AstFileReferenceDeferredCacheInterface) *CacheableFileSubscriber {
	return &CacheableFileSubscriber{
		deferredCache: deferredCache,
	}
}

func (s *CacheableFileSubscriber) InvokeEventSubscriber(rawEvent interface{}, stopPropagation func()) error {
	switch rawEvent.(type) {
	case *ast.PreCreateAstMapEvent:
		err := s.deferredCache.Load()
		if err != nil {
			return err
		}
	case *ast.PostCreateAstMapEvent:
		err := s.deferredCache.Write()
		if err != nil {
			return err
		}
	}

	return nil
}

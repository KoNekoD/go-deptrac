package cache

import (
	ast_contract2 "github.com/KoNekoD/go-deptrac/pkg/ast_contract"
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
	case *ast_contract2.PreCreateAstMapEvent:
		err := s.deferredCache.Load()
		if err != nil {
			return err
		}
	case *ast_contract2.PostCreateAstMapEvent:
		err := s.deferredCache.Write()
		if err != nil {
			return err
		}
	}

	return nil
}

package cache

import (
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Ast/PostCreateAstMapEvent"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Ast/PreCreateAstMapEvent"
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
	case *PreCreateAstMapEvent.PreCreateAstMapEvent:
		err := s.deferredCache.Load()
		if err != nil {
			return err
		}
	case *PostCreateAstMapEvent.PostCreateAstMapEvent:
		err := s.deferredCache.Write()
		if err != nil {
			return err
		}
	}

	return nil
}

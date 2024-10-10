package subscribers

import "github.com/KoNekoD/go-deptrac/pkg/ast_map"

type CacheableFileSubscriber struct {
	deferredCache ast_map.AstFileReferenceDeferredCacheInterface
}

func NewCacheableFileSubscriber(deferredCache ast_map.AstFileReferenceDeferredCacheInterface) *CacheableFileSubscriber {
	return &CacheableFileSubscriber{
		deferredCache: deferredCache,
	}
}

func (s *CacheableFileSubscriber) InvokeEventSubscriber(rawEvent interface{}, stopPropagation func()) error {
	switch rawEvent.(type) {
	case *ast_map.PreCreateAstMapEvent:
		err := s.deferredCache.Load()
		if err != nil {
			return err
		}
	case *ast_map.PostCreateAstMapEvent:
		err := s.deferredCache.Write()
		if err != nil {
			return err
		}
	}

	return nil
}

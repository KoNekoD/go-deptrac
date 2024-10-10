package event_handlers

import "github.com/KoNekoD/go-deptrac/pkg/ast_map"

type CacheableFile struct {
	deferredCache ast_map.AstFileReferenceDeferredCacheInterface
}

func NewCacheableFile(deferredCache ast_map.AstFileReferenceDeferredCacheInterface) *CacheableFile {
	return &CacheableFile{deferredCache: deferredCache}
}

func (s *CacheableFile) HandleEvent(rawEvent interface{}, stopPropagation func()) error {
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
